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
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	tiny    bool
	sd      bool
	sl      bool
	outfile string = "tiny.wasm"
)

func NewBuildCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Builds the wasm file from the saved encoded dictionary file",
		Run:   Build,
	}
}
func init() {
	buildCmd := NewBuildCmd()
	buildCmd.PersistentFlags().BoolVarP(&tiny, "tinygo", "y", false, "enable the tinygo optimization")
	buildCmd.PersistentFlags().BoolVarP(&sd, "sdebug", "p", false, "enable the strip debug optimization")
	buildCmd.PersistentFlags().BoolVarP(&sl, "sline", "l", false, "enable the strip line number optimization")
	rootCmd.AddCommand(buildCmd)
}

func Build(cmd *cobra.Command, args []string) {
	log.Debugf("Building the final wasm file with log level: %v", Verbose)

	if len(args) > 0 {
		if args[0] != "" {
			in = args[0]
		}
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting working directory: %v", err)
	}

	c := exec.Command("bash", "-c", `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../assets`)
	c.Dir = dir
	if err := c.Run(); err != nil {
		log.Errorf("error copying the wasm_exec.js file: %v", err.Error())
	}

	c = exec.Command("bash", "-c", "mv ../build/index.go wasmgo/")
	c.Dir = dir
	if err := c.Run(); err != nil {
		log.Errorf("error moving compiled file: %v", err.Error())
	}

	c = exec.Command("bash", "-c", `GOOS=js GOARCH=wasm go build  -ldflags="-s -w" -o `+outfile)
	c.Dir = dir + "/wasmgo"
	if err := c.Run(); err != nil {
		log.Errorf("error running go compiler: %v", err.Error())
	}

	// c = exec.Command("bash", "-c", `tinygo build -o `+outfile+` -target wasm main.go`)
	// c.Dir = dir + "/wasmgo"
	// if err := c.Run(); err != nil {
	// 	log.Errorf("error running go compiler: %v", err.Error())
	// }

	c = exec.Command("bash", "-c", "mv "+outfile+" ../../assets/")
	c.Dir = dir + "/wasmgo"
	if err := c.Run(); err != nil {
		log.Errorf("error moving compiled file: %v", err.Error())
	}

	c = exec.Command("bash", "-c", "rm -rf index.go")
	c.Dir = dir + "/wasmgo"
	if err := c.Run(); err != nil {
		log.Errorf("error moving compiled file: %v", err.Error())
	}

}
