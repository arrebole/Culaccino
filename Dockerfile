FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

# 数据库位置
ENV DB_NAME="./database/culaccino.db"

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN go build -o culaccino ./src/main.go && mkdir database

CMD [ "./culaccino" ]

EXPOSE 3000 3000