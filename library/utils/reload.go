package utils

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/pkg/errors"
)

var Server = new(server)

type server struct{}

func (s *server) Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("系统不支持")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
