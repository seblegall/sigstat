package sigstat

import (
	"time"
	"bytes"
)

type Command struct {
	Command   []string
	Timeout   time.Duration
	StdOut    bytes.Buffer
	StdErr    bytes.Buffer
	StartedAt time.Time
	StoppedAt time.Time
	Killed    bool
	ExitCode  int
}

type CommandService interface {

}



