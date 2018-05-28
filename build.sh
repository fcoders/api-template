#!/bin/sh
go build -ldflags "-X astropay/go-web-template/settings.Version=`date -u +%Y%m%d.%H%M%S` -X astropay/go-web-template/settings.CommitHash=$(git log --pretty=format:'%h' -n 1)"
