package parser

import (
	"bufio"
	"html"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/kljensen/snowball"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/thoas/go-funk"
)

var startWords = [...]string{"http", "src", "srcset", "style", "class", "href", "id", "sizes", "height", "width", "loading", "aria"}

type Parser interface {
	Parse(domain string, stem bool, rman bool, rmsw bool) error
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

func (p *parser) Parse(domain string, stem bool, rman bool, rmsw bool) error {

	titlePath, err1 := jp.ParseString("[*].title")
	slugPath, err2 := jp.ParseString("[*].slug")
	contentPath, err3 := jp.ParseString("[*].content")
	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatalf("error parsing json data: %v %v %v", err1, err2, err3)
	}
	titles := titlePath.Get(p.source)
	slugs := slugPath.Get(p.source)
	contents := contentPath.Get(p.source)

	for it, vt := range titles {

		// string into an array of words
		ft := strings.Fields(vt.(string))
		fc := strings.Fields(contents[it].(string))

		var words []string

		// iterate over all the title content
		for _, vs := range ft {
			if !hasStartWord(vs) {
				word := stripAlphaAndUnescape(stripHTMLTags(vs))

				// fmt.Printf("the word being analyzed: %v\n", word)

				if yes, found := hasMultipleWords(word); yes {
					for _, v := range found {
						// fmt.Printf("analyze word: %v\n", v)
						checkStemAndAdd(v, &words)
					}
				} else {
					checkStemAndAdd(word, &words)
				}
			}
		}

		// iterate over all the content
		for _, vs := range fc {
			// fmt.Printf("input word being analyzed: %v\n", vs)

			// 1. check any of the words starting with target start words
			// 2. strip any html tags from the word
			// 3. strip alpha numerics and unescape (if on)
			// 4. check if the word is multiple words (if multiple repeat below for each)
			// 5. check if the word is a stopword (if on)
			// 6. stem the word (if on)
			// 7. finally, add the word

			if !hasStartWord(vs) {
				word := stripAlphaAndUnescape(stripHTMLTags(vs))

				// fmt.Printf("the word being analyzed: %v\n", word)

				if yes, found := hasMultipleWords(word); yes {
					for _, v := range found {
						// fmt.Printf("analyze word: %v\n", v)
						checkStemAndAdd(v, &words)
					}
				} else {
					checkStemAndAdd(word, &words)
				}
			}
		}
		// fmt.Printf("before funk: %v\n", len(words))
		words = funk.UniqString(words) //use funk magic to remove duplicates
		// fmt.Printf("after funk: %v\n", len(words))
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

func checkStemAndAdd(word string, words *[]string) {
	if ok, _ := isStopword(strings.TrimSpace(word)); !ok {
		// fmt.Printf("not a stopword, adding: %v\n", word)
		*words = append(*words, stemWord(word))
	}
}

func stripHTMLTags(word string) string {
	return strip.StripTags(word)
}

func stripAlphaAndUnescape(word string) string {
	reg, _ := regexp.Compile("[^'a-zA-Z0-9]+")
	return reg.ReplaceAllString(html.UnescapeString(word), " ")
}

func stemWord(word string) string {
	out, err := snowball.Stem(word, "english", true)
	if err != nil {
		log.Fatalf("error unescaping and stemming titles: %v", err)
	}
	return out
}

func hasStartWord(word string) bool {
	for _, v := range startWords {
		if strings.HasPrefix(word, v) {
			return true
		}
	}
	return false
}

func hasMultipleWords(word string) (bool, []string) {
	s := strings.Split(strings.TrimSpace(word), " ")
	if len(s) > 1 {
		return true, s
	}
	return false, nil
}

func isStopword(word string) (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	// small fix for running tests and normal mode
	if strings.HasSuffix(dir, "cmd") {
		dir = dir + "/parser"
	}

	f, err := os.Open(dir + "/stopwords")
	if err != nil {
		return false, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == strings.ToLower(word) {
			return true, nil
		}
		if err := scanner.Err(); err != nil {
			return false, err
		}
	}
	return false, nil
}
