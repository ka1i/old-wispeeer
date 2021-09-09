#!/bin/bash --posix

# Args: source
# eg: install.sh ./main.go

AUTH=$(whoami)
ENV=$(uname -snr)
VER=$(cat .version)
UPT=$(date +"%Y/%m/%d %T %z")
TAG=$(echo $(git rev-parse --short HEAD)$([ -n "$(git status -s)"  ] && echo "-dev" || echo ""))

echo "$1 installing ..."
mkdir -p bin
echo "Version:${TAG}"

go mod tidy

go install -ldflags "                                \
    -installsuffix 'static'                          \
    -s -w                                            \
    -X '$(go list -m)/pkg/version.verStr=${VER}'     \
    -X '$(go list -m)/pkg/version.tagStr=${TAG}'     \
    -X '$(go list -m)/pkg/version.uptStr=${UPT}'     \
    -X '$(go list -m)/pkg/version.envStr=${ENV}'     \
    -X '$(go list -m)/pkg/version.authStr=${AUTH}'   \
    "                                                \
    ./...
