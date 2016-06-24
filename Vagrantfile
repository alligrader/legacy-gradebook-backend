# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = "rem89/alligrader"
  config.vm.network "forwarded_port", guest: 443, host: 443
  config.vm.network "forwarded_port", guest: 8000, host: 8000
  config.vm.network "forwarded_port", guest: 5000, host: 5000
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  # config.vm.network "forwarded_port", guest: 3306, host: 3306
  config.vm.network "forwarded_port", guest: 5672, host: 5672

  # removes the default shared folder.
  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.synced_folder ".", "/opt/gopath/src/github.com/alligrader/gradebook-backend"

  config.ssh.forward_agent = true
  config.vm.provider "virtualbox" do |v|
    v.memory = 4096 # this line is required to make the box large enough for all the docker deps
    v.cpus = 2
  end

  # config.vm.provision "shell", path: ".appdeps/provision.bash"
  # config.vm.provision "shell", path: ".deploy/install_go.bash"
  config.vm.provision "shell", path: ".deploy/as_user.bash", privileged: false
  #config.vm.provision "docker" do |d|
  #  d.run "mysql",  image: "mysql",    args: "-p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=root -e MYSQL_DATABASE=alligrader"
  #  d.run "rabbit", image: "rabbitmq:3.6.0-management", args: "-p 8080:15672 -p 5672:5672"
  #end

  # Define a Vagrant Push strategy for pushing to Atlas. Other push strategies
  # such as FTP and Heroku are also available. See the documentation at
  # https://docs.vagrantup.com/v2/push/atlas.html for more information.
  # config.push.define "atlas" do |push|
  #   push.app = "YOUR_ATLAS_USERNAME/YOUR_APPLICATION_NAME"
  # end

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", path: ".deploy/provision.sh"
end
