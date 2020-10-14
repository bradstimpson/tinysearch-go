// Package cmd setroot - this is a convenience function to set the
// root domain used in the output URLs
/*
Copyright © 2020 Brad Stimpson <brad.stimpson@gmail.com>

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
package cmd

import (
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setrootCmd represents the setroot command
var setrootCmd = &cobra.Command{
	Use:   "setroot",
	Short: "Sets the output root domain in the local config file",
	Long:  `This will set the root domain and use whatever variable is set in env or config file for src.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Calling the setroot command with log level: %v", Verbose)
		setRoot(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setrootCmd)
}

func setRoot(domain string) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	viper.Set(DefaultRootName, domain)
	viper.Set(DefaultSrcName, SrcDomain)
	if err := viper.WriteConfig(); err != nil {
		log.Error(err)
	}
	log.Infof("Set the root domain for output url to: %v", domain)
}
