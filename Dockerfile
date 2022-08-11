FROM alpine:3.15.1 as alpine

COPY app/bin app/bin
COPY react-app/build react-app/build

CMD ["app/bin"]