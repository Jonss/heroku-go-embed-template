FROM alpine:3.15.1 as alpine

COPY app/bin app/bin
COPY ui/build ui/build

CMD ["app/bin"]