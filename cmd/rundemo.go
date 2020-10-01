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

var Port int

// rundemoCmd represents the rundemo command
var rundemoCmd = &cobra.Command{
	Use:   "rundemo",
	Short: "Run the demo server to test the generated wasm files",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			log.Error(err)
		}
		log.Infof("Running the dev server port: %v for the demo with directory %v", Port, path+"/assets")
		server.Run(":"+strconv.Itoa(Port), path+"/assets")
	},
}

func init() {
	rootCmd.AddCommand(rundemoCmd)
	rundemoCmd.PersistentFlags().IntVarP(&Port, "port", "p", 9090, "dev server port")
}
