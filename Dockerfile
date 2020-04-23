#----------------------------------------------->

FROM node:latest as htmlBuilder

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN cd theme \
    && npm install \
    && npm run build

# ----------------------------------------------->

FROM golang:latest as goBuilder

LABEL authorMail "concurrent.exec@gmail.com"

ENV CGO_ENABLED=0
WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN go build -tags netgo -a -o culaccino.out ./src/main.go

# ----------------------------------------------->

FROM alpine:latest

LABEL authorMail "concurrent.exec@gmail.com"

WORKDIR /srv/culaccino

# 数据库位置
ENV DB_NAME="./database/culaccino.db"

# 数据库匿名卷
VOLUME ./database

COPY ./ /srv/culaccino
COPY --from=htmlBuilder /srv/culaccino/theme/build /srv/culaccino/theme/build
COPY --from=goBuilder /srv/culaccino/culaccino.out /srv/culaccino/culaccino.out

ENTRYPOINT [ "./culaccino.out" ]

EXPOSE 3000 3000