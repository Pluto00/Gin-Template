# Gin Template

## Usage

- conf：存放配置文件

- docs：存放 swagger 相关的，一般不用动，用 `swag init` 命令自动生成
  
- middleware：存放项目用到的 middleware 函数

- models：数据库相关

- pkg：引用的第三方库

- router：路由目录

- service：业务代码

## commit message

- feat：新功能（feature）
- fix：修补 bug
- docs：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试 
- chore：构建过程或辅助工具的变动

## build in linux

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GinTemplate .`

## go mod

`go mod init Gin-Template`

## swagger

`swag init`