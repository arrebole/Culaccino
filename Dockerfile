FROM postgres
LABEL authorMail "c.y.concurrent@gmail.com"

ENV POSTGRES_PASSWORD=e8y28diq9hx POSTGRES_USER=postgres

RUN apt update && apt upgrade

VOLUME "/var/lib/postgresql/data/pgdata" "/var/lib/postgresql/data"

EXPOSE 5432 5432