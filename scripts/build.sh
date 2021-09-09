#!/bin/bash --posix

# Args: source, target
# eg: build.sh ./main.go ./bin/program

AUTH=$(whoami)
ENV=$(uname -snr)
VER=$(cat .version)
UPT=$(date +"%Y/%m/%d %T %z")
TAG=$(echo $(git rev-parse --short HEAD)$([ -n "$(git status -s)"  ] && echo "-dev" || echo ""))

echo "making $2"
mkdir -p bin
echo "Version:${TAG}"

go mod tidy

go build -ldflags "                                     \
    -installsuffix 'static'                             \
    -s -w                                               \
    -X '$(go list -m)/pkg/version.verStr=${VER}'\
    -X '$(go list -m)/pkg/version.tagStr=${TAG}'\
    -X '$(go list -m)/pkg/version.uptStr=${UPT}'\
    -X '$(go list -m)/pkg/version.envStr=${ENV}'\
    -X '$(go list -m)/pkg/version.authStr=${AUTH}'\
    " \
    -o $2 $1
