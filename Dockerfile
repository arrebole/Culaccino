FROM ubuntu

LABEL authorMail "concurrent.exec@gmail.com"

ENV GIN_MODE = "release"
RUN apt update && apt upgrade; mkdir -p /var/www/Culaccino
COPY ./ /var/www/Culaccino
WORKDIR /var/www/Culaccino

VOLUME "./cuaccino.db" "/var/www/Culaccino/data/culaccino.db"

CMD [ "./bin/service.out" ]

EXPOSE 3000 3000