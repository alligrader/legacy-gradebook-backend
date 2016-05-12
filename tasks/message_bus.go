package tasks

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/signatures"
	log "github.com/Sirupsen/logrus"
)

func init() {
	Bus = &MessageBus{ConnectServer()}
	Bus.AddTasks()
}

// TODO get this configuration from the environment
// TODO swap the default queue with redis or something ?
var (
	cnf = config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		ResultBackend: "amqp://guest:guest@localhost:5672/",
		Exchange:      "machinery_exchange",
		ExchangeType:  "direct",
		DefaultQueue:  "machinery_tasks",
		BindingKey:    "machinery_task",
	}

	Bus *MessageBus
)

type MessageBus struct {
	*machinery.Server
}

func ConnectServer() *machinery.Server {
	server, err := machinery.NewServer(&cnf)
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func (bus *MessageBus) AddTasks() {
	bus.RegisterTask("checkstyle", Checkstyle)
	bus.RegisterTask("findbugs", Findbugs)
}

// Make the connection a shared global
func (bus *MessageBus) PushFindbugs(submissionID int) {

	task := signatures.TaskSignature{
		Name: "findbugs",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "int",
				Value: submissionID,
			},
		},
	}

	_, err := bus.SendTask(&task)
	if err != nil {
		// failed to send the task
		log.Error(err)
	}
}

func (bus *MessageBus) PushCheckstyle(submissionID int) {

	task := signatures.TaskSignature{
		Name: "checkstyle",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "int",
				Value: submissionID,
			},
		},
	}

	_, err := bus.SendTask(&task)
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
