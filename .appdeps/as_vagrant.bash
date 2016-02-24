touch $HOME/.profile

export APPDIR="/opt/gopath/src/github.com/gradesham/gradebook-backend"
export GOPATH="/opt/gopath"
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

echo "export GOPATH=/opt/gopath" >> $HOME/.profile
echo "export GOBIN=\$GOPATH/bin" >> $HOME/.profile
echo "export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin" >> $HOME/.profile
source $HOME/.profile

go get github.com/Masterminds/glide
cd $APPDIR
glide up
cd -

function install_autoenv {
    git clone git://github.com/kennethreitz/autoenv.git ~/.autoenv
    echo 'source ~/.autoenv/activate.sh' >> ~/.profile
    source ~/.profile
}
install_autoenv
