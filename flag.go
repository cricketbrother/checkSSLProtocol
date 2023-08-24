package main

import "flag"

var (
	version string
)

// 获取参数
func initFlag() (string, string, string, int) {
	flag.Usage = func() {
		println("A tool to check SSL/TLS protocol support")
		println("Version:    " + version)
		println()
		println("Usage: checkSSLProtocol2 [-d domain] [-p port] [-f file] [-m max threads]")
		println("Options:")
		flag.PrintDefaults()
		println("Notice: if -d, -p and -f are specified at the same time, -d and -p will be ignored")
		println("Examples:")
		println("\tcheckSSLProtocol2 -d www.baidu.com -p 443")
		println("\tcheckSSLProtocol2 -d www.baidu.com")
		println("\tcheckSSLProtocol2 -f sites.txt")
		println("\tcheckSSLProtocol2 -f sites.txt -m 20")
	}
	domain := flag.String("d", "", "domain")
	port := flag.String("p", "443", "port")
	file := flag.String("f", "", "a file contains sites, one site per line")
	maxThreads := flag.Int("m", 10, "max threads")
	flag.Parse()
	return *domain, *port, *file, *maxThreads
}
