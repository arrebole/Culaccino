#---------------- build assert ----------------->

FROM node:latest as htmlBuilder

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN cd theme \
    && npm install \
    && npm run build

# --------------- build server ------------------->

FROM golang:alpine as goBuilder

LABEL authorMail "concurrent.exec@gmail.com"

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN apk add --update gcc musl-dev \
    && go build -o culaccino.out ./src/main.go

# -------------- build docker image -------------->

FROM alpine:latest

LABEL authorMail "concurrent.exec@gmail.com"

WORKDIR /srv/culaccino

# 数据库位置
ENV DB_NAME="./database/culaccino.db"

RUN mkdir database

COPY ./ /srv/culaccino
COPY --from=htmlBuilder /srv/culaccino/theme/build /srv/culaccino/theme/build
COPY --from=goBuilder /srv/culaccino/culaccino.out /srv/culaccino/culaccino.out

CMD [ "./culaccino.out" ]

EXPOSE 80