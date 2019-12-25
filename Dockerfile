FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

RUN mkdir -p /var/www/culaccino

COPY ./ /srv/culaccino
WORKDIR /srv/culaccino

RUN go build -o service.out ./service/main.go

CMD [ "./service.out" ]

EXPOSE 3000 3000