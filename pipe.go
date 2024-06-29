package core

// copyright https://github.com/babarot/go-pipe

import (
	"bytes"
	"io"
	"os/exec"
)

// Command is a os/exec.Commnad wrapper for UNIX pipe
func Command(stdout *bytes.Buffer, stack ...*exec.Cmd) (err error) {
	var stderr bytes.Buffer
	pipeStack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		inPipe, outPipe := io.Pipe()
		stack[i].Stdout = outPipe
		stack[i].Stderr = &stderr
		stack[i+1].Stdin = inPipe
		pipeStack[i] = outPipe
	}
	stack[i].Stdout = stdout
	stack[i].Stderr = &stderr

	return call(stack, pipeStack)
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}
