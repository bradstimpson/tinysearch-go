// Package downloader - this is used to download the corpus.json from
// the target source domain.
package downloader

import (
	"bytes"
	"encoding/json"
	"errors"

	log "github.com/sirupsen/logrus"
)

// Downloader interface to specify the main functions that
// a downloader must have (namely Get)
type Downloader interface {
	Get(domain string) (Posts, error)
}

var (
	dlService DLService
	jl        bool = false
	hu        bool = false
)

type downloader struct{}

// NewDownloader is the implementation that expects to have
// a DLService (service to actually get the content) and one of the
// supported service types (w=wordpress, j=jekyll, h=hugo)
func NewDownloader(d DLService, w bool, j bool, h bool) Downloader {
	jl, hu = j, h
	dlService = d
	return &downloader{}
}

// Posts is a slice of all the extracted posts from the target system
// in the desired format.
type Posts []Post

// Post is the desired format for the output
type Post struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
}

func (*downloader) Get(domain string) (Posts, error) {

	r1, dlErr := dlService.Fetch(domain)

	var err error = <-dlErr
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	result := <-r1
	if result == nil {
		return nil, errors.New("an error occured fetching from " + domain)
	}

	_, err1 := buf.ReadFrom(result.Body)
	if err1 != nil {
		return nil, err1
	}

	var p Posts

	if jl {
		return nil, errors.New("Jekyll marshalling method not currently supported")
	} else if hu {
		return nil, errors.New("Hugo marshalling method not currently supported")
	} else {
		err2 := p.unmarshalWP(buf.Bytes())
		if err2 != nil {
			log.Errorf("an error occured: %v", err2)
		}
	}

	return p, nil
}

func (p *Posts) unmarshalWP(b []byte) error {
	internal := []struct {
		Slug    string `json:"slug"`
		Content struct {
			Rendered string `json:"rendered"`
		} `json:"content"`
		Title struct {
			Rendered string `json:"rendered"`
		} `json:"title"`
	}{}

	if err := json.Unmarshal(b, &internal); err != nil {
		return err
	}

	for _, value := range internal {
		*p = append(*p, Post{
			Content: value.Content.Rendered,
			Title:   value.Title.Rendered,
			Slug:    value.Slug,
		})

	}
	return nil
}

// func isJSONString(s string) bool {
// 	var js string
// 	return json.Unmarshal([]byte(s), &js) == nil

// }

// func isJSON(s string) bool {
// 	var js map[string]interface{}
// 	return json.Unmarshal([]byte(s), &js) == nil

// }
