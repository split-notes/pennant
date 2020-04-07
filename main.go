/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/split-notes/pennant/cli/cmd"
	"github.com/split-notes/pennant/cli/config"
	"log"
	"os"
)

func main() {
	cobra.OnInitialize(initConfig)

	cmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	cfgFile := home + "/.pennant.yaml"

	// If the config file doesn't exist, create it
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Println("Required config not found at:", cfgFile)
		fmt.Println("Creating initial file:", cfgFile)
		if err := config.GenerateSampleConfigYaml(cfgFile); err != nil {
			log.Println(err.Error())
			panic("error creating the missing config file")
		}
	}

	// Once we know the config file exists, load it in
	var conf *config.Configurations
	conf, err = config.LoadConfigYaml(cfgFile)
	if err != nil {
		log.Println(err.Error())
		panic("error opening file")
	}

	// Store the config in viper
	config.StoreConfigInViper(conf)
}
