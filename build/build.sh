#!/bin/bash

set -e

go get github.com/PuerkitoBio/goquery
GOOS=linux GOARCH=386 go build -o gouscis.linux ..
docker build -t gouscis .

echo
echo 'Docker container was successfully built'
echo 'Run "docker run -e EMAIL=my@email.com -e SMTP_SERVER=127.0.0.1:25 -e CASE_NUMBER=MSC0000000000 --restart=always gouscis" to start it.'
