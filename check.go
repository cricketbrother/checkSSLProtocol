package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// 获取网站
func getSites(domain string, port string, file string) ([]string, error) {
	var sites []string
	if file != "" {
		sites, err := getSitesFromFile(file)
		if err != nil {
			return nil, err
		}
		return sites, nil
	} else if domain != "" {
		sites = append(sites, domain+":"+port)
		return sites, nil
	} else {
		return nil, fmt.Errorf("No domain or file specified\n")
	}
}

// 从文件中获取网站
func getSitesFromFile(file string) ([]string, error) {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sites []string
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			sites = append(sites, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}
	return sites, nil
}

// 获取工具版本
func getToolVersion() (string, string, bool, error) {
	cmd := exec.Command("nmap", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", false, err
	}

	nmapVersionRegex := regexp.MustCompile(`\d+\.\d+`)
	nmapVersion := nmapVersionRegex.FindString(string(out))

	opensslVersionRegex := regexp.MustCompile(`openssl-\d+\.\d+\.\d+`)
	opensslVersion := strings.Replace(opensslVersionRegex.FindString(string(out)), "openssl-", "", 1)

	var tlsv13Flag bool
	if compareOpenSSLVersion(opensslVersion) {
		tlsv13Flag = true
	} else {
		tlsv13Flag = false
	}

	return nmapVersion, opensslVersion, tlsv13Flag, nil
}

// 比较openssl版本
func compareOpenSSLVersion(version string) bool {
	versionSlice := strings.Split(version, ".")
	majorVersion, _ := strconv.Atoi(versionSlice[0])
	minorVersion, _ := strconv.Atoi(versionSlice[1])
	if majorVersion > 1 {
		return true
	}
	if majorVersion == 1 && minorVersion > 0 {
		return true
	}
	return false
}

// 通过nmap检查网站是否支持某个协议
func checkSiteByNmap(site string, tlsv13Flag bool) []string {
	domain := site[:strings.Index(site, ":")]
	port := site[strings.Index(site, ":")+1:]
	cmd := exec.Command("nmap", "--script", "ssl-enum-ciphers", "-p", port, domain, "2>&1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		println(err.Error())
		return nil
	}
	passFlag := true
	siteFlag := false
	row := []string{site}
	ipv4Regex := regexp.MustCompile(`Nmap scan report for.*\)`)
	tmpSlice := strings.Split(ipv4Regex.FindString(string(out)), " ")
	ipv4 := "--"
	if len(tmpSlice) > 1 {
		ipv4 = strings.ReplaceAll(strings.ReplaceAll(tmpSlice[len(tmpSlice)-1], ")", ""), "(", "")
	}
	row = append(row, ipv4)
	if strings.Contains(string(out), "SSLv3") {
		passFlag = false
		siteFlag = true
		row = append(row, "Yes")
	} else {
		row = append(row, "No")
	}
	if strings.Contains(string(out), "TLSv1.0") {
		passFlag = false
		siteFlag = true
		row = append(row, "Yes")
	} else {
		row = append(row, "No")
	}
	if strings.Contains(string(out), "TLSv1.1") {
		passFlag = false
		siteFlag = true
		row = append(row, "Yes")
	} else {
		row = append(row, "No")
	}
	if strings.Contains(string(out), "TLSv1.2") {
		siteFlag = true
		row = append(row, "Yes")
	} else {
		row = append(row, "No")
	}
	if tlsv13Flag {
		if strings.Contains(string(out), "TLSv1.3") {
			siteFlag = true
			row = append(row, "Yes")
		} else {
			row = append(row, "No")
		}
	} else {
		row = append(row, "Unknown")
	}
	if siteFlag {
		if passFlag {
			row = append(row, "PASS")
		} else {
			row = append(row, "FAIL")
		}
	} else {
		row = []string{site, ipv4, "Unknown", "Unknown", "Unknown", "Unknown", "Unknown", "Unknown"}
	}
	return row
}
