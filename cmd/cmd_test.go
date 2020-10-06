package cmd_test

import (
	"os"
	"strings"
	"testing"
	"tinysearch/cmd"

	"github.com/spf13/cobra"
)

func TestDownloadCmd(t *testing.T) {
	if strings.Compare(os.Getenv("TEST_INT"), "true") == 1 {
		t.Skip("Skipping download integration test")
	}
	cmd.UpdateSrcDomain("http://192.168.0.28:8009?rest_route=/wp/v2/posts&page=1&per_page=100")
	cmd.Download(&cobra.Command{}, []string{"../build/corpus.json"})
}

func TestParseCmd(t *testing.T) {
	if strings.Compare(os.Getenv("TEST_INT"), "true") == 1 {
		t.Skip("Skipping parse integration test")
	}
	cmd.DisableStemming()
	cmd.UpdateRootDomain("http://example.test")
	cmd.Parse(&cobra.Command{}, []string{"../build/corpus.json", "../build/index.bin"})
}

func TestSaveCmd(t *testing.T) {
	if strings.Compare(os.Getenv("TEST_INT"), "true") == 1 {
		t.Skip("Skipping save integration test")
	}
	cmd.SetGzipFlag()
	cmd.Save(&cobra.Command{}, []string{"../build/index.bin", "../build/index.go"})
}

func TestBuildCmd(t *testing.T) {
	if strings.Compare(os.Getenv("TEST_INT"), "true") == 1 {
		t.Skip("Skipping build integration test")
	}
	cmd.Build(&cobra.Command{}, []string{"NOP"})
}

func TestSearchCmd(t *testing.T) {
	if strings.Compare(os.Getenv("TEST_INT"), "true") == 1 {
		t.Skip("Skipping search integration test")
	}
	cmd.DisableOutput()
	cmd.Search(&cobra.Command{}, []string{"test", "../build/index.bin"})
}
