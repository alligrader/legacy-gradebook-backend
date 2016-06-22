#!/bin/bash

readonly ME=$(whoami)
readonly DIST='vivid'

function update {
    sudo apt-add-repository "deb http://download.virtualbox.org/virtualbox/debian $DIST contrib"
    wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -
    sudo apt-get update
}

function install_deps {
    sudo apt-get install -yq \
        git \
        mysql-client \
        curl \
        wget \
        zip \
        unzip \
        vim-nox \ 
        npmi \
        virtualbox-5.0
}

function make_gopath {
    mkdir -p "/opt/gopath/src" | 0
    mkdir -p "/opt/gopath/bin" | 0
    mkdir -p "/opt/gopath/pkg" | 0
    sudo chown -R $ME:$ME /opt  # Work around bug in Vagrant
}

function install_go {
    if [ ! -d /usr/local/go ]; then
        # curl -O -J -L "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
        wget -q "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
        sudo tar -C /usr/local -xvf 'go1.7beta1.linux-amd64.tar.gz'
    fi
}

function install_packer {
    if [ ! -d /usr/local/go ]; then
        # curl -O -J -L "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
        wget -q 'https://releases.hashicorp.com/packer/0.10.1/packer_0.10.1_linux_amd64.zip'
        sudo unzip -d /usr/local 'packer_0.10.1_linux_amd64.zip'
    fi
}

function install_dredd {
    sudo npm install -g npm@latest
    sudo ln -s `which nodejs` /usr/bin/node
    sudo npm install -g dredd@latest
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
