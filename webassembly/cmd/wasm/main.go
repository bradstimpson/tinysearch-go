// +build js,wasm

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"regexp"
	"strings"
	"syscall/js"

	"github.com/gobuffalo/packr"
	"github.com/kljensen/snowball"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	cuckoo "github.com/seiflotfy/cuckoofilter"
)

const (
	ROOT_DOMAIN = "https://www.andthenhome.com/"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var filters []*cuckoo.Filter
var urls []string
var names []string

func buildIndex() {
	//load json file
	result, err := parseIndex()
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
		names = append(names, vt.(string))
		urls = append(urls, ROOT_DOMAIN+slugs[it].(string)+"/")

	}
}

func search(term string) (interface{}, error) {

	var found []interface{}
	// iterate through the filters and return indices of matches
	for i, v := range filters {
		if v.Lookup([]byte(term)) {
			var jsonData []byte
			jsonData, err := json.Marshal(Result{
				Name: names[i],
				Url:  urls[i],
			})
			if err != nil {
				log.Println(err)
			}
			found = append(found, string(jsonData))
		}
	}

	return found, nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		searchTerm := args[0].String()
		fmt.Printf("input %s\n", searchTerm)
		urls, err := search(searchTerm) //returns
		if err != nil {
			fmt.Printf("unable to find term", err)
			return err.Error()
		}
		fmt.Println(urls)
		return urls
	})
	return jsonFunc
}

func main() {
	fmt.Println("Go Tiny Search")
	buildIndex()
	js.Global().Set("search", jsonWrapper())
	<-make(chan bool)
}

func parseIndex() (interface{}, error) {
	box := packr.NewBox("./fixtures")
	jsonfile, err := box.FindString("index.json")
	if err != nil {
		log.Fatal(err)
	}

	// jsonfile, err := t.Execute(os.Open("index.json"))
	// if err != nil {
	// 	log.Fatalln("Couldn't open the json file", err)
	// 	return nil, err
	// }
	// defer jsonfile.Close()

	// var buf strings.Builder
	// _, err = io.Copy(&buf, jsonfile)
	// if err != nil {
	// 	// handle error
	// 	log.Fatalln("Could not convert file to string", err)
	// 	return nil, err
	// }
	// s := buf.String()

	//parse the file
	result, err := oj.ParseString(jsonfile)
	if err != nil {
		log.Fatalln("Could not parse json file", err)
		return nil, err
	}
	return result, nil
}
