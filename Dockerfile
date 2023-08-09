FROM golang:alpine as builder

USER root
RUN mkdir /go/app

ADD . /go/app
WORKDIR /go/app

# Install tesseract and data
RUN apk add --no-cache --update g++ musl-dev tesseract-ocr-dev

# Build ocr server
RUN go get all; \
    go build -o bin/gocrapi;

FROM alpine as runner

USER root
RUN apk add --no-cache --update bash tesseract-ocr-dev

# Get best trainned data
ENV OCR_LANGS=vie,jpn,fra
ENV TESSDATA_PREFIX=/usr/share/tessdata/
COPY get-traineddatas.sh /tmp/
COPY --from=builder /go/app/bin/gocrapi /usr/local/bin/gocrapi
RUN bash /tmp/get-traineddatas.sh

ENV PORT "8080"

EXPOSE 8080

CMD ["gocrapi"]
