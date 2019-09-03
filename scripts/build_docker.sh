if ! [ -x "$(command -v docker)" ]; then
    echo "install docker ..."
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
    add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable edge test"
    #apt-cache policy docker-ce
fi
apt update
# 更新docker
apt install -y docker-ce
# 构建镜像
docker build -t arrebole/culaccino:latest arrebole/culaccino:v1.0 --rm .