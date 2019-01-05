# elastic使用

## 节点查看 [https://github.com/lmenezes/cerebro](https://github.com/lmenezes/cerebro)
```
bin/cerebro -Dhttp.port=1234 -Dhttp.address=127.0.0.1
```
## 版本必须为6.x,采用6.5.4
- 启动一个节点
```
./bin/elasticsearch -Ecluster.name=base_framework -Enode.name=node1 -Epath.data=base_framework_node1 -Ehttp.port=5200 -d
./bin/elasticsearch -Ecluster.name=base_framework -Enode.name=node2 -Epath.data=base_framework_node2 -Ehttp.port=5300 -d
./bin/elasticsearch -Ecluster.name=base_framework -Enode.name=node3 -Epath.data=base_framework_node3 -Ehttp.port=5400 -d

```
- 浏览器输入 http://localhost:1234 添加 http://localhost:5200 查看3个节点的集群