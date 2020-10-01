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
	"tinysearch/cmd/parser"
	"tinysearch/cmd/persister"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	sw  bool   = false
	an  bool   = false
	in  string = "build/corpus.json"
	out string = "build/index.bin"
)

// parseCmd represents the parse command
func NewParseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "parse",
		Short: "Parse the index.json file into the search dictionary",
		Long: `The download subcommand takes two arguments, the path and filename to 
corpus.json and where to store the index.bin.  If not provided it defaults to build 
directory with filenames corpus.json and index.bin.`,
		Run: Parse,
	}
}

func UpdateRootDomain(domain string) {
	RootDomain = domain
}

func init() {
	parseCmd := NewParseCmd()
	parseCmd.Flags().BoolVarP(&sw, "stopwords", "w", false, "Remove stopwords from dictionary")
	parseCmd.Flags().BoolVarP(&an, "alphanumerics", "j", false, "Remove alphanumerics from dictionary")
	rootCmd.AddCommand(parseCmd)
}

func Parse(cmd *cobra.Command, args []string) {
	log.Debugf("Parsing with root domain %s with log level: %v", RootDomain, Verbose)

	if len(args) > 0 {
		if args[0] != "" {
			in = args[0]
		}
	}
	if len(args) > 1 {
		if args[1] != "" {
			out = args[1]
		}
	}

	par := parser.NewParser()
	err := par.Source(in)
	if err != nil {
		log.Errorf("error parsing corpus.json: %v", err)
	}

	err = par.Parse(RootDomain)
	if err != nil {
		log.Errorf("error parsing to final dictionary: %v", err)
	}

	filters, urls, names := par.Encode()
	per := persister.NewPersistor()
	err = per.SaveGOB(out, struct {
		F [][]byte
		U []string
		N []string
	}{
		F: filters,
		U: urls,
		N: names,
	})
	if err != nil {
		log.Errorf("error saving index.bin: %v", err)
	}
}
