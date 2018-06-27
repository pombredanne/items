#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o adv_admin
rsync -az --delete --progress /home/tonnn/go/src/zonstfe_api/advertiser_admin/adv_admin   -e 'ssh ' tonnn@111.231.137.127:/home/tonnn/zonstfe_api/
rm adv_admin
rm advertiser_admin




