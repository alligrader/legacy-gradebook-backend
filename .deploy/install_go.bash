mkdir -p "/opt/gopath/src" | 0
mkdir -p "/opt/gopath/bin"
mkdir -p "/opt/gopath/pkg"

sudo chown -R vagrant:vagrant /opt  # Work around bug in Vagrant
sudo apt-get update
sudo apt-get install -yq \
    git \
    mysql-client \
    curl \
    vim-nox

if [ ! -d /usr/local/go ]; then
    curl -O -J -L "https://storage.googleapis.com/golang/go1.7beta2.linux-amd64.tar.gz"
    # wget -q "https://storage.googleapis.com/golang/go1.7beta1.linux-amd64.tar.gz"
    sudo tar -C /usr/local -xvf go1.7beta2.linux-amd64.tar.gz
fi
