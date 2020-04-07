package golang

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/config/submodule_config"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os"
	"os/exec"
)

var MockCmd = &cobra.Command{
	Use:     "mock",
	Aliases: []string{"mo"},
	Short:   "Mock go submodule's services",
	Long:    `Searches submodules for go projects and mocks their services`,
	Run:     Mock,
}

func Mock(_ *cobra.Command, _ []string) {
	submodules, err := submodule_config.IdentifySubmodules()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, configData := range submodules {
		// Check for necessary stuff
		if _, err := exec.LookPath("mockgen"); err != nil {
			log.Println(err.Error())
			return
		}

		if  configData.Language == "golang"{
			f, err := os.Open(fmt.Sprintf("%s/services", configData.ProjectPath))
			if err != nil {
				log.Println(err.Error())
				return
			}
			fileInfo, err := f.Readdir(-1)
			f.Close()
			if err != nil {
				log.Println(err.Error())
				return
			}

			for _, file := range fileInfo {
				if file.IsDir() {
					command := bash.GoMock(file.Name())
					if err := utils.Exec(command, &configData.ProjectPath); err != nil {
						log.Println(err.Error())
						return
					}
				}
			}
		}
	}
}

func init() {
	GolangCmds.AddCommand(MockCmd)
}
