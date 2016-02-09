#!/bin/bash

function set_env {
    export GO15VENDOREXPERIMENT=1
    export SHAMAN_ENV='DEVELOPMENT'
}

function install_deps {
    sudo chown -R vagrant /opt/gopath
    go get github.com/Masterminds/glide
    glide up
}

function install_autoenv {
    git clone git://github.com/kennethreitz/autoenv.git ~/.autoenv
    echo 'source ~/.autoenv/activate.sh' >> ~/.bashrc
    source ~/.bashrc
    echo 'You might have to run source ./~bashrc before everything will work.'
}

function main {
    set_env
    install_deps
    install_autoenv
}

main
