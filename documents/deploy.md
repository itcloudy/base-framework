# 部署

* 在GOPATH的src/github.com下创建一个文件夹 itcloudy,然后克隆代码,如需要把项目放到src目录下

删除项目代码中的包引入前缀github.com/itcloudy
```sh
git clone https://github.com/itcloudy/base-framework.git
```
* 编译
```sh
go build
```
* 项目目录下创建下列文件夹
logs 用于保存日志

system-data   用于保存进程信息和系统文件

user-data 用于保存用户上传文件

* 生成配置文件，修改配置文件信息
```sh
./base-framework config #生成配置文件
````

* 初始化数据库
```sh
./base-framework  dbinit
```

* 启动服务
```sh
./base-framework  start
```

ps: 前期为了方便在start命令中创建了系统接口的表，若要看系统接口示例

停止start进程,执行下面的命令再启动start
```sh
./base-framework  loaddata
```