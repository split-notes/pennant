package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var FetchPullCmd = &cobra.Command{
	Use:     "fetch_pull",
	Aliases: []string{"fp"},
	Short:   "fetch and pull current branch and submodules",
	Long:    `git fetch --recurse-submodules && git pull --recurse-submodules`,
	Run:     FetchPull,
}

func FetchPull(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitFetchRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	command = bash.GitPullRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	GitCmds.AddCommand(FetchPullCmd)
}
