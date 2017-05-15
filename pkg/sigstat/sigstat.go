package sigstat

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

//Command represents a shell command.
//It could be a binary as well as a script
//or any kind of command it is possible to run in a shell.
type Command struct {
	Group     string
	GroupID   int64
	ID        int64
	Command   string
	Path      string
	Timeout   time.Duration
	Status    string
	StdOut    bytes.Buffer
	StdErr    bytes.Buffer
	StartedAt time.Time
	StoppedAt time.Time
	Killed    bool
	ExitCode  int
}

//Client is an interface that must be implemented by any kind of sigstat client.
//Client could be an http client, a tcp client or event a io.writer that log informations.
type Client interface {
	CommandService() CommandService
}

//CommandService define the way command should me manipulated throw the client.
type CommandService interface {
	CreateCommand(cmd Command) (int64, error)
	UpdateStatus(cmd Command)
}

//Exec actualy exec the command and send status update using
//the CommandService defined in a given service.
func (c *Command) Exec(client Client) {

	c.StartedAt = time.Now()

	cmd := exec.Command("/bin/sh", "-c", c.Command)
	cmd.Dir = c.Path
	cmd.Stdout = &c.StdOut
	cmd.Stderr = &c.StdErr
	cmd.Env = os.Environ()

	//Start process. Exit code 127 if process fail to start.
	if err := cmd.Start(); err != nil {
		c.StdErr.WriteString("\n" + err.Error() + "\n")
		c.ExitCode = 127
	} else {
		var timer *time.Timer
		if c.Timeout > 0 {
			timer = time.NewTimer(c.Timeout)
			go func(timer *time.Timer, cmd *exec.Cmd) {
				for _ = range timer.C {
					c.Killed = true
					if err := cmd.Process.Kill(); err != nil {
						c.StdErr.WriteString(fmt.Sprintf("\nUnabled to kill the process: %s\n", err))
					}
				}
			}(timer, cmd)
		}

		// Create a ticker that will let us send status information at regular time duration.
		ticker := time.NewTicker(time.Millisecond)
		go func(ticker *time.Ticker, cmd *exec.Cmd, c *Command, client Client) {
			for _ = range ticker.C {
				err := cmd.Process.Signal(syscall.Signal(0))
				if err == nil {
					c.Status = "running"
					client.CommandService().UpdateStatus(*c)
				}
			}
		}(ticker, cmd, c, client)

		err := cmd.Wait()

		ticker.Stop()

		if c.Timeout > 0 {
			timer.Stop()
		}

		c.Status = "stopped"

		if err != nil {
			// unsuccessful exit code?
			c.ExitCode = -1
			if exitError, ok := err.(*exec.ExitError); ok {
				c.ExitCode = exitError.Sys().(syscall.WaitStatus).ExitStatus()
			}
		}
	}

	c.StartedAt = time.Now()
}
