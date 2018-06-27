#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o static_upload
rsync -az --delete --progress /home/tonnn/go/src/zonstfe_api/static_upload/static_upload   -e 'ssh ' tonnn@111.231.141.43:/home/tonnn/service/
