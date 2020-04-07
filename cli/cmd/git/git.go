package git

import (
	"github.com/spf13/cobra"
)

var GitCmds = &cobra.Command{
	Use:     "git",
	Aliases: []string{"gi"},
	Short:   "git command start",
	Run:     Git,
}

func Git(_ *cobra.Command, _ []string) {}
