package cmd

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/config"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var WreckCmd = &cobra.Command{
	Use:     "wreck",
	Aliases: []string{"wreck"},
	Short:   "wreck docker",
	Long:    `stops docker and removes any unused images (all of them since it is stopped)`,
	Run:     WreckDocker,
}

func WreckDocker(_ *cobra.Command, _ []string) {
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

	if err := utils.Exec(bash.DockerWreck(), nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(WreckCmd)
}
