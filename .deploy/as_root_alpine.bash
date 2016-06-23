#!/bin/bash
set -e
echo 'Hello Alpine Linux!'

function update {
    apk update
}

function install_deps {
    apk add alpine-sdk
    apk add \
        nodejs \
        git \
        bzr \
        mercurial \
        wget
}

function make_gopath {
    echo '~~ Preparing the gopath ~~'
    mkdir -p "/opt/gopath/src" || 0
    mkdir -p "/opt/gopath/bin" || 0
    mkdir -p "/opt/gopath/pkg" || 0
}

function install_go {

    local readonly GO_BIN_VERSION="go1.7beta2"
    local readonly URL="storage.googleapis.com/golang/${GO_BIN_VERSION}.linux-amd64.tar.gz"

    echo "~~ Installing $GO_BIN_VERSION ~~"
    if [ ! -d /usr/local/go ]; then
        echo 'About to wget'
        wget -q "http://$URL"
        echo 'About to untar'
        sudo tar -C /usr/local -xvzf "${GO_BIN_VERSION}.linux-amd64.tar.gz"
        echo 'about to rm the package'
        rm "${GO_BIN_VERSION}.linux-amd64.tar.gz"
    fi
    echo "~~ $GO_BIN_VERSION Installation complete ~~"
}

function install_dredd {
    echo '~~ Installing dredd ~~'
    sudo npm install -g npm@latest
    sudo ln -s $(which nodejs) /usr/bin/node
    sudo npm install -g dredd@stable
    echo '~~ dredd installed successfully ~~'
}

readonly ME=$(whoami)
touch $HOME/.profile

function set_env {
    export PATH="/usr/local/go/bin":$PATH
    export GOROOT="/usr/local/go"
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-backend"
    export GOPATH="/opt/gopath"
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    echo "export PATH=/usr/local/go/bin:\$PATH"
    echo "export GOROOT=/usr/local/go"
    echo "export GOPATH=/opt/gopath" >> $HOME/.profile
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
    echo "cd $APPDIR" >> $HOME/.bashrc
    source $HOME/.profile
}

function main {
    update
    install_deps
    make_gopath
    install_go
    set_env
    install_dredd
}

main
