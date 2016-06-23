# Vagrant

I'm assuming you have basic knowledge of Vagrant. Install Vagrant, then run `vagrant up` to start the VM. If you're shared folders aren't mounting because there's a problem with your Guest Additions, run `vagrant plugin install vagrant-vbguest`. That should solve your problem.

# MySQL

If you want to connect to the database from the command line, you need to tell the MySQL client to use `127.0.0.1` instead of localhost.

	mysql -h 127.0.0.1 -P 3306 -u root -p

The password is `root`. Check out the username and password settings in `.env`, and specified in the `Vagrantfile`.

Copy and paste that command from inside the Vagrant box to connect to MySQL. Then run `use alligrader` to switch to the right database. If you want to see the tables that are created, run `show tables;`. Easy enough, right? If you fucked up something in your database and want a fresh start, run `drop database alligrader; create database alligrader;`.

There's also a command line tool `goose` for switching to the right migration. `goose up` will apply all migrations. `goose down` will unmigrations (or the last migration? I'm not sure. See the goose documentation for more info).

To dump the database contents out into a file, or make a barebones schema:

    mysqldump --compatible=ansi --add-drop-database --skip-comments --add-drop-table --disable-keys --skip-set-charset --host=127.0.0.1 -u root --password=root shaman > mysqldump.sql

# Golang

To run the tests, execute `go test ./spec/...`. The `...` is a special syntax for recursively running tests.

To start the webserve, run `go run main.go`.
To build the binary, run `go build main.go`.

To run the Dredd tests, run `go build main.go` and then run `dredd`. Before submitting a PR, your commit should pass both the `dredd` tests and the `go test` tests. 
