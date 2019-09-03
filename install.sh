# echo "删除所有容器..."
# docker rm -f $(docker ps -aq)
echo "创建文件夹 $(pwd)/wwwroot/static"
mkdir -p $(pwd)/wwwroot/static
echo "创建docker网络..."
docker network create culaccino
# 启动容器
echo "启动容器..."
# 运行nginx
docker run \
-net culaccino \
-p 80:80 \
-p 443:443 \
--name nginx-master \
-v /root/Repository/default.conf:/etc/nginx/conf.d/default.conf \
-v /root/.acme.sh/gkdark.xyz/gkdark.xyz.key:/root/gkdark.xyz.key \
-v /root/.acme.sh/gkdark.xyz/fullchain.cer:/root/fullchain.cer \
-d nginx
# 运行博客
docker run \
--net culaccino \
--name blog \
-e PASSWORD root@password \
-v $(pwd)/database:/var/www/cuaccino/data \
-v $(pwd)/wwwroot/static:/var/www/culaccino/wwwroot/static \
-d arrebole/culaccino