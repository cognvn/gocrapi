FROM ghcr.io/cognvn/gocrapi:latest

WORKDIR /go/app

ENV PORT=7000
EXPOSE 7000
ENTRYPOINT ["/go/app/bin/gocrapi"]