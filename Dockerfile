#----------------------------------------------->

FROM node:latest as htmlBuilder

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN cd theme \
    && npm install \
    && npm run build

# ----------------------------------------------->

FROM golang:latest

LABEL authorMail "concurrent.exec@gmail.com"

WORKDIR /srv/culaccino

# 数据库位置
ENV DB_NAME="./database/culaccino.db"

# 数据库匿名卷
VOLUME ./database

COPY ./ /srv/culaccino
COPY --from=htmlBuilder /srv/culaccino/theme/build /srv/culaccino/theme/build

RUN go build -o culaccino.out ./src/main.go

CMD [ "./culaccino.out" ]

EXPOSE 3000 3000