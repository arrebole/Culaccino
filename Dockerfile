FROM golang

LABEL authorMail "concurrent.exec@gmail.com"

RUN mkdir -p /srv/culaccino

COPY ./ /srv/culaccino
WORKDIR /srv/culaccino

RUN go build

CMD [ "./culaccino" ]

EXPOSE 3000 3000