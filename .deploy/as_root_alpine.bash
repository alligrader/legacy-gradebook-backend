#!/bin/bash
set -e
echo 'Hello Alpine Linux!'

touch $HOME/.profile
readonly ME=$(whoami)
readonly GOLANG_VERSION='1.7beta2'
readonly GOLANG_SRC_URL="http://golang.org/dl/go$GOLANG_VERSION.src.tar.gz"
readonly GOLANG_SRC_SHA256='88840e78905bdff7c8e408385182b4f77e8bdd062cac5c0c6382630588d426c7'

function install_deps {
    apk add --update --virtual \
        .build-deps \
        bash \
        alpine-sdk \
        ca-certificates \
        musl-dev \
        openssl \
        go \
        nodejs \
        git \
        bzr \
        mercurial \
        wget
    echo 'Completing installing deps'
}

function build_go {
    echo 'Building Go'
    export GOROOT_BOOTSTRAP="$(which go)"
    echo 'Wgetting'
    wget -q --no-check-certificate "$GOLANG_SRC_URL" -O golang.tar.gz
    echo 'Checking the checksum'
    echo "$GOLANG_SRC_SHA256  golang.tar.gz" | sha256sum -c -
    tar -C /usr/local -xzf golang.tar.gz
    rm golang.tar.gz
    cd /usr/local/go/src
    ./make.bash
    apk del .build-deps
    cd -
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
        echo 'About to rm the package'
        rm "${GO_BIN_VERSION}.linux-amd64.tar.gz"
	cd /usr/local/go/src
	./all.bash
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

function set_env {
    export PATH="/usr/local/go/bin":$PATH
    export GOROOT="/usr/local/go"
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-backend"
    export GOPATH="/opt/gopath"
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    echo "export GOROOT=/usr/local/go" >> $HOME/.profile
    echo "export GOROOT=/usr/local/go" >> $HOME/.bashrc
    echo "export GOPATH=/opt/gopath" >> $HOME/.profile
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.bashrc
    echo "cd $APPDIR" >> $HOME/.bashrc
    source $HOME/.profile
}

function main {
    install_deps
    build_go
    #make_gopath
    #install_go
    #set_env
    #install_dredd
}

main
