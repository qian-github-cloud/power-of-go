package shell

import (
	"errors"
	"os/exec"
	"strings"
)

func CmdFromString(input string) (*exec.Cmd, error) {

	args := strings.Fields(input)
	if args < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil

}
