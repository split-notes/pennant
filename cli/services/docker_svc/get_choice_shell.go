package docker_svc

import (
	"errors"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/utils"
	"strings"
)

func GetChoiceShell(container string) (string, error) {
	err, out, errout := utils.ExecGetOutput(bash.DockerListShellsOfContainer(container), nil)
	if err != nil {
		return "", err
	}
	if errout != "" && errout != "Unable to use a TTY - input is not a terminal or the right kind of file\n" {
		return "", errors.New("error getting containers, maybe docker hasn't been started yet")
	}
	trimmedOutput := strings.TrimSpace(out)
	splitOutput := strings.Split(trimmedOutput, "\n")
	if len(splitOutput) < 1 {
		return "", errors.New("no pods found")
	}

	for _, v := range splitOutput {
		if v == "/bin/bash" {
			return "bash", nil
		}
	}
	return "sh", nil
}
