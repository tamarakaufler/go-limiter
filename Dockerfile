FROM alpine:3.7
RUN apk add --no-cache openssh ca-certificates

RUN mkdir /app
WORKDIR /app
COPY burstlimiter /app
ENTRYPOINT ["./burstlimiter"]
