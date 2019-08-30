FROM redis:5.0.4-alpine

# 创建目录
RUN mkdir -p /redis/{db,conf}

# 挂载配置文件
ADD redis.conf /redis/conf/redis.conf

# 启动
CMD [ "redis-server", "/redis/conf/redis.conf" ]