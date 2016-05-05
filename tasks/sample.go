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

// Make the connection a shared global
func PushFindbugs(submissionID int) {
	s := ConnectServer()

	task := signatures.TaskSignature{
		Name: "findbugs",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "int",
				Value: submissionID,
			},
		},
	}

	_, err := s.SendTask(&task)
	if err != nil {
		// failed to send the task
		log.Error(err)
	}
}

func PushCheckstyle(submissionID int) {
	s := ConnectServer()

	task := signatures.TaskSignature{
		Name: "checkstyle",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "int",
				Value: submissionID,
			},
		},
	}

	_, err := s.SendTask(&task)
	if err != nil {
		// failed to send the task
		log.Error(err)
	}
}

func Checkstyle(submissionID int) {
	// Fire off a docker container that contains the checkstlye code
	// with the repo copied from the filesystem
	// Take the results and put them into a results table that matches the submission ID to the output.
	log.Info("Processing Checkstyle")
}

func Findbugs(submissionID int) {
	// Fire off a docker container that contains the findbugs code
	// with the repo copied from the filesystem
	// Take the results and put them into a results table that matches the submission ID to the output.
	log.Info("Processing Findbugs")

}
