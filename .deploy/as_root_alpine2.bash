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

function set_env {
    touch $HOME/.profile
    export PATH="/usr/local/go/bin":$PATH
    export GOROOT="/usr/local/go"
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-backend"
    export GOPATH="/opt/gopath"
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    echo "export GOPATH=/opt/gopath" >> $HOME/.profile
    echo "export GOPATH=/opt/gopath" >> $HOME/.bashrc
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.bashrc
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.bashrc
    echo "cd $APPDIR" >> $HOME/.bashrc
    source $HOME/.profile
}


function main {
    update
    install_deps
    make_gopath
    install_dredd
    set_env
}

main
echo '~~ Provisioning complete ~~'
