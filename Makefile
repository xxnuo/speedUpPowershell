.PHONY: all build compress

all: build compress

build:
	go build -ldflags "-s -w -H=windowsgui" -o speedUpPowershell.exe main.go

compress: build
	upx --brute speedUpPowershell.exe
