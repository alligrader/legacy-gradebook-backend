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
    echo "cd $APPDIR" >> /home/vagrant/.bashrc
    source $HOME/.profile
}

function install_autoenv {
    git clone git://github.com/kennethreitz/autoenv.git ~/.autoenv
    echo 'source ~/.autoenv/activate.sh' >> ~/.profile
    source ~/.profile
}

function install_deps {
    go get github.com/snikch/goodman
    go get bitbucket.org/liamstask/goose/cmd/goose
    go get github.com/Masterminds/glide
    cd $APPDIR
    glide up
    cd -
}

function install_dredd {
    sudo apt-get install -yq npm
    sudo npm install npm -g
    sudo ln -s `which nodejs` /usr/bin/node
    sudo npm install -g dredd
}

function make_storage {
    mkdir /home/vagrant/storage
}

function main {
    set_env
    install_deps
    install_autoenv
    install_dredd
    make_storage
}

main
