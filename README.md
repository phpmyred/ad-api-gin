# GIN-API

用GIN 搭建的api 服务服务

初始化
go mod tidy

打包命令
## linux
```
go env -w GOOS=linux
go build main.go
```
## window
```
go env -w GOOS=windows
go build -o main.exe main.go

```

## linux 部署命令

```
nohup ./main  > start.log 2>&1 &
```

