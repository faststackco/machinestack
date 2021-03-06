package scheduler

import (
	"io"

	"gitlab.com/faststack/machinestack/config"
	"gitlab.com/faststack/machinestack/driver"
)

// LocalScheduler runs all machine on localhost
type LocalScheduler struct {
	driverOptions *config.DriverOptions
}

// NewLocalScheduler creates a new LocalScheduler
func NewLocalScheduler(options *config.DriverOptions) (Scheduler, error) {
	return &LocalScheduler{
		driverOptions: options,
	}, nil
}

// Create creates a new machine
func (c *LocalScheduler) Create(name, image, driverName string, attrs driver.MachineAttributes) (string, error) {
	driver, err := driver.NewDriver(driverName, *c.driverOptions)
	if err != nil {
		return "", err
	}

	if err := driver.Create(name, image, attrs); err != nil {
		return "", err
	}

	return "", nil
}

// Delete deletes a machine
func (c *LocalScheduler) Delete(name, driverName, node string) error {

	driver, err := driver.NewDriver(driverName, *c.driverOptions)
	if err != nil {
		return err
	}

	if err := driver.Delete(name); err != nil {
		return err
	}

	return nil
}

// Session creates an new exec session
func (c *LocalScheduler) Session(name, driverName, node string, stdin io.ReadCloser, stdout io.WriteCloser, control chan driver.ControlMessage, width, height int) error {
	driver, err := driver.NewDriver(driverName, *c.driverOptions)
	if err != nil {
		return err
	}

	return driver.Session(name, stdin, stdout, control, width, height)
}
