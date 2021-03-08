FROM golang:alpine

USER root
RUN mkdir /go/app

ADD . /go/app
WORKDIR /go/app

# Install tesseract and data
RUN apk update && apk upgrade && \
    apk add gcc g++ bash tesseract-ocr-dev

# Get best trainned data
ENV OCR_LANGS=vie,jpn,fra
ENV TESSDATA_DIR=/usr/share/tessdata/
RUN chmod +x ./get-traineddatas.sh
RUN ./get-traineddatas.sh

# Build ocr server
RUN go get all && \
    go build -o bin/gocrapi

ENV PORT "8080"

EXPOSE 8080

ENTRYPOINT ["/go/app/bin/gocrapi"]
