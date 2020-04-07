package golang

import (
	"github.com/spf13/cobra"
)

var GolangCmds = &cobra.Command{
	Use:     "golang",
	Aliases: []string{"go"},
	Short:   "golang command start",
	Run:     Golang,
}

func Golang(_ *cobra.Command, _ []string) {}
