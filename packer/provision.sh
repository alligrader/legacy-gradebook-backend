#!/bin/bash
set -e
echo 'Hello Alpine!'
adduser -S alligrader
apk update
apk add alpine-sdk
apk add --update python python-dev py-pip
apk add --update nodejs
apk add git bzr mercurial
sudo npm install -g npm@latest
sudo npm install -g dredd@latest

oe() { $@ 2>&1 | logger -t otto > /dev/null; }
ol() { echo "[otto] $@"; }

# If we have Go, then do nothing
if command -v go >/dev/null 2>&1; then
    ol "Go already installed! Otto won't install Go."
    exit 0
fi

oe mkdir -p "/opt/gopath/src"
oe mkdir -p "/opt/gopath/bin"
oe mkdir -p "/opt/gopath/pkg"

ol "Downloading Go 1.7 beta 2..."
oe wget -q -O /home/alligrader/go.tar.gz https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz
ol "Untarring Go..."
oe sudo tar -C /usr/local -xzf /home/alligrader/go.tar.gz
ol "Setting up PATH..."
echo 'export PATH=/opt/gopath/bin:/usr/local/go/bin:$PATH' >> /home/alligrader/.profile
echo 'export GOPATH=/opt/gopath' >> /home/alligrader/.profile
ol "Configuring Go to use SSH instead of HTTP..."
git config --global url."git@github.com:".insteadOf "https://github.com/"
