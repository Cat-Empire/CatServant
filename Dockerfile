FROM golang:1.17-alpine

RUN apk update && apk add make \ 
                          git \
                          bash \
                          openssh-client && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /usr/src/bot
WORKDIR /usr/src/bot

COPY . .

CMD ["make", "run"]
