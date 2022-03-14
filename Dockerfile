FROM golang:1.17-alpine

RUN mkdir -p /usr/src/bot
WORKDIR /usr/src/bot

COPY . .

CMD ["make", "run"]
