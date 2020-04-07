package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/services/input_svc"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var WreckCmd = &cobra.Command{
	Use:     "wreck",
	Aliases: []string{"w"},
	Short:   "runs the `Reset` and `Clean` command",
	Long:    `Resets all tracked files and removes all untracked files`,
	Run:     Wreck,
}

func Wreck(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("git"); err != nil {
		log.Println(err.Error())
		return
	}

	input_svc.UserConfirms("Are you sure you want to do this? This will reset all tracked files in project. Recommend first stashing changes")

	// RESET
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

	// CLEAN
	command = bash.GitClean()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
	command = bash.GitCleanRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	// SWITCH TO MASTER
	command = bash.GitBranch()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}
	command = bash.GitBranchRecurse()
	if err := utils.Exec(command, nil); err != nil {
		log.Println(err.Error())
		return
	}

	// FETCH AND PULL
	command = bash.GitFetchRecurse()
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
	GitCmds.AddCommand(WreckCmd)
}
