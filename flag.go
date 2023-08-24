package main

import "flag"

var (
	version string = "local-build"
)

// 获取参数
func initFlag() (string, string, string, int) {
	flag.Usage = func() {
		println("checkSSLProtocol " + version + ", a tool to check SSL/TLS protocol support")
		println()
		println("Usage:")
		println("  checkSSLProtocol [-d domain] [-p port] [-f file] [-m max threads]")
		println("Options:")
		flag.PrintDefaults()
		println("Notice:")
		println("  if -d, -p and -f are specified at the same time, -d and -p will be ignored")
		println("Examples:")
		println("  checkSSLProtocol -d www.baidu.com -p 443")
		println("  checkSSLProtocol -d www.baidu.com")
		println("  checkSSLProtocol -f sites.txt")
		println("  checkSSLProtocol -f sites.txt -m 20")
	}
	domain := flag.String("d", "", "domain")
	port := flag.String("p", "443", "port")
	file := flag.String("f", "", "a file contains sites, one site per line")
	maxThreads := flag.Int("m", 10, "max threads")
	flag.Parse()
	return *domain, *port, *file, *maxThreads
}
