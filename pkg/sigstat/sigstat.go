package sigstat

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

//Command represents a shell command.
//It could be a binary as well as a script
//or any kind of command it is possible to run in a shell.
type Command struct {
	Command   []string
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
	UpdateStatus(cmd *Command)
}

//Exec actualy exec the command and send status update using
//the CommandService defined in a given service.
func (c *Command) Exec(client Client) {

	c.StartedAt = time.Now()

	cmd := exec.Command("/bin/sh", "-c", strings.Join(c.Command, " "))
	cmd.Dir = c.Path
	cmd.Stdout = &c.StdOut
	cmd.Stderr = &c.StdErr
	cmd.Env = os.Environ()

	//Start process. Exit code 127 if process fail to start.
	if err := cmd.Start(); err != nil {
		c.StdErr.WriteString("\n" + err.Error() + "\n")
		c.ExitCode = 127
	} else {
		// Create a ticker that outputs elapsed time
		ticker := time.NewTicker(time.Millisecond * 100)
		go func(ticker *time.Ticker, cmd *exec.Cmd, c *Command, client Client) {
			for _ = range ticker.C {
				err := cmd.Process.Signal(syscall.Signal(0))
				if err == nil {
					c.Status = "running"
					client.CommandService().UpdateStatus(c)
				}
			}
		}(ticker, cmd, c, client)

		err := cmd.Wait()

		ticker.Stop()

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
