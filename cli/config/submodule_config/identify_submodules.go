package submodule_config

import (
	"encoding/json"
	"fmt"
	"github.com/split-notes/pennant/cli/config"
	"io/ioutil"
)


func IdentifySubmodules() ([]Submodule, error) {
	var configs = config.GetConfigFromViper()

	submodulesFilePath := fmt.Sprintf("%s/deployments/submodules", configs.PennantConfig.ProjectFilePath)
	files, err := ioutil.ReadDir(submodulesFilePath)
	if err != nil { return nil, err}

	var submodules []Submodule
	// Array of absolute paths to each submodule
	for _, f := range files {
		projectPath := fmt.Sprintf("%s/%s", submodulesFilePath, f.Name())

		configFileLocation := fmt.Sprintf("%s/.config.json", projectPath)

		configFile, err := ioutil.ReadFile(configFileLocation)
		if err != nil {
			return nil, err
		}

		configData := Submodule{
			ProjectPath: projectPath,
			ProjectName: f.Name(),
		}
		_ = json.Unmarshal(configFile, &configData)

		submodules = append(submodules, configData)
	}

	return submodules, nil
}
