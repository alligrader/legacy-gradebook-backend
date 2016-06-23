#!/bin/bash

readonly ME=$(whoami)
readonly DIST='vivid'

function update {
    echo '~~ Updating apt-get ~~'
    sudo apt-add-repository "deb http://download.virtualbox.org/virtualbox/debian $DIST contrib"
    wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -
    sudo apt-get update
}

function install_deps {
    echo '~~ Installing via apt-get ~~'
    sudo apt-get install -yq \
        git \
        mysql-client \
        curl \
        wget \
        zip \
        unzip \
        vim-nox \
        npm
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
        # curl -O -J -L "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
        wget -q "https://$URL"
        sudo tar -C /usr/local -xvf "${GO_BIN_VERSION}.linux-amd64.tar.gz"
        rm "${GO_BIN_VERSION}.linux-amd64.tar.gz"
    fi
    echo "~~ $GO_BIN_VERSION Installation complete ~~"
}

function install_packer {
    local readonly VERSION="0.10.1"
    local readonly PACKER_VERSION="packer_${VERSION}"
    local readonly URL="releases.hashicorp.com/packer/${VERSION}/${PACKER_VERSION}_linux_amd64.zip"
    echo "~~ Installing ${PACKER_VERSION} ~~"

    if [ ! -d /usr/local/packer ]; then
        wget -q "https://$URL"
        sudo unzip -d /usr/local/bin "${PACKER_VERSION}_linux_amd64.zip"
        rm "${PACKER_VERSION}_linux_amd64.zip"
    fi
    echo "~~ ${PACKER_VERSION} installed ~~"
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
    install_go
    install_packer
    install_dredd
}

main
echo '~~ Provisioning complete ~~'
