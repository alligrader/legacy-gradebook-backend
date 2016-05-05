package tasks

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/signatures"
	log "github.com/Sirupsen/logrus"
)

var cnf = config.Config{
	Broker:        "amqp://guest:guest@localhost:5672/",
	ResultBackend: "amqp://guest:guest@localhost:5672/",
	Exchange:      "machinery_exchange",
	ExchangeType:  "direct",
	DefaultQueue:  "machinery_tasks",
	BindingKey:    "machinery_task",
}

func ConnectServer() *machinery.Server {
	server, err := machinery.NewServer(&cnf)
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func AddTasks(s *machinery.Server) {
	s.RegisterTask("checkstyle", Checkstyle)
	s.RegisterTask("findbugs", Findbugs)
}

func Checkstyle(submissionID int) {
	// Fire off a docker container that contains the checkstlye code
	// with the repo copied from the filesystem
	// Take the results and put them into a results table that matches the submission ID to the output.
}

func Findbugs(submission int) {
	// Fire off a docker container that contains the findbugs code
	// with the repo copied from the filesystem
	// Take the results and put them into a results table that matches the submission ID to the output.

}

func SendTask() {

	server := ConnectServer()
	AddTasks(server)

	// SendTask no longer works!
	task := signatures.TaskSignature{
		Name: "add",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "int64",
				Value: 4,
			},
			signatures.TaskArg{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	_, err := server.SendTask(&task)
	if err != nil {
		// failed to send the task
		log.Fatal(err)
	}
}
