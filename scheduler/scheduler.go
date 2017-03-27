package scheduler

import (
	"fmt"
	"io"

	"github.com/faststackco/machinestack/config"
	"github.com/gorilla/websocket"
)

type Scheduler interface {
	Create(name, image, driverName string) (string, error)
	Delete(name, driverName, node string) error
	Exec(name, driverName, node string, stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, controlHandler func(*websocket.Conn)) error
}

func NewScheduler(name string, options *config.DriverOptions) (Scheduler, error) {
	switch name {
	case "local":
		return NewLocalScheduler(options)
	case "consul":
		return NewConsulScheduler(options)
	default:
		return nil, fmt.Errorf("unknown scheduler '%s'", name)
	}
}
