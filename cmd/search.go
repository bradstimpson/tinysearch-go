// Package cmd search - this is a convenience function to search the
// generated index files to evaluate the search quality.
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
	"encoding/json"
	"fmt"
	"tinysearch/cmd/persister"

	cuckoo "github.com/seiflotfy/cuckoofilter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	searchTerm string = ""
	gobFile    string = "build/index.bin"
	rslts      int    = 5
	nostd      bool   = false
)

// Data is the expected struct format after by gobbered
type Data struct {
	F [][]byte
	U []string
	N []string
}

// Result is the expect format of the output json for search
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var data Data

// NewSearchCmd is the search cobra command
func NewSearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Searches the dictionary for the specified term",
		Run:   Search,
	}
}

func init() {
	searchCmd := NewSearchCmd()
	searchCmd.PersistentFlags().IntVarP(&rslts, "results", "l", 5, "number of results to return (default=5)")
	searchCmd.PersistentFlags().BoolVar(&nostd, "no-output", false, "flag to disable output")
	rootCmd.AddCommand(searchCmd)
}

// DisableOutput is used primarily in testing to programatically change whether
// to output the results of the search to the screen
func DisableOutput() {
	nostd = true
}

// Search is called by NewSearchCmd and is exported for testing
// purposes.  The args are passed via the command line.
func Search(cmd *cobra.Command, args []string) {
	//*TODO* Update with the new serializer option
	log.Debug("Search subcommand run with log level: ", Verbose, "\n")

	if len(args) > 0 {
		if args[0] != "" {
			searchTerm = args[0]
		}
	}
	if len(args) > 1 {
		if args[1] != "" {
			gobFile = args[1]
		}
	}

	np := persister.NewPersistor(persister.GOB)
	err := np.Load(gobFile, &data)
	if err != nil {
		log.Errorf("error loading gob file: %v", err)
	}

	filters, urls, names := data.F, data.U, data.N

	var found []interface{}
	// iterate through the filters and return indices of matches
	for i, v := range filters {
		filter, _ := cuckoo.Decode(v)
		if filter.Lookup([]byte(searchTerm)) {
			if len(found) >= rslts {
				break
			}
			found = append(found, Result{
				Name: names[i],
				URL:  urls[i],
			})
		}
	}

	if !nostd {
		// ** PRETTY OUTPUT FOR USE AT COMMAND LINE **
		foundJSON, err := json.MarshalIndent(found, "", "  ")
		if err != nil {
			log.Errorf("error in marshalling the found results: %v", err)
		}
		fmt.Printf("Search Results:\n %s\n", string(foundJSON))
	}

}
