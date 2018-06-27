#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o ad_default
rsync -az --delete --progress /home/tonnn/go/src/zonstfe_api/ad_default/ad_default   -e 'ssh ' tonnn@111.231.141.43:/home/tonnn/service/
rm ad_default