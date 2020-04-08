package git_svc

import (
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/split-notes/pennant/cli/config/submodule_config"
)

func SelectSubmodules(languageFilter *string, transportFilter *string) ([]submodule_config.Submodule, error) {
	submodules, err := submodule_config.IdentifySubmodules(languageFilter, transportFilter)
	if err != nil {
		return nil, err
	}
	results, err := fuzzyfinder.FindMulti(submodules,
		func(i int) string {
			return submodules[i].ProjectName
		})
	if err != nil {
		// nothing selected, abort.
		return nil, err
	}
	var selected []submodule_config.Submodule
	for _, i := range results {
		selected = append(selected, submodules[i])
	}
	return selected, nil
}
