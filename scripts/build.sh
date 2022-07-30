mkdir -p ./build
version=`cat version` # Version
timestamp=`date "+%Y-%m-%d %H:%M:%S"` # Timestamp

# Build for Linux 🎁
env GOOS=linux GOARCH=amd64    go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-amd64 ./src/
echo Built Linux-amd64!
env GOOS=linux GOARCH=arm64    go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-arm64 ./src/
echo Built Linux-arm64!
env GOOS=linux GOARCH=arm      go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-arm ./src/
echo Built Linux-arm!
env GOOS=linux GOARCH=mips     go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-mips ./src/
echo Built Linux-mips!
env GOOS=linux GOARCH=mipsle   go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-mipsle ./src/
echo Built Linux-mipsle!
env GOOS=linux GOARCH=mips64   go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-mips64 ./src/
echo Built Linux-mips64!
env GOOS=linux GOARCH=mips64le go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/linux/AGDDoS-linux-mips64le ./src/
echo Built Linux-mips64le!

# Build for Macos(Darwin) 🎁
env GOOS=darwin GOARCH=amd64   go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/darwin/AGDDoS-darwin-amd64 ./src/
echo Built MacOS-amd64!
env GOOS=darwin GOARCH=arm64   go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/darwin/AGDDoS-darwin-arm64 ./src/
echo Built MacOS-arm64!

# Build for Windows 🎁
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-amd64.exe ./src/
echo Built Windows-amd64!
env CGO_ENABLED=0 GOOS=windows GOARCH=386  go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-x86.exe ./src/
echo Built Windows-x86!
env CGO_ENABLED=0 GOOS=windows GOARCH=arm64  go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-arm64.exe ./src/
echo Built Windows-arm64!

env CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags="-H windowsgui -s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-amd64-hidden.exe ./src/
echo Built Windows-amd64-hidden!
env CGO_ENABLED=0 GOOS=windows GOARCH=386  go build -ldflags="-H windowsgui -s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-x86-hidden.exe ./src/
echo Built Windows-x86-hidden!
env CGO_ENABLED=0 GOOS=windows GOARCH=arm64  go build -ldflags="-H windowsgui -s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/windows/AGDDoS-arm64-hidden.exe ./src/
echo Built Windows-arm64-hidden!

# Build for Freebzd 🎁
env GOOS=freebsd GOARCH=amd64  go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/freebsd/AGDDoS-freebsd-amd64 ./src/
echo Built Freebzd-amd64!
env GOOS=freebsd GOARCH=386  go build -ldflags "-s -w -X main.version=$version -X \"main.timestamp=$timestamp\"" -o ./build/freebsd/AGDDoS-freebsd-x86 ./src/
echo Built Freebzd-386!

# Exit
exit 0
