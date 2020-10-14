// Package cmd root - this is the base command used mainly to setup
// the configuration and kick-off the tool.
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
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// SrcDomain is the source domain to retrieve the corpus
var SrcDomain string

// RootDomain is the domain to append to all outgoing URLs
var RootDomain string

// Verbose is the log level
var Verbose string = "warning"

// DefaultConfigName is the default configuration file name (appended with yml)
var DefaultConfigName string = ".tinysearch"

// DefaultConfigPath is the default path to find the config
var DefaultConfigPath string = "."

// internal variables
var cfgFile string

// DefaultSrcName is the name to look for in the config files for the source domain
var DefaultSrcName string = "SrcDomain"

// DefaultRootName is the name to look for in the config files for the root domain
var DefaultRootName string = "RootDomain"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tinysearch",
	Short: "Tinysearch is small utility to build a wasm search index",
	Long: `++TINYSEARCH-GO++
=================
Inspired by endler.dev this is the Golang equivalent of his Rust application.  

This tool can pull the post index from cms/site generator, parse it into a cuckoo 
dictionary filter, save/load it, and build it into a wasm by calling the go or 
tinygo compiler.

Calling tinysearch directly without a subcommand will search the parsed index.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("Running prerun with log level: ", Verbose, "\n")
		if err := setUpLogs(os.Stdout, Verbose); err != nil {
			return err
		}
		return nil
	},
}

// Execute is used by cobra to kick-off the CLI
func Execute() {
	log.Debug("Running the main execute function with log level: ", Verbose, "\n")
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Debug("Completed the main execute function with log level: ", Verbose, "\n")
}

func init() {
	log.Debug("Running init with log level: ", Verbose, "\n")
	if err := setUpLogs(os.Stdout, Verbose); err != nil {
		log.Error(err)
	}
	initConfig()
	// cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is .tinysearch.yaml in pwd)")
	rootCmd.PersistentFlags().StringVarP(&SrcDomain, "src", "s", "", "the src domain to build the input corpus.json")
	rootCmd.PersistentFlags().StringVarP(&RootDomain, "root", "r", "", "the root domain for all output URLs")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Verbose, "verbosity", "v", log.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
}

//setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level string) error {
	log.Debug("Running setuplogs with log level: ", Verbose, "\n")
	log.SetOutput(out)
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.Infof("Setting the log level to: %v\n", level)
	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Debug("Running initconfig with log level: ", Verbose, "\n")
	// Check if command line vars set, if not check ENV vars, then finally config file
	if SrcDomain == "" || RootDomain == "" {
		viper.SetEnvPrefix("tiny")
		viper.AutomaticEnv()

		//

		if viper.GetString(DefaultSrcName) != "" && viper.GetString(DefaultRootName) != "" {

			//
			SrcDomain = viper.GetString(DefaultSrcName)
			RootDomain = viper.GetString(DefaultRootName)
			log.Info("Using environment variables for necessary URLs: TINY_SRCDOMAIN and TINY_ROOTDOMAIN")
		} else {
			if cfgFile != "" {
				// Use config file from the flag.
				viper.SetConfigFile(cfgFile)
			} else {

				// Search config in home directory with name ".tinysearch" (without extension).
				viper.AddConfigPath(DefaultConfigPath)
				viper.SetConfigName(DefaultConfigName)
			}

			// If a config file is found, read it in.
			if err := viper.ReadInConfig(); err == nil {
				log.Info("Using config file for necessary URLs:", viper.ConfigFileUsed())

				//
				SrcDomain = viper.GetString(DefaultSrcName)
				RootDomain = viper.GetString(DefaultRootName)
			}
		}
	} else {
		log.Info("Using command line variables for necessary URLs: -s and -r")
	}
	log.Debug("The source domain for the input corpus: ", SrcDomain)
	log.Debug("The root domain for the output dictionary: ", RootDomain)

}
