echo "删除所有容器..."
docker rm -f $(docker ps -aq)
echo "创建文件夹 $(pwd)/wwwroot/static"
mkdir -p $(pwd)/wwwroot/static
echo "创建docker网络..."
docker network create culaccino
echo "docker build"
docker build -t arrebole/culaccino:v1.2 -t arrebole/culaccino:latest .
# 启动容器
echo "启动容器..."
# 运行博客
docker run \
--net culaccino \
--name blog \
-e PASSWORD=root@password \
-v $(pwd)/database:/var/www/culaccino/database \
-v $(pwd)/wwwroot/static:/var/www/culaccino/wwwroot/static \
-d arrebole/culaccino
# 运行nginx
docker run \
--net culaccino \
-p 80:80 \
-p 443:443 \
--name nginx-master \
-v $(pwd)/config/nginx.conf:/etc/nginx/conf.d/default.conf \
-v ~/.acme.sh/gkdark.xyz/gkdark.xyz.key:/root/gkdark.xyz.key \
-v ~/.acme.sh/gkdark.xyz/fullchain.cer:/root/fullchain.cer \
-d nginx