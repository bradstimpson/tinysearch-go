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
	"bytes"
	"compress/gzip"
	"tinysearch/cmd/persister"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	gz       bool   = false
	finalin  string = "build/index.bin"
	finalout string = "build/index.go"
)

func NewSaveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "save",
		Short: "Save the encoded dictionary to a byteslice file",
		Long: `This uses the cuckoo filter encode capability along
with the binary representations of the urls and names to create a bin
file and/or go file to embded.`,
		Run: Save,
	}
}

func init() {
	saveCmd := NewSaveCmd()
	saveCmd.Flags().BoolVarP(&gz, "gzip", "z", false, "Enable gzip on the dictionary binary")
	rootCmd.AddCommand(saveCmd)
}

func SetGzipFlag() {
	gz = true
}

func Save(cmd *cobra.Command, args []string) {
	log.Debugf("Saving the dictionary to an index file for embedding with log level: %v", Verbose)

	if len(args) > 0 {
		if args[0] != "" {
			finalin = args[0]
		}
	}
	if len(args) > 1 {
		if args[1] != "" {
			finalout = args[1]
		}
	}
	var data []byte
	per := persister.NewPersistor(persister.BIN)
	err := per.Load(finalin, &data)
	if err != nil {
		log.Errorf("error loading binary gob file: %v", err)
	}

	var buf bytes.Buffer
	if gz {
		z := gzip.NewWriter(&buf)
		bw, err := z.Write(data)
		if err != nil {
			log.Errorf("error compressing data: %v", err)
		}
		log.Debugf("compressed %d bytes to data variable", bw)
		z.Flush()
		z.Close()
		data = buf.Bytes()
	}

	per = persister.NewPersistor(persister.GO)
	err = per.Save(finalout, &data)
	if err != nil {
		log.Errorf("error saving final output go file: %v", err)
	}

}
