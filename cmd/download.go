// Package cmd download - this will download the corpus.json
// from the specified system.
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
	"tinysearch/cmd/persister"

	"tinysearch/cmd/downloader"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	wp       bool   = false
	jl       bool   = false
	hu       bool   = false
	filename string = "build/corpus.json"
)

// NewDownloadCmd is the cobra command for downloading the corpus
func NewDownloadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "Downloads and prepares the corpus.json file from target cms/ssg",
		Long: `The download subcommand takes one argument, the path and filename to 
store the corpus.json.  If not provided it defaults to build/corpus.json.`,
		Run: Download,
	}
}

func init() {
	downloadCmd := NewDownloadCmd()
	downloadCmd.Flags().BoolVarP(&wp, "wordpress", "w", false, "Download against WP")
	downloadCmd.Flags().BoolVarP(&jl, "jekyll", "j", false, "Download against Jekyll")
	downloadCmd.Flags().BoolVarP(&hu, "hugo", "u", false, "Download against Hugo")
	rootCmd.AddCommand(downloadCmd)
}

// UpdateSrcDomain is used primarily in testing to programatically change the source
// domain to download from.
func UpdateSrcDomain(domain string) {
	SrcDomain = domain
}

// Download is called by NewDownloadCmd and is exported for testing
// purposes.  The args are passed via the command line.
func Download(cmd *cobra.Command, args []string) {
	log.Debugf("Downloading from %s with log level: %v", SrcDomain, Verbose)

	var svc downloader.DLService = downloader.NewDLService()
	var dw downloader.Downloader = downloader.NewDownloader(svc, wp, jl, hu)

	if len(args) > 0 {
		if args[0] != "" {
			filename = args[0]
		}
	}

	result, err := dw.Get(SrcDomain)
	if err != nil {
		log.Errorf("error getting source domain %s: %v", SrcDomain, err)
	}

	p := persister.NewPersistor(persister.JSON)
	err = p.Save(filename, result)
	if err != nil {
		log.Errorf("error saving corpus.json: %v", err)
	}
}
