package util

import (
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func ShellOut(config *DBConfig) {

	var (
		name = config.Flavor
		args []string
	)

	if name == "mysql" {
		args = []string{"-h", "127.0.0.1", "-P", "3306", "-u", "root", "--password=root"}
	} else {
		log.Fatal("Not really sure how to build this, boss.")
	}

	schema, err := os.Open(config.SchemaFile)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = schema
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
