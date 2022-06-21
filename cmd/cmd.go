package cmd

import (
	"adb/tool"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

var (
	AdbPath = ""
)

type Command struct {
	cmd       *exec.Cmd
	exitCode  int
	errOutput string
}

func getCmd(arg ...string) *exec.Cmd {
	if runtime.GOOS == "linux" {
		return exec.Command("bash", "-c", strings.Join(arg, " "))
	} else if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", strings.Join(arg, " "))
	}
	return nil
}

func New(arg ...string) *Command {
	args := append([]string{AdbPath}, arg...)
	cmd := getCmd(args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return &Command{
		cmd: cmd,
	}
}

func NewScrcpy(arg ...string) *Command {
	_, err := os.Stat("scrcpy/scrcpy.exe")
	if err == nil {
		scrcpy, _ := filepath.Abs("scrcpy/scrcpy.exe")
		args := append([]string{scrcpy}, arg...)
		cmd := getCmd(args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return &Command{
			cmd: cmd,
		}
	}
	tool.Error(err, 1)
	return nil
}

func (c *Command) SyncRun() error {
	err := c.cmd.Run()
	tool.Error(err, 1)
	return err
}

func (c *Command) SyncReadString() string {
	byt, err := c.cmd.CombinedOutput()
	if err != nil {
		tool.Error(err, 1)
		return err.Error()
	}
	return tool.ConvertByte2String(byt)
}

func (c *Command) String() string {
	return c.cmd.String()
}

func (c *Command) Start(cb func(line string, err error)) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	if c.cmd != nil {
		stdout, err := c.cmd.StdoutPipe()
		if err != nil {
			c.exitCode = -998
			c.errOutput += err.Error()
			return
		}
		errout, err := c.cmd.StderrPipe()
		if err != nil {
			c.exitCode = -998
			c.errOutput += err.Error()
			return
		}
		go func() {
			defer wg.Done()
			c.read(stdout, cb)
		}()
		go func() {
			defer wg.Done()
			c.read(errout, cb)
		}()
		err = c.cmd.Start()
		if err != nil {
			c.exitCode = -997
			c.errOutput += err.Error()
			return
		}
		wg.Wait()
		_ = c.cmd.Wait()
		c.exitCode = c.cmd.ProcessState.ExitCode()
	}
}

func (c *Command) read(in io.ReadCloser, cb func(line string, err error)) {
	reader := bufio.NewReader(in)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			cb("", err)
			break
		}
		if runtime.GOOS == "windows" {
			line = tool.ConvertByte2String([]byte(line))
		}
		cb(line, nil)
	}
}

func (c *Command) Quit() error {
	var err error
	if c.cmd.Process != nil {
		if runtime.GOOS == "linux" {
			err = c.cmd.Process.Kill()
		}
		if runtime.GOOS == "windows" {
			maxCloseC := exec.Command("taskkill.exe", "/pid", fmt.Sprintf("%d", c.cmd.Process.Pid), "-t", "-f")
			err = maxCloseC.Run()
		}
	}
	tool.Error(err, 1)
	return err
}
