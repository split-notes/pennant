package golang

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/config/submodule_config"
	"github.com/split-notes/pennant/cli/utils"
	"log"
)

var TestCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"te"},
	Short:   "Test go submodule's services",
	Long:    `Searches submodules for go projects and runs 'go mod vendor'`,
	Run:     Test,
}

func Test(_ *cobra.Command, _ []string) {
	submodules, err := submodule_config.IdentifySubmodules()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, configData := range submodules {
		if configData.Language == "golang" {
			command := bash.GoTest()
			if err := utils.Exec(command, &configData.ProjectPath); err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func init() {
	GolangCmds.AddCommand(TestCmd)
}
