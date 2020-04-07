package git

import (
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/bash"
	"github.com/split-notes/pennant/cli/utils"
	"log"
	"os/exec"
)

var OnboardGitCmd = &cobra.Command{
	Use:     "onboard",
	Aliases: []string{"on"},
	Short:   "Onboard git stuff",
	Long:    `Sets up your git repos`,
	Run:     OnboardGit,
}

func OnboardGit(_ *cobra.Command, _ []string) {
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
}

func init() {
	GitCmds.AddCommand(OnboardGitCmd)
}
