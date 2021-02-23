go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go env
go build -o gfva cmd/main.go
./gfva initdb