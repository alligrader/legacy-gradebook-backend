#!/bin/bash
set -e
echo 'Hello Alpine!'
apk update
apk add alpine-sdk
apk add --update python python-dev py-pip
apk add --update nodejs
sudo npm install -g npm@latest
sudo npm install -g dredd@latest
