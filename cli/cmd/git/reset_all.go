package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/services/input_svc"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var ResetAllCmd = &cobra.Command{
	Use:     "reset_all",
	Aliases: []string{"ra"},
	Short:   "reset current branch and submodules",
	Long:    `Resets all tracked files in base repo and all sub repos to last commit`,
	Run:     ResetAll,
}

func ResetAll(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	input_svc.UserConfirms("Are you sure you want to do this? This will reset all tracked files in project. Recommend first stashing changes")

	command := bash.GitResetHard()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	command = bash.GitResetHardRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	GitCmds.AddCommand(ResetAllCmd)
}
