### redis容器
```
cd make/common/redis
docker build -t "my-redis:v1" .
docker run --name my-redis -p 6379:6379 -v /Users/liushaobo/data/redis:/redis/db --restart=always -d my-redis:v1
```

### mysql容器
```
cd make/common/mysql
docker build -t "my-mysql:v1" .
docker run --name my-mysql -p 3306:3306 -v /Users/liushaobo/data/mysql:/var/lib/mysql --restart=always -e MYSQL_ROOT_PASSWORD=root -d go-gin-web:v1
```

```
docker run --name my-mysql -p 3306:3306 -v /Users/liushaobo/data/mysql:/var/lib/mysql --restart=always -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
```

### 性能测试
```
# 登陆
wrk -t1 -c10 -d60s -T5s --script=sigin.lua --latency http://127.0.0.1:9090

# 注册
wrk -t1 -c10 -d60s -T5s --script=sigup.lua --latency http://127.0.0.1:9090
```
限流测试
seq 10 | xargs -P10 -I% curl localhost:9090/health

# todo list
[] 自动重启
[] service优化
[] 前端管理站开发
[] 监控实现
[] 定时任务实现
[] 异步框架
[] mongodb集成
[] es集成
[] 工具函数库
[] 容器化
[] 自动化测试
[] 性能测试
[] 配置文件
[] 校验库

