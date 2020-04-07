package docker_svc

import (
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/utils"
	"strings"
)

func SelectContainer() (*string, error) {
	err, out, errout := utils.ExecGetOutput(bash.DockerListContainersByName, nil)
	if err != nil {
		return nil, err
	}
	if errout != "" {
		return nil, errors.New("error getting containers, maybe docker isn't running")
	}
	trimmedOutput := strings.TrimSpace(out)
	splitOutput := strings.Split(trimmedOutput, "\n")
	if len(splitOutput) < 1 {
		return nil, errors.New("no containers found")
	}
	selected, err := fuzzyfinder.Find(splitOutput,
		func(i int) string {
			return splitOutput[i]
		})
	if err != nil {
		return nil, errors.New("no container selected, aborting")
	}
	return &splitOutput[selected], nil
}
