#!/usr/bin/env bash
ssh tonnn@111.231.137.127 "cp -f /home/tonnn/zonstfe_api/adv_admin /home/tonnn/zonstfe_api/adv_admin_bak"
GOOS=linux GOARCH=amd64 go build -o adv_admin
rsync -az --delete --progress /home/tonnn/go/src/adv-admin-api/adv_admin   -e 'ssh ' tonnn@111.231.137.127:/home/tonnn/zonstfe_api/
rm adv_admin
rm adv-admin-api