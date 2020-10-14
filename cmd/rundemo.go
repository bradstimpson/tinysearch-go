// Package cmd rundemo - this is used to run the demo system mainly
// for testing and evaluation purposes.
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
	"os"
	"strconv"
	"tinysearch/cmd/server"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var p int
var g bool

// rundemoCmd represents the rundemo command
var rundemoCmd = &cobra.Command{
	Use:   "rundemo",
	Short: "Run the demo server to test the generated wasm files",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			log.Error(err)
		}
		log.Infof("Running the dev server port: %v for the demo with directory %v", p, path+"/assets")
		err = server.Run(":"+strconv.Itoa(p), g, path+"/assets", true)
		return err
	},
}

func init() {
	rootCmd.AddCommand(rundemoCmd)
	rundemoCmd.PersistentFlags().IntVarP(&p, "port", "p", 9090, "dev server port")
	rundemoCmd.PersistentFlags().BoolVarP(&g, "gzip", "g", false, "enable gzip compression")
}
