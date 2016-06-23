#!/bin/bash

readonly ME=$(whoami)
readonly DIST='vivid'

function update {
    apk update
}

function install_deps {
    echo '~~ Installing via apk ~~'
    apk add \
        zip \
        unzip \
        vim \
        python \
        python-dev \
        py-pip \
        alpine-sdk \
        nodejs
}

function make_gopath {
    echo '~~ Preparing the gopath ~~'
    mkdir -p "/opt/gopath/src" || 0
    mkdir -p "/opt/gopath/bin" || 0
    mkdir -p "/opt/gopath/pkg" || 0
}

function install_dredd {
    echo '~~ Installing dredd ~~'
    sudo npm install -g npm@latest
    sudo ln -s $(which nodejs) /usr/bin/node
    sudo npm install -g dredd@stable
    echo '~~ dredd installed successfully ~~'
}

function main {
    update
    install_deps
    make_gopath
    install_dredd
}

main
echo '~~ Provisioning complete ~~'
