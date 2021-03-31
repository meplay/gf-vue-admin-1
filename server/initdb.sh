go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go env
go run cmd/gfva/main.go initdb
