package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var StashAllCmd = &cobra.Command{
	Use:     "stash_all",
	Aliases: []string{"sa"},
	Short:   "stash current branch and submodules",
	Long:    ``,
	Run:     StashAll,
}

func StashAll(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	command := bash.GitStash()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	command = bash.GitStashRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	GitCmds.AddCommand(StashAllCmd)
}
