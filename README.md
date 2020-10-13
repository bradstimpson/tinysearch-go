# Tinysearch-Go

This is the story behind the creation of Tinysearch-go.  This application was created to build a tinysearch, inspired from endler.dev, from a wordpress system converted into a static site.  Over the course of several weeks, I did many passes over different ideas to build the tinysearch with a goal of getting the binary size down into the same range as endler.dev (150KB).

## The First Pass - Crude Application (check out branch pass1)

The first pass at building a tiny search in golang consisted of:

1. Parsing the fixtures within the binary
2. Building the wasm and using packr to embed the fixtures
3. Manually testing from the command line

We used packr to embed the fixtures into the compile wasm file:

```bash
GOOS=js GOARCH=wasm packr build -o  ../../assets/json.wasm
```

Checking the size of the resulting file: `stat -ln ../../assets/json.wasm`

```bash
-rwxr-xr-x 1 brad staff 11796518 Sep 19 11:52:01 2020 ../../assets/json.wasm
```

An 11.8MB wasm file certainly doesn't constitute as a 'tiny search'.  We can do better.  Also to make matters worse, there is a major delay at first load in the index.html to build the search index.

Let's try separating the dictionary building from the wasm build.

### Extract Data from Wordpress API

The first step is to extract the data we want to use from the Wordpress JSON API:

```bash
curl http://192.168.0.28:8009?rest_route=/wp/v2/posts\&page\=2\&per_page\=90 | jq '.[] |{title:.title.rendered,slug:.slug,content:.content.rendered}'
```

this put the data in a 'near' JSON form - I simply had to go through it and add commas plus array.

The resulting corpus file is 2.1MB.

### Copy over Javascript Glue Code

Run the following from within the project directory.

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets/
```

### Webassembly

Compile the wasm script

```bash
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm
```

this outputs the wasm file that needs to be loaded by javascript

## Second Pass (check out first master push)

In this refactoring we will simplify the directory structure, adopt the use of viper and cobra to build a functional parsing tool and finally output the built dictionary into a separate file that can be packed into the final wasm without bundling the large unparsed file.  We will also update terminology/filenames to make usage easier to understand.

Using Viper allows us manage configuration for the tool.  These are the variables read in from configuration:

* RootDomain: this is the domain name that will be prepended to all URLs in the dictionary (cmd: -r, env: TINY_ROOTDOMAIN)
* SrcDomain: this is where to query for the post index (cmd: -s, env: TINY_SRCDOMAIN)

The priority of reading config vars in are: 1) command line, 2) environment variables, and 3) local config file.

The main application is split into four modules: 1) downloader, 2) parser, 3) persister, and 4) builder.  These modules correspond to the 4 stages we run through to prepare the tinysearch module.

### Run Demo

There is a convience method call run demo which will start a simple http server sharing the assets folder.  To enable gzip compression use the -g flag (no gzip is the default) and to set the port use the -p flag (port 9090 is default).

`./tinysearch rundemo -g -p 8080`

### Downloader

The downloader uses the 3 flags to determine which CMS/SSG to use `-w` wordpress, `-j` jekyll, `-h` hugo (if none are specificed wordpress is assumed). From there it will build a corpus.json in the following format:

```json
[
  {
    "title": "Article 1",
    "url": "https://example.com/article1",
    "body": "This is the body of article 1."
  },
  {
    "title": "Article 2",
    "url": "https://example.com/article2",
    "body": "This is the body of article 2."
  }
]
```

This relies on the RootDomain configuration variable being set as it prepends it to the slugs it pulls from CMS/SSG. Downloader relies on persistor to save/load the json files to disk.

### Parser

The parser will take the downloaded corpus.json and convert it into a cuckoo dictionary with urls/names for the search tool.  The output is index.bin.  We can enable optimizations such as removing stopwords and/or non-alphanumeric content. Parser relies on persistor to save/load the bin files to disk.

### Persister

This will take the binary dictionary and persist it to an index.go file that allows it to be easily packaged with the final binary.  We convert the data struct to a byteslice that is embedded within the index.go file.  We can enable optimizations by compressing the byteslice with gzip prior to persisting.

### Build - this is not a module

This will build the final wasm binary and output it to the assets directory.  We can set optimization flags such as: 1) remove debugging information, 2) strip line numbers, 3) use tinygo.  We check if the wasm_exec.js is present and add it if not present.

### Summary

Checking the size of the resulting file: `stat -ln assets/tiny.wasm`

```bash
-rwxr-xr-x 1 brad staff 3259152 Sep 30 21:46:46 2020 assets/tiny.wasm
```

A 3.3MB wasm file is a big improvement (72% reduction) but certainly doesn't constitute as a 'tiny search'.  We can do better.  On the positive side the first load of the javascript is much snappier as there is no need to build the search index.  Also we removed the need to use packr to embded the external file by being crafty with byteslices and gobs.

Also interesting to note that corpus grew between pass 1 and 2, sitting now at 2.5MB.  The resulting stripped out dictionary using bloom filters with meta data is only 140KB.

Let's try optimizations.

## Third Pass

Finally we optimize by trying flag optimizations, general build tweaks, and/or replacing the main compiler with TinyGo.  The goal is to get the final tinysearch.wasm as small as possible.  The rust equivalent is 121kB or 51kB gzipped.  The optimizations are flagged in the build stage.

### Gzip Implementation

Implementing gzip content encoding for the server shrunk the tiny.wasm to 993KB (from 3.3MB) another 71% reduction in filesize.  To test for the right encoding:

```bash
go run main.go rundemo -g -p 8080
curl -v -sH "Accept-Encoding:gzip" localhost:8080/ | gunzip -
```

Next enabling gzip compression on the building of the index.go file reduced it from 485KB to 298KB.  With the layering on of the gzip compression in the server, the new file size is 985KB.  Not a massize savings overall but the hope is this will have a big impact the smaller we go.

### Stopwords, Remove Alpha-numerics and Stemming

With these optimizations the size of the index.bin was reduced from 140KB to 98KB.  The final tiny.wasm is 972KB.  Further clean up of the words in the dictionary, including removing duplicates and unecessary words such as HTML tags reduced the index.bin to 35KB on an original corpus.json of 1.7MB.  The tiny.wasm is now 944KB on a 3.2MB total size.  Disabling Stemming leads to a slightly larger index.bin (42KB) but leads to a better search experience.  Flag leaves the decision to the user to enable/disable.

### Binary optimizations

Final step is to optimize the binary by looking at flags, line number removals and other wasm tweaks.  By enabling `-ldflags="-s -w"` we are able to reduce the binary size by ~100KB (934KB after gzip compression):

```bash
-rwxr-xr-x 1 brad staff 3199938 Oct  4 13:40:29 2020 assets/tiny.wasm
```

This did net an improvement but still doesn't get us to the realm of tinysearch in rust (~150KB pre gzip).  I looked into other techniques to reduce binary size but none of them (e.g. stripping line numbers) could net us the binary size reduction needed.  So to get there our only choice is to move to the tinygo compiler.

### Tinygo

First round with tinygo - I used `encoding/gob` to encode the structures to binary and then embed into the wasm file. Installed tinygo for mac via brew and used the following to build the wasm file:

```bash
tinygo build -o assets/tiny.wasm -target wasm main.go
```

I discovered that gob is not supported in tinygo (as of 0.15.0) as many of the reflect features are not implemented.  I looked around to see if I can find another serialization technique that can decode without reflection - so far no luck.  To get there I decided to make my own serializer that optimizes for fast decode time and smaller output format (while only using reflection on the encoder side).  My inspiration is coming from 3 libraries: a) gotiny, b) sereal, c) msgpack.  But more importantly, I aimed for a simple encoder/decoder.

#### Terial Encoding

I call it Terial for tiny serial with the design philosphy to keep it simple and tiny.  We obviously are not 'complete' in order to keep it compact - many variable types are missing.

#### General Points

Strictness = must detect invalid documents without crash
Big Endian = as this protocol is intended to serialize data for transfer via a network, therefore we are using network order
IEEE Floats = floating points are in the IEEE format

#### Header

To encode the data without reflection we need to add relevant data in a header to infer what is encoded.  Using a variation of the Sereal header format combined with ASN.1 TLV + EOC data encoding.

```bash
Header = <encoding-format (32bit)><version (8bit)><num-of-fields (8bit)><user-meta-data (64bit)>
```

The header is always a fixed size (112bit) and is never compressed.

* encoding-format: this is a string that identifies the type of encoding-format.  Allows for a broader future change similar to Sereal with its magic string.  (0x54455231 or TER1)
* version: 1 indicates without compression, 2 with gzip compression, 3 with zstd compression
* num-of-fields: X indicates how many fields are encoded in the body from the originating struct
* user-meta-data: general meta data that the user wants to send along - in the case of tinysearch we send along the total number of elements in the arrays (as each array is equal size)

#### Body

Each variable encoded in the body is in an ASN.1 'like' TLNV+EOC format and it is assumed that whatever goes into the body is a `struct` or `map[string]` format:

```bash
Body Variable = <type-code(8bit)><length (int64)><value><name (string)><2 x end-of-content-octet (32bit)>
```

For the type-codes, we adopt a similar approach to Msgpack, the length is byteslice length, the value is the actual data to be encoded and the end-of-content octet.  Here are the codes used:

| Name    |  Code |
|:---------|:-----:|
| EOC     |  0xff |
| Nil     |  0xc0 |
| False   |  0xc1 |
| True    |  0xc2 |
| Bool    |  0xc3 |
| Float64 |  0xc4 |
| String  |  0xc5 |
| Uint8     |  0xd0 |
| []Uint8     |  0xd1 |
| Uint64   |  0xd2 |
| []Uint64    |  0xd3 |
| Int64    |  0xd4 |
| []Int64 |  0xd5 |
| Byte  |  0xd6 |
| []Byte  |  0xd7 |

To encode the string, Hello World!: `<0xcc><0x0C><0x48 0x70 ...><0xFFFFFFFF>` this is converted to a `[]byte` which is then converted to a `[]Uint8`.

## Final Structure

* assets --> everything served
* build --> output files from build steps
* cmd --> the go code
* -downloader --> extract information from wordpress, jekyll (TODO), hugo (TODO)
* -parser --> parse corpus.json into index.bin
* -persister --> marshal and write to different formats
* -serializer --> the terial encoder/decoder
* -server --> simple dev server
* -wasmgo --> code being compiled to wasm with native golang
* -fixtures --> the test files

## Release Approach

Using `github.com/nwillc/gorelease` we update the .version file with our target semver and then run `gorelease`.  This creates a `gen/version.go` file and will automatically tag and push to github.

To run this a convenient target is in the Makefile: `make release v0.2.2`.  The git repository needs to be clean prior to doing a release.

## Testing Approach

* go test ./... -v -cover
* go test ./... -coverprofile=test/coverage/c.out
* go tool cover -html=test/coverage/c.out
* go test -run ' '
* golangci-lint run

## References

1. https://endler.dev/2019/tinysearch
1. https://github.com/tinysearch/tinysearch
1. https://golangbot.com/webassembly-using-go/
1. https://www.stavros.io/posts/bloom-filter-search-engine/