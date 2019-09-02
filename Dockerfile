FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

ENV GIN_MODE="release"

RUN mkdir -p /var/www/culaccino

COPY ./ /var/www/culaccino
WORKDIR /var/www/culaccino

RUN go build -o ./bin/service.out ./service/main.go

CMD [ "./bin/service.out" ]

EXPOSE 3000 3000