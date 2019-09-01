docker build -t arrebole/Culaccino .
docker network create net
mkdir -p $(pwd)/wwwroot/static
docker run \
--net net \
--name culaccino \
-v $(pwd)/data:/var/www/Cuaccino/data \
-v $(pwd)/wwwroot/static:/var/www/Culaccino/wwwroot/static \
-d arrebole/Culaccino