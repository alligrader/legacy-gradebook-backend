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

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func Multiply(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	return sum, nil
}

func AddTasks(s *machinery.Server) {
	s.RegisterTask("add", Add)
	s.RegisterTask("multiply", Multiply)
}

func SendTask() {

	server := ConnectServer()
	AddTasks(server)

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
