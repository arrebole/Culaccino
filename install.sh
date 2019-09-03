echo "删除所有容器..."
docker rm -f $(docker ps -aq)
echo "创建文件夹 $(pwd)/wwwroot/static"
mkdir -p $(pwd)/wwwroot/static
echo "创建docker网络..."
docker network create culaccino
# 启动容器
echo "启动容器..."
docker run \
--net culaccino \
--name culaccino \
-p 8080:3000 \
-e PASSWORD root@password \
-v $(pwd)/database:/var/www/cuaccino/data \
-v $(pwd)/wwwroot/static:/var/www/culaccino/wwwroot/static \
-d arrebole/culaccino