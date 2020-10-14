// +build integration

package tests

import (
	"testing"
	"tinysearch/cmd"

	"github.com/spf13/cobra"
)

func TestDownloadCmd(t *testing.T) {
	cmd.UpdateSrcDomain("http://192.168.0.28:8009?rest_route=/wp/v2/posts&page=1&per_page=100")
	cmd.Download(&cobra.Command{}, []string{"../build/corpus.json"})
}

func TestParseCmd(t *testing.T) {
	cmd.DisableStemming()
	cmd.UpdateRootDomain("http://example.test")
	cmd.Parse(&cobra.Command{}, []string{"../build/corpus.json", "../build/index.bin"})
}

func TestSaveCmd(t *testing.T) {
	cmd.SetGzipFlag()
	cmd.Save(&cobra.Command{}, []string{"../build/index.bin", "../build/index.go"})
}

func TestBuildCmd(t *testing.T) {
	cmd.Build(&cobra.Command{}, []string{"NOP"})
}

func TestSearchCmd(t *testing.T) {
	cmd.DisableOutput()
	cmd.Search(&cobra.Command{}, []string{"test", "../build/index.bin"})
}
