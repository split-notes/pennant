package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/services/input_svc"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var CleanAllCmd = &cobra.Command{
	Use:     "clean_all",
	Aliases: []string{"ca"},
	Short:   "clean current branch and submodules",
	Long:    `Removes untracked files and directories in base repo and all sub repos`,
	Run:     CleanAll,
}

func CleanAll(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	input_svc.UserConfirms("Are you sure you want to do this? This will remove all untracked files. Recommend stashing first")
	command := bash.GitClean()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	command = bash.GitCleanRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	GitCmds.AddCommand(CleanAllCmd)
}
