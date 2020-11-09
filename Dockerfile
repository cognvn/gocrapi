FROM golang:alpine

USER root
RUN mkdir /src

ADD . /src
WORKDIR /src
RUN apk update && apk upgrade && \
    apk add gcc g++ tesseract-ocr-dev tesseract-ocr-data-vie

RUN go get -u all && \
    go build -o ocrviet

ENV PORT "8080"
CMD ["/src/ocrviet"]
