FROM alpine

WORKDIR /

COPY ./main  /

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 8080

RUN apk upgrade --update-cache --available && \
        apk add openssl && \
        rm -rf /var/cache/apk/*

ENTRYPOINT ./main