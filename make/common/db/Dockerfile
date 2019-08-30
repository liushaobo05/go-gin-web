FROM mysql:5.7

WORKDIR /tmp

ADD db.sql db.sql

ADD docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod u+x /docker-entrypoint.sh

# 启动命令
ENTRYPOINT ["docker-entrypoint.sh"]