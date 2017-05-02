package sigstat

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type Command struct {
	Command   []string
	Path      string
	Timeout   time.Duration
	StdOut    bytes.Buffer
	StdErr    bytes.Buffer
	StartedAt time.Time
	StoppedAt time.Time
	Killed    bool
	ExitCode  int
}

func (c *Command) Exec() {

	c.StartedAt = time.Now()

	cmd := exec.Command(c.Command[0], c.Command[1:]...)
	cmd.Dir = c.Path
	cmd.Stdout = &c.StdOut
	cmd.Stderr = &c.StdErr
	cmd.Env = os.Environ()
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

		err := cmd.Wait()
		if c.Timeout > 0 {
			timer.Stop()
		}
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
