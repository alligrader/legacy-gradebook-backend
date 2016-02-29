If you want to connect to the database from the command line, you need to tell the MySQL client to use `127.0.0.1` instead of localhost.

	mysql -h 127.0.0.1 -P 3306 -u root -p

The password is `root`. Check out the username and password settings in `.env`, and specified in the `Vagrantfile`.

To dump the database contents out into a file, or make a barebones schema:

    mysqldump --compatible=ansi --add-drop-database --skip-comments --add-drop-table --disable-keys --skip-set-charset --host=127.0.0.1 -u root --password=root shaman > mysqldump.sql
