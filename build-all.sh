#!/bin/bash


BUILD_PLATFORMS="linux windows darwin freebsd"
BUILD_ARCHS="amd64 386"

# For each platform
for OS in ${BUILD_PLATFORMS[@]}; do

    # For each arch
    for ARCH in ${BUILD_ARCHS[@]}; do

        #
        # Windows binaries would be more useful if they had
        # an .EXE suffix.
        #
        SUFFIX=""
        if [ "$OS" = "windows" ]; then
            SUFFIX=".exe"
        fi

        # Get the dependencies
        GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 go get -t -v -d $(go list ./...)

        # Do the build
        GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 go build -ldflags "-X main.version=$(git describe --tags)" -o puppet-summary-${OS}-${ARCH}${SUFFIX} .
    done
done
