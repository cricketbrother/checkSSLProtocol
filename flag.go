package main

import "flag"

// 获取参数
func initFlag() (string, string, string, int) {
	flag.Usage = func() {
		println("A tool to check SSL/TLS protocol support")
		println("Usage: checkSSLProtocol2 [-d domain] [-p port] [-f file] [-m max threads]")
		println()
		println("Options:")
		flag.PrintDefaults()
		println()
		println("Notice: if -d, -p and -f are specified at the same time, -d and -p will be ignored")
		println()
		println("Example 1: checkSSLProtocol2 -d www.baidu.com -p 443")
		println("Example 2: checkSSLProtocol2 -d www.baidu.com")
		println("Example 3: checkSSLProtocol2 -f sites.txt")
		println("Example 4: checkSSLProtocol2 -f sites.txt -m 20")
	}
	domain := flag.String("d", "", "domain")
	port := flag.String("p", "443", "port")
	file := flag.String("f", "", "a file contains sites, one site per line")
	maxThreads := flag.Int("m", 10, "max threads")
	flag.Parse()
	return *domain, *port, *file, *maxThreads
}
