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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setsrcCmd represents the setsrc command
var setsrcCmd = &cobra.Command{
	Use:   "setsrc",
	Short: "Sets the input src domain to get the corpus.json",
	Long:  `This will set the src domain and use whatever variable is set in env or config file for root.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Calling the setsrc command with log level: %v", Verbose)
		setSrc(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setsrcCmd)
}

func setSrc(domain string) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	viper.Set(DefaultRootName, RootDomain)
	viper.Set(DefaultSrcName, domain)
	if err := viper.WriteConfig(); err != nil {
		log.Error(err)
	}
	log.Infof("Set the source domain for corpus.json to: %v", domain)
}
