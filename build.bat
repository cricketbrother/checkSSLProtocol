@echo off

cls

cd ./bin

echo Initialize Golang Environment
go env -u GOARCH
go env -u GOOS

go env -w GOARCH=amd64

echo Build Windows Executable Program
go env -w GOOS=windows
go build ../

echo Build Linux Executable Program
go env -w GOOS=linux
go build ../

echo Reset Golang Environment
go env -u GOARCH
go env -u GOOS

echo Done
