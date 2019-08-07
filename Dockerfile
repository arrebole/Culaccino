FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

RUN apt update && apt upgrade; mkdir -p /var/www/Culaccino

WORKDIR /var/www/Culaccino
ENV GIN_MODE = "release"

COPY ./ /var/www/Culaccino
VOLUME "./cuaccino.db" "/var/www/Culaccino/data/culaccino.db"

CMD [ "./bin/service.out" ]

EXPOSE 3000 3000