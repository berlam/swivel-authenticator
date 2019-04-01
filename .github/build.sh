#!/bin/sh

cd ${GITHUB_WORKSPACE:-.}

mkdir -p .dist

echo "0.1.0" > .version

for GOOS in darwin linux windows; do
	SUFFIX=`[ $GOOS = "windows" ] && echo ".exe"`
	for GOARCH in 386 amd64; do
		for D in swivelp swivelt; do go build -v -o .release/$GOOS/$GOARCH/$D$SUFFIX cmd/$D/main.go; done
		tar --transform 's/.*\///g' -czvf .dist/swivel-authenticator-$GOOS-$GOARCH.tar.gz .release/$GOOS/$GOARCH/* README.md LICENSE
	done
done
