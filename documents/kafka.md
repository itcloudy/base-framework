# Kafka部署和使用


###  Docker部署kafka
```sh
docker pull landoop/fast-data-dev:2.0.1

docker run --rm -p 2181:2181 -p 3030:3030 -p 8081-8083:8081-8083 \
       -p 9581-9585:9581-9585 -p 9092:9092 -e ADV_HOST=192.168.1.194 \
       landoop/fast-data-dev:2.0.1
```

### 浏览器查看
[http://192.168.1.190:3030](http://192.168.1.190:3030)