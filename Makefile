all: build_all

build_all:
	mkdir -p ./build
	env GOOS=linux GOARCH=amd64    go build -o ./build/GDDoS-linux-amd64
	env GOOS=linux GOARCH=arm64    go build -o ./build/GDDoS-linux-arm64
	env GOOS=linux GOARCH=arm      go build -o ./build/GDDoS-linux-arm
	env GOOS=linux GOARCH=mips     go build -o ./build/GDDoS-linux-mips
	env GOOS=linux GOARCH=mipsle   go build -o ./build/GDDoS-linux-mipsle
	env GOOS=linux GOARCH=mips64   go build -o ./build/GDDoS-linux-mips64
	env GOOS=linux GOARCH=mips64le go build -o ./build/GDDoS-linux-mips64le
	env GOOS=darwin GOARCH=amd64   go build -o ./build/GDDoS-macos-amd64
	env GOOS=darwin GOARCH=arm64   go build -o ./build/GDDoS-macos-arm64

.PHONY: build_all
