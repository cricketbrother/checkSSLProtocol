package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Job参数结构体
type JobArgs struct {
	Id   int
	Site string
}

// 生产者
func producer(sites []string, jobArgsChan chan<- JobArgs) {
	for i, site := range sites {
		jobArgsChan <- JobArgs{Id: i + 1, Site: site}
	}
	close(jobArgsChan)
}

// 消费者
func consumer(jobArgsChan <-chan JobArgs, tlsv13Flag bool, resultChan chan<- []string, progressChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for jobArgs := range jobArgsChan {
		resultChan <- checkSiteByNmap(jobArgs.Site, tlsv13Flag)
		progressChan <- jobArgs.Id
	}
}

// 检查网站
func checkSites(sites []string, tlsv13Flag bool, maxThreads int) [][]string {
	fmt.Printf("Max Threads: %d\n", maxThreads)
	jobArgsChan := make(chan JobArgs, len(sites))
	resultChan := make(chan []string, len(sites))
	completedJobsChan := make(chan int, len(sites))
	var wg sync.WaitGroup
	go producer(sites, jobArgsChan)
	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go consumer(jobArgsChan, tlsv13Flag, resultChan, completedJobsChan, &wg)
	}
	go func() {
		count := 0
		for range completedJobsChan {
			count++
			progress := float64(count) / float64(len(sites)) * 100
			fmt.Printf("\rProgress: %.2f%%", progress)
		}
	}()
	wg.Wait()
	close(resultChan)
	close(completedJobsChan)
	var data [][]string
	for result := range resultChan {
		data = append(data, result)
	}
	return data
}

func main() {
	// 获取命令行参数
	domain, port, file, maxThreads := initFlag()

	// 获取网站列表
	sites, err := getSites(domain, port, file)
	if err != nil {
		println("unable to get sites")
		return
	}

	// 获取工具信息
	nmapVersion, opensslVersion, tlsv13Flag, err := getToolVersion()
	if err != nil {
		println("unable to get tool version")
		return
	}
	fmt.Printf("Nmap version:    %s\n", nmapVersion)
	fmt.Printf("OpenSSL version: %s\n", opensslVersion)
	if tlsv13Flag {
		fmt.Printf("TLSv1.3 Support: %s\n\n", color.New(color.FgGreen, color.Bold).Sprint("Yes"))
	} else {
		fmt.Printf("TLSv1.3 Support: %s\n\n", color.New(color.FgRed, color.Bold).Sprint("No"))
	}

	// 检查网站
	data := checkSites(sites, tlsv13Flag, maxThreads)

	// 打印结果
	time.Sleep(3 * time.Second)
	tableRender(data)
	csvRender(data)
}
