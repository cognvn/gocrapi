FROM golang:alpine

USER root
RUN mkdir /go/app

ADD . /go/app
WORKDIR /go/app

# Install tesseract and data
RUN apk add --no-cache --update g++ bash musl-dev tesseract-ocr-dev

# Get best trainned data
ENV OCR_LANGS=vie,jpn,fra
ENV TESSDATA_PREFIX=/usr/share/tessdata/
RUN chmod +x ./get-traineddatas.sh
RUN ./get-traineddatas.sh

# Build ocr server
RUN go get all && \
    go build -o /usr/local/bin/gocrapi

ENV PORT "8080"

EXPOSE 8080

CMD ["gocrapi"]
