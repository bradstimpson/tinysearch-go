// +build js,wasm

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"

	cuckoo "github.com/seiflotfy/cuckoofilter"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var filters []*cuckoo.Filter
var urls []string
var names []string

type Dictionary struct {
	F [][]byte
	U []string
	N []string
}

func search(term string, results int) (interface{}, error) {

	var found []interface{}
	var total int
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
			if total < results {
				found = append(found, string(jsonData))
			}
			total++
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
		urls, err := search(searchTerm, 5) //returns
		if err != nil {
			fmt.Printf("unable to find term", err)
			return err.Error()
		}
		fmt.Print(urls)
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

func buildIndex() error {

	var ix Dictionary
	r := bytes.NewReader(Index)
	gr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	err = gob.NewDecoder(gr).Decode(&ix)
	if err != nil {
		return err
	}

	for i, v := range ix.F {
		filter, _ := cuckoo.Decode(v)
		filters = append(filters, filter)
		urls = append(urls, ix.U[i])
		names = append(names, ix.N[i])
	}
	return nil
}
