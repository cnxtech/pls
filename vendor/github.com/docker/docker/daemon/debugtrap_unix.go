// +build !windows

package daemon

import (
	"os"
	"os/signal"
	"syscall"

	stackdump "github.com/docker/docker/pkg/signal"
	"github.com/sirupsen/logrus"
)

func (d *Daemon) setupDumpStackTrap(root string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			path, err := stackdump.DumpStacks(root)
			if err != nil {
				logrus.WithError(err).Error("failed to write goroutines dump")
			} else {
				logrus.Infof("goroutine stacks written to %s", path)
			}
		}
	}()
}
