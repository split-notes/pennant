package cmd

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/config"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var StopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"so"},
	Short:   "stop docker",
	Long:    ``,
	Run:     StopDocker,
}

func StopDocker(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("docker"); err != nil {
		log.Println(err.Error())
		return
	}

	var configs = config.GetConfigFromViper()
	command := bash.DockerStop(configs.PennantConfig.ProjectFilePath)
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(StopCmd)
}
