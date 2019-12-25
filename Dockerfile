FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

RUN mkdir -p /var/www/culaccino

COPY ./ /var/www/culaccino
WORKDIR /var/www/culaccino

RUN go build -o service.out ./service/main.go

CMD [ "./service.out" ]

EXPOSE 3000 3000