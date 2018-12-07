# 数据库表关系svg图

该功能是基于[https://github.com/gmarik/go-erd](https://github.com/gmarik/go-erd)来实现

base-framework中的全表关系在`pkg/models/all_fileds`下

## 操作步骤

* 获得该模块
```sh
go get github.com/gmarik/go-erd
```

* 在项目根目录下执行下面的操作，生成到documents目录下
```sh
go-erd -path $(pwd)/pkg/models/all_fields |dot -Tsvg > $(pwd)/documents/db-table-relation.svg

```

* 查看数据库表关系svg图
```sh
open $(pwd)/documents/db-table-relation.svg
```