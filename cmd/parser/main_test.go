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
	Url  string `json:"url"`
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

	err = testParser.Parse("http://example.test")
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
				Url:  urls[i],
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

	assert.Equal(t, "http://example.test/test1/", r0.Url)
	assert.Equal(t, "http://example.test/test2/", r1.Url)

	// f, err := os.OpenFile("../../build/index.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal("Couldn't open file")
	// }
	// defer f.Close()
	// err = binary.Write(f, binary.BigEndian, uint64(len(filters)))
	// for i, _ := range filters {
	// 	var data = struct {
	// 		f uint64
	// 		u uint64
	// 		n uint64
	// 		i uint16
	// 	}{binary.BigEndian.Uint64(filters[i]), binary.BigEndian.Uint64([]byte(urls[i])), binary.BigEndian.Uint64([]byte(names[i])), uint16(i)}

	// 	err = binary.Write(f, binary.BigEndian, data)
	// }
	// if err != nil {
	// 	log.Fatal("Write failed")
	// }

	// f2, err := os.OpenFile("../../build/index.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal("Couldn't open file")
	// }
	// defer f2.Close()

	// file, err := os.Open("../../build/index.bin")
	// if err != nil {
	// 	log.Fatal("Couldn't open file")
	// }
	// defer file.Close()

	// stats, statsErr := file.Stat()
	// if statsErr != nil {
	// 	log.Fatal("stats error")
	// }

	// var size int64 = stats.Size()
	// bytes := make([]byte, size)

	// bufr := bufio.NewReader(file)
	// _, err = bufr.Read(bytes)

	// _, err = f2.WriteString("package main\n")
	// _, err = f2.WriteString("var Index []byte = []byte{")
	// for i, v := range bytes {
	// 	_, err = f2.Write([]byte(fmt.Sprintf("%d", v)))
	// 	if i != len(bytes) {
	// 		_, err = f2.WriteString(",")
	// 	}
	// }

	// _, err = f2.WriteString("}\n")

	// if err != nil {
	// 	log.Fatal("Couldn't open file")
	// }
}
