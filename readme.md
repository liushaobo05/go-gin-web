### redis容器
cd make/common/redis
docker build -t "my-redis:v1" .
docker run --name my-redis -p 6379:6379 -v /Users/liushaobo/data/redis:/redis/db --restart=always -d my-redis:v1

### mysql容器
cd make/common/mysql
docker build -t "my-mysql:v1" .
docker run --name my-mysql -p 3306:3306 -v /Users/liushaobo/data/mysql:/var/lib/mysql --restart=always -e MYSQL_ROOT_PASSWORD=root -d go-gin-web:v1


docker run --name my-mysql -p 3306:3306 -v /Users/liushaobo/data/mysql:/var/lib/mysql --restart=always -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7


