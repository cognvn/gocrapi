#!/bin/bash
wget -O "${TESSDATA_PREFIX}eng.traineddata" "https://github.com/tesseract-ocr/tessdata_best/raw/main/eng.traineddata"
langarray=(`echo $OCR_LANGS | tr ',' ' '`)
for lang in "${langarray[@]}"
do
  url="https://github.com/tesseract-ocr/tessdata_best/raw/main/${lang}.traineddata"
  if [[  `wget -S --spider $url  2>&1 | grep 'HTTP/1.1 200 OK'` ]] ; then
    wget -O "${TESSDATA_PREFIX}${lang}.traineddata" $url
  fi
done