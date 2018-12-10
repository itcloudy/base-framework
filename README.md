# 说明

## 介绍

* 改进项目结构，采用行业标准项目结构
* 配置文件模板代码生成，toml格式
* 支持数据库升级，数据库访问接口化
* 日志写入到文件，文件备份大小可配置，同时支持日志写到Kafka
* 系统表对应的struct间的关系可视化
* 同时提供rest接口和rpc接口
* 支持controller，service层测单元测试

## 相关命令

```sh
 ./base-framework -h
base-framework application

Usage:
  base-framework [command]

Available Commands:
  config      Initial config generation
  dbinit      init database,create tables
  dbupdate    update database, add tables and columns or modify columns
  help        Help about any command
  loaddata    load data from init folder
  start       Starting node

Flags:
      --config string   filepath to config.toml (default "/Users/cloudy/Documents/go/src/github.com/itcloudy/base-framework/config/config.toml")
  -h, --help            help for base-framework

Use "base-framework [command] --help" for more information about a command.

```


相关文档在 documents下