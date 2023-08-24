# checkSSLProtocol
## Description
A tool to check site(s) SSL/TLS protocol support.
## Requirements
Nmap and OpenSSL must been installed in your operating system. 

OpenSSL version needs 1.0.2 or higher. 

OpenSSL 1.1.1 or higher additional supports check TLSv1.3, otherwise only support check SSLv3, TLSv1.0, TLSv1.1 and TLSv1.2.
## How To Use
```
PS > .\checkSSLProtocol.windows.amd64.exe -h
checkSSLProtocol v2023.08.24.104500, a tool to check SSL/TLS protocol support

Usage:
  checkSSLProtocol [-d domain] [-p port] [-f file] [-m max threads]
Options:
  -d string
        domain
  -f string
        a file contains sites, one site per line
  -m int
        max threads (default 10)
  -p string
        port (default "443")
Notice:
  if -d, -p and -f are specified at the same time, -d and -p will be ignored
Examples:
  checkSSLProtocol -d www.baidu.com -p 443
  checkSSLProtocol -d www.baidu.com
  checkSSLProtocol -f sites.txt
  checkSSLProtocol -f sites.txt -m 20
```
## Snapshot
![](https://github.com/cricketbrother/checkSSLProtocol/raw/main/snapshot.png)