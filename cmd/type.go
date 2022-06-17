package cmd

import (
	"os/exec"
)

var (
	AdbPath = ""
)

type Command struct {
	cmd       *exec.Cmd
	exitCode  int
	errOutput string
}
