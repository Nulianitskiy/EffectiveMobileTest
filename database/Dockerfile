FROM postgres:alpine3.18

ENV POSTGRES_PASSWORD=stoic\
    POSTGRES_USER=dbuser\
    POSTGRES_DB=tasktrackerdb

COPY ./db.sql /docker-entrypoint-initdb.d/db.sql

EXPOSE 5432