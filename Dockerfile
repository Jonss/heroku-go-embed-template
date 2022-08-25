FROM alpine:3.15.1 as alpine

COPY app/bin app/bin
COPY ui/build ui/build
COPY pkg/db/migrations pkg/db/migrations

CMD ["app/bin"]