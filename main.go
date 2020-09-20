package main

import (
	"html"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kljensen/snowball"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	cuckoo "github.com/seiflotfy/cuckoofilter"
)

const (
	ROOT_DOMAIN = "https://www.andthenhome.com/"
)

var filters []*cuckoo.Filter
var urls []interface{}

func main() {
	//load json file
	result, err := parseFile("output.json")
	if err != nil {
		log.Println("an error occurred parsing the file", err)
	}
	titlePath, err1 := jp.ParseString("[*].title")
	slugPath, err2 := jp.ParseString("[*].slug")
	contentPath, err3 := jp.ParseString("[*].content")
	if err1 != nil || err2 != nil || err3 != nil {
		log.Println("error parsing json data")
	}
	titles := titlePath.Get(result)
	slugs := slugPath.Get(result)
	contents := contentPath.Get(result)

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for it, vt := range titles {

		// string into an array of words
		ft := strings.Fields(vt.(string))
		fc := strings.Fields(contents[it].(string))

		var words []string
		// iterate over all the titles
		for _, vs := range ft {
			ps := reg.ReplaceAllString(html.UnescapeString(vs), "")
			stemmed, err := snowball.Stem(ps, "english", true)
			if err == nil {
				words = append(words, stemmed)
			}
		}
		// iterate over all the content
		for _, vs := range fc {
			ps := reg.ReplaceAllString(html.UnescapeString(vs), "")
			stemmed, err := snowball.Stem(ps, "english", true)
			if err == nil {
				words = append(words, stemmed)
			}
		}
		// fmt.Println(sTitles)
		cf := cuckoo.NewFilter(uint(len(words)))
		for _, value := range words {
			cf.InsertUnique([]byte(value))
		}

		filters = append(filters, cf)

		urls = append(urls, ROOT_DOMAIN+slugs[it].(string)+"/")

	}
}

func search(term string) ([]interface{}, error) {
	var found []interface{}
	// iterate through the filters and return indices of matches
	for i, v := range filters {
		if v.Lookup([]byte(term)) {
			found = append(found, urls[i])
		}
	}

	return found, nil
}

func parseFile(fileName string) (interface{}, error) {
	jsonfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the json file", err)
		return nil, err
	}
	defer jsonfile.Close()

	var buf strings.Builder
	_, err = io.Copy(&buf, jsonfile)
	if err != nil {
		// handle error
		log.Fatalln("Could not convert file to string", err)
		return nil, err
	}
	s := buf.String()

	//parse the file
	result, err := oj.ParseString(s)
	if err != nil {
		log.Fatalln("Could not parse json file", err)
		return nil, err
	}
	return result, nil
}
