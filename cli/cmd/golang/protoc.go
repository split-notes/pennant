package golang

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/services/git_svc"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os"
	"os/exec"
)

var ProtocCmd = &cobra.Command{
	Use:     "protoc",
	Aliases: []string{"po"},
	Short:   "Run Protoc go submodule's servers",
	Long:    `Searches submodules for go projects and runs Protoc on each of their servers`,
	Run:     Protoc,
}

func Protoc(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("protoc"); err != nil {
		log.Println(err.Error())
		return
	}

	transportFilter := "grpc"
	languageFilter := "golang"
	submodules, err := git_svc.SelectSubmodules(&languageFilter, &transportFilter)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, configData := range submodules {
		if  configData.Transport == "grpc" && configData.Language == "golang"{
			f, err := os.Open(fmt.Sprintf("%s/servers", configData.ProjectPath))
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
					command := bash.GoProtoc(fmt.Sprintf("./servers/%s", file.Name()))
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
	GolangCmds.AddCommand(ProtocCmd)
}
