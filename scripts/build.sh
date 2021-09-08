#!/bin/bash --posix

# Args: source, target
# eg: build.sh ./main.go ./bin/program

TAG=$(git describe --tags --always --dirty="-dev")
UPT=$(date +"%Y/%m/%d %T %z")
ENV=$(uname -snr)
AUTH=$(whoami)

echo "making $2"
mkdir -p bin
echo "Version:${TAG}"

go mod tidy

go build -ldflags "                                     \
    -installsuffix 'static'                             \
    -s -w                                               \
    -X '$(go list -m)/pkg/version.tagStr=${TAG}'\
    -X '$(go list -m)/pkg/version.uptStr=${UPT}'\
    -X '$(go list -m)/pkg/version.envStr=${ENV}'\
    -X '$(go list -m)/pkg/version.authStr=${AUTH}'\
    " \
    -o $2 $1
