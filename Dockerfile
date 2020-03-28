FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

WORKDIR /srv/culaccino

COPY ./ /srv/culaccino

RUN go build

CMD [ "./culaccino" ]

EXPOSE 3000 3000