if ! [ -x "$(command -v docker)" ]; then
    echo "install docker ..."
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
    add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable edge test"
    #apt-cache policy docker-ce
fi
apt-get update
# 更新docker
apt-get install -y docker-ce
# 删除全部容器
docker rm -f $(docker ps -aq)
# 创建docker网络
docker network create futanaicha
# 构建镜像
docker build -t arrebole/Culaccino:v0.8 .
mkdir -p $(pwd)/wwwroot/static
# 启动容器
docker run \
--net net \
--name culaccino \
-p 8080:3000 \
-v $(pwd)/data:/var/www/Cuaccino/data \
-v $(pwd)/wwwroot/static:/var/www/Culaccino/wwwroot/static \
-d arrebole/Culaccino