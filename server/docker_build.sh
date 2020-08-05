export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64
export GOPROXY=https://goproxy.io,direct
export CGO_ENABLED=1
export CC=/usr/local/gcc-4.8.0-for-linux64/bin/x86_64-pc-linux-gcc
go env
go mod tidy
go mod verify
go build -o main .