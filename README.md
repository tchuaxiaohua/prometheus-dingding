### 项目介绍

#### Windows 平台构建 Linux下运行
~~~shell
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build  -ldflags "-s -w" -o dingtalk  main.go
~~~


