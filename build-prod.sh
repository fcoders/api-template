#!/usr/bin/env bash

# Configure GO variables
GOOS=linux
GOARCH=x86_64
export GOPATH=$(dirname $(dirname $(dirname $(pwd))))

# Resolve internal dependencies
#rm -Rf "$GOPATH/src/github.com/astropay/apc-common"
#git clone -q ssh://git-codecommit.eu-west-1.amazonaws.com/v1/repos/apc-common "$GOPATH/src/github.com/astropay/apc-common" --branch master

# Execute compilation
echo "Compiling project..."
go get ./...
go build -ldflags "-X astropay/go-web-template/settings.Version=$1 -X astropay/go-web-template/settings.CommitHash=$(git log --pretty=format:'%h' -n 1)" -o target/go-web-template
