#!/bin/sh
go build -ldflags "-X github.com/fcoders/api-template/settings.Version=`date -u +%Y%m%d.%H%M%S` -X github.com/fcoders/api-template/settings.CommitHash=$(git log --pretty=format:'%h' -n 1)"
