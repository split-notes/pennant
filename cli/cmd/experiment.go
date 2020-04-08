package cmd

import (
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/config/submodule_config"
	"log"
)

var ExperimentCmd = &cobra.Command{
	Use:     "experiment",
	Aliases: []string{"exp"},
	Short:   "a command endpoint to experiment with stuff",
	Long:    ``,
	Run:     Experiment,
}

func Experiment(_ *cobra.Command, _ []string) {
	submodules, err := submodule_config.IdentifySubmodules(nil, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	results, err := fuzzyfinder.FindMulti(submodules,
		func(i int) string {
			return submodules[i].ProjectName
		})
	if err != nil {
		log.Println("no submodules selected")
		return
	}
	var selected []submodule_config.Submodule
	for _, i := range results {
		selected = append(selected, submodules[i])
	}
	log.Println(selected)

	//return &splitOutput[selected], nil
}

func init() {
	rootCmd.AddCommand(ExperimentCmd)
}
