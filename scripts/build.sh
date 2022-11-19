if [ "$1" = "dev" ]; then
  version="gh-actions"
else
  version=$(git describe --abbrev=0 --tags)
fi
appName="AGDDoS"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor=$(git show -s --format='format:%aN <%ae>' HEAD)
gitCommit=$(git log --pretty=format:"%h" -1)
ldflags="\
-w -s \
-X 'main.AppName=$appName' \
-X 'main.BuiltAt=$builtAt' \
-X 'main.GoVersion=$goVersion' \
-X 'main.GitAuthor=$gitAuthor' \
-X 'main.GitCommit=$gitCommit' \
-X 'main.Version=$version' \
"

mkdir -p ./build
# Build for Windows 🎁
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags="$ldflags" -o ./build/windows/AGDDoS-amd64.exe ./src/
echo Built Windows-amd64!
env CGO_ENABLED=0 GOOS=windows GOARCH=386    go build -ldflags="$ldflags" -o ./build/windows/AGDDoS-x86.exe ./src/
echo Built Windows-x86!
env CGO_ENABLED=0 GOOS=windows GOARCH=arm64  go build -ldflags="$ldflags" -o ./build/windows/AGDDoS-arm64.exe ./src/
echo Built Windows-arm64!

# Build for Linux 🎁
env GOOS=linux GOARCH=amd64    go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-amd64 ./src/
echo Built Linux-amd64!
env GOOS=linux GOARCH=arm64    go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-arm64 ./src/
echo Built Linux-arm64!
env GOOS=linux GOARCH=arm      go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-arm ./src/
echo Built Linux-arm!
env GOOS=linux GOARCH=mips     go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-mips ./src/
echo Built Linux-mips!
env GOOS=linux GOARCH=mipsle   go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-mipsle ./src/
echo Built Linux-mipsle!
env GOOS=linux GOARCH=mips64   go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-mips64 ./src/
echo Built Linux-mips64!
env GOOS=linux GOARCH=mips64le go build -ldflags="$ldflags" -o ./build/linux/AGDDoS-linux-mips64le ./src/
echo Built Linux-mips64le!

# Build for Macos(Darwin) 🎁
env GOOS=darwin GOARCH=amd64   go build -ldflags="$ldflags" -o ./build/darwin/AGDDoS-darwin-amd64 ./src/
echo Built MacOS-amd64!
env GOOS=darwin GOARCH=arm64   go build -ldflags="$ldflags" -o ./build/darwin/AGDDoS-darwin-arm64 ./src/
echo Built MacOS-arm64!

# Build for Freebzd 🎁
env GOOS=freebsd GOARCH=amd64  go build -ldflags="$ldflags" -o ./build/freebsd/AGDDoS-freebsd-amd64 ./src/
echo Built Freebzd-amd64!
env GOOS=freebsd GOARCH=386  go build -ldflags="$ldflags" -o ./build/freebsd/AGDDoS-freebsd-x86 ./src/
echo Built Freebzd-386!

# Exit
exit 0
