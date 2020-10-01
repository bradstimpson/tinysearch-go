package parser

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

type Parser interface {
	Parse(domain string) error
	Source(fileName string) error
	Encode() ([][]byte, []string, []string)
	GetSource() interface{}
}

type parser struct {
	source  interface{}
	filters []*cuckoo.Filter
	urls    []string
	names   []string
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(domain string) error {

	titlePath, err1 := jp.ParseString("[*].title")
	slugPath, err2 := jp.ParseString("[*].slug")
	contentPath, err3 := jp.ParseString("[*].content")
	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatalf("error parsing json data: %v %v %v", err1, err2, err3)
	}
	titles := titlePath.Get(p.source)
	slugs := slugPath.Get(p.source)
	contents := contentPath.Get(p.source)

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatalf("error with regular expression step removing unneeded characters: %v", err)
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
			} else {
				log.Fatalf("error unescaping and stemming titles: %v", err)
				return err
			}
		}
		// iterate over all the content
		for _, vs := range fc {
			ps := reg.ReplaceAllString(html.UnescapeString(vs), "")
			stemmed, err := snowball.Stem(ps, "english", true)
			if err == nil {
				words = append(words, stemmed)
			} else {
				log.Fatalf("error unescaping and stemming titles: %v", err)
				return err
			}
		}
		// fmt.Println(sTitles)
		cf := cuckoo.NewFilter(uint(len(words)))
		for _, value := range words {
			cf.InsertUnique([]byte(value))
		}

		p.filters = append(p.filters, cf)
		p.names = append(p.names, vt.(string))
		p.urls = append(p.urls, domain+"/"+slugs[it].(string)+"/")
	}
	return nil
}

func (p *parser) Source(fileName string) error {
	jsonfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the json file", err)
		return err
	}
	defer jsonfile.Close()

	var buf strings.Builder
	_, err = io.Copy(&buf, jsonfile)
	if err != nil {
		// handle error
		log.Fatalln("Could not convert file to string", err)
		return err
	}
	s := buf.String()

	//parse the file
	source, err := oj.ParseString(s)
	if err != nil {
		log.Fatalln("Could not parse json file", err)
		return err
	}
	p.source = source
	return nil
}

func (p *parser) GetSource() interface{} {
	return p.source
}

//returns the encoded cuckoo filters in a byteslice array along with an array of urls
func (p *parser) Encode() ([][]byte, []string, []string) {
	var ef [][]byte
	var eu []string
	var en []string
	for i, v := range p.filters {
		ef = append(ef, v.Encode())
		eu = append(eu, p.urls[i])
		en = append(en, p.names[i])
	}
	return ef, eu, en
}
