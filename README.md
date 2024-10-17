# chat-server

chat-server 相关服务

初始化
go mod tidy

打包命令
go env -w GOOS=linux
go build main.go

go env -w GOOS=windows
部署命令

nohup ./main  > start.log 2>&1 &

go env -w CGO_ENABLED=0

go build -o main.exe main.go

