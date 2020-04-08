package cmd

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/config/submodule_config"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
	"strings"
)

var OnboardEverythingCmd = &cobra.Command{
	Use:     "onboard",
	Aliases: []string{"on"},
	Short:   "setup all projects",
	Long:    `Calls the 'git onboard' command and then sets up repos. Node could run 'npm install' here`,
	Run:     OnboardEverything,
}

func OnboardEverything(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitOnboard()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	submodules, err := submodule_config.IdentifySubmodules(nil, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, configData := range submodules {
		if configData.Language == "golang" {
			if err := utils.Exec(bash.GoModVendor, &configData.ProjectPath); err != nil {
				log.Println(err.Error())
				return
			}
		}

		if strings.HasSuffix(configData.Language, "js") {
			if err := utils.Exec(bash.NpmInstall, &configData.ProjectPath); err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(OnboardEverythingCmd)
}
