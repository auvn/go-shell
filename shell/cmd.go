package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	baseCmd     string
	baseCmdArgs []string
	args        []string
}

func newShellCommand(args ...string) *Command {
	return &Command{
		baseCmd:     "/bin/sh",
		baseCmdArgs: []string{"-c"},
		args:        args,
	}
}

func (c *Command) Run(args ...string) (*CommandResult, error) {
	cmd := exec.Command(c.baseCmd, append(c.baseCmdArgs, c.cmdString(args...))...)

	var (
		stderr  bytes.Buffer
		stdout  bytes.Buffer
		success bool
	)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	ok, err := c.run(cmd)
	if err != nil {
		return nil, err
	}

	success = ok

	return &CommandResult{
		Success: success,
		Stderr:  stderr,
		Stdout:  stdout,
	}, nil

}

func (c *Command) run(cmd *exec.Cmd) (bool, error) {
	err := cmd.Run()
	if err != nil {
		switch exactErr := err.(type) {
		case *exec.ExitError:
			return exactErr.Success(), nil
		default:
			return false, err
		}
	}
	return true, nil
}

func (c *Command) cmdString(args ...string) string {
	cmdFields := append(c.args, args...)
	return strings.Join(cmdFields, " ")
}

func Cmd(args ...string) *Command {
	return newShellCommand(args...)
}

type RunFunc func(args ...string) (*CommandResult, error)

func (f RunFunc) Unsafe() UnsafeRunFunc {
	return func(args ...string) CommandResult {
		r, err := f(args...)
		if err != nil {
			panic(fmt.Sprintf("unsafe func: %v", err))
		}
		return *r
	}
}

type UnsafeRunFunc func(args ...string) CommandResult

func (f UnsafeRunFunc) Output() OutputOnlyRunFunc {
	return func(args ...string) string {
		r := f(args...)
		return r.AnyOutputString()
	}
}

type OutputOnlyRunFunc func(args ...string) string

func CmdFunc(args ...string) RunFunc {
	cmd := Cmd(args...)
	return cmd.Run
}

type CommandResult struct {
	Stderr  bytes.Buffer
	Stdout  bytes.Buffer
	Success bool
}

func (r *CommandResult) AnyOutputString() string {
	bb := r.anyOutput(r.Stderr, r.Stdout)
	return string(bb)
}

func (r *CommandResult) anyOutput(bufs ...bytes.Buffer) []byte {
	var bb []byte
	for i := range bufs {
		bb = bufs[i].Bytes()
		if len(bb) != 0 {
			break
		}
	}
	return bb
}
