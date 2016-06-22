#!/bin/bash

readonly ME=$(whoami)
touch $HOME/.profile

function set_env {
    export APPDIR="/opt/gopath/src/github.com/alligrader/gradebook-backend"
    export GOPATH="/opt/gopath"
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    echo "export GOPATH=/opt/gopath" >> $HOME/.profile
    echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
    echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
    echo 'export GO15VENDOREXPERIMENT=1' >> ~/.profile
    echo "cd $APPDIR" >> $HOME/.bashrc
    source $HOME/.profile
}

function install_godeps {
    go get github.com/snikch/goodman
    go get bitbucket.org/liamstask/goose/cmd/goose
    go get github.com/Masterminds/glide
    cd $APPDIR
    glide up
    cd -
}

function install_autoenv {
    git clone git://github.com/kennethreitz/autoenv.git ~/.autoenv
    echo 'source ~/.autoenv/activate.sh' >> ~/.profile
    source ~/.profile
}

function make_storage {
    mkdir $HOME/storage
}

function main {
    set_env
    install_autoenv
    install_godeps
    make_storage
}

main
