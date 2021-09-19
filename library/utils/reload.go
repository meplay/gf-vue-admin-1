package utils

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var Service = new(service)

type service struct{}

func (s *service) Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("系统不支持")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
