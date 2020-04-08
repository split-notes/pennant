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
	"strings"
)

var MockCmd = &cobra.Command{
	Use:     "mock",
	Aliases: []string{"mo"},
	Short:   "Mock go submodule's services",
	Long:    `Searches submodules for go projects and mocks their services`,
	Run:     Mock,
}

func Mock(_ *cobra.Command, _ []string) {

	languageFilter := "golang"
	submodules, err := git_svc.SelectSubmodules(&languageFilter, nil)
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
			servicesDir := fmt.Sprintf("%s/services", configData.ProjectPath)
			f, err := os.Open(servicesDir)
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
					if strings.HasPrefix(file.Name(), "grpc") {
						grpcF, err := os.Open(fmt.Sprintf("%s/%s", servicesDir, file.Name()))
						if err != nil {
							log.Println(err.Error())
							return
						}
						grpcFileInfo, err := grpcF.Readdir(-1)
						grpcF.Close()

						for _, grpcFile := range grpcFileInfo {
							if strings.HasSuffix(grpcFile.Name(), ".pb.go"){
								mockName := strings.TrimSuffix(grpcFile.Name(), ".pb.go")

								command := bash.GoMockReflect(configData.ProjectName, mockName)
								if err := utils.Exec(command, &configData.ProjectPath); err != nil {
									log.Println(err.Error())
									return
								}
							}
						}
					} else {
						command := bash.GoMockSource(file.Name())
						if err := utils.Exec(command, &configData.ProjectPath); err != nil {
							log.Println(err.Error())
							return
						}
					}
				}
			}
		}
	}
}

func init() {
	GolangCmds.AddCommand(MockCmd)
}
