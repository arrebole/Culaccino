#!/bin/bash
cat << EOF
------------------------------------------
|                                        |
|        install cuaccino                |
|                                        |
------------------------------------------
EOF

RMcontainer(){
    echo "删除所有容器..."
    docker rm -f $(docker ps -aq)
}

MkdirStatic(){
    echo "创建文件夹 $(pwd)/wwwroot/static"
    mkdir -p $(pwd)/wwwroot/static
}

createDockerNetWork(){
    echo "创建docker网络..."
    docker network create culaccino
}

DockerBuild(){
    echo "docker build"
    docker build -t arrebole/culaccino:v1.2 -t arrebole/culaccino:latest .

}

Run(){
    echo -n "请输入root密码 ->"
    read PASSWD
   # 运行redis
    docker run \
        --net culaccino \
        --name redis-master \
        -v $(pwd)/database:/data \
        -d redis
    # 运行博客
    docker run \
        --net culaccino \
        --name blog \
        -e PASSWORD=$PASSWD \
        -e DB_ADDR=redis-master:6379 \
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
}

main(){
    RMcontainer
    MkdirStatic
    createDockerNetWork
    DockerBuild
    Run
}

main $@