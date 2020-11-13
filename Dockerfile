FROM golang:alpine

USER root
RUN mkdir /src

ADD . /src
WORKDIR /src
# Install tesseract and data
RUN apk update && apk upgrade && \
    apk add gcc g++ tesseract-ocr-dev tesseract-ocr-data-vie
# Get best trainned data
RUN wget -O /usr/share/tessdata/vie.traineddata https://github.com/tesseract-ocr/tessdata_best/blob/master/vie.traineddata?raw=true
# Build ocr server
RUN go get -u all && \
    go build -o ocrviet

ENV PORT "8080"
CMD ["/src/ocrviet"]
