package parser

import (
	"encoding/json"
	"testing"

	"github.com/ohler55/ojg/jp"
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/stretchr/testify/assert"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func TestRunParserSource(t *testing.T) {
	testParser := NewParser()
	err := testParser.Source("../fixtures/test-wp-corpus.json")
	result := testParser.GetSource()
	assert.Nil(t, err)
	assert.NotNil(t, result)

	titlePath, _ := jp.ParseString("[*].title")
	titles := titlePath.Get(result)
	assert.Equal(t, "Test Title 1", titles[0])
}

func TestRunParser(t *testing.T) {
	testParser := NewParser()
	err := testParser.Source("../fixtures/test-wp-corpus.json")
	assert.Nil(t, err)

	err = testParser.Parse("http://example.test", true, true, true)
	assert.Nil(t, err)
	filters, urls, names := testParser.Encode()

	var found []interface{}
	// iterate through the filters and return indices of matches
	for i, v := range filters {
		filter, _ := cuckoo.Decode(v)
		if filter.Lookup([]byte("test")) {
			var jsonData []byte
			jsonData, err := json.Marshal(Result{
				Name: names[i],
				URL:  urls[i],
			})
			assert.Nil(t, err)
			found = append(found, string(jsonData))
		}
	}
	r0, r1 := Result{}, Result{}
	err = json.Unmarshal([]byte(found[0].(string)), &r0)
	assert.Nil(t, err)
	err = json.Unmarshal([]byte(found[1].(string)), &r1)
	assert.Nil(t, err)

	assert.Equal(t, "http://example.test/test1/", r0.URL)
	assert.Equal(t, "http://example.test/test2/", r1.URL)
}
