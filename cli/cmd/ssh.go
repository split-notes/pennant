package cmd

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/services/docker_svc"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var SSHCmd = &cobra.Command{
	Use:     "ssh",
	Aliases: []string{"sh"},
	Short:   "docker ssh",
	Long:    `ssh into a docker container of your choosing`,
	Run:     DockerSSH,
}

func DockerSSH(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("docker"); err != nil {
		log.Println(err.Error())
		return
	}

	container, err := docker_svc.SelectContainer()
	if err != nil {
		log.Println(err.Error())
		return
	}

	shell, err := docker_svc.GetChoiceShell(*container)
	if err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.DockerSSH(*container, shell)
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(SSHCmd)
}
