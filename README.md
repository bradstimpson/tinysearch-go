# Tinysearch-Go

## Extract Data from Wordpress API

```bash
curl http://192.168.0.28:8009?rest_route=/wp/v2/posts\&page\=2\&per_page\=90 | jq '.[] |{title:.title.rendered,slug:.slug,content:.content.rendered}'
```

this put the data in a 'near' JSON form - I simply had to go through it and add commas plus array.

The resulting corpus file is 2.1MB.

## Copy over Javascript Glue Code

Run the following from within the project directory.

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets/
```

## Setup Structure

* assets --> everything served
* cmd --> the go code
* -server --> simple dev server
* -wasm --> code being compiled to wasm
* --fixtures --> the index file

## Webassembly

Compile the wasm script

```bash
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm
```

this outputs the wasm file that needs to be loaded by javascript

## First Pass

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

## Second Pass

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

With these optimizations the size of the index.bin was reduced from 140KB to 98KB.  The final tiny.wasm is 972KB

## Release Approach

Using `github.com/nwillc/gorelease` we update the .version file with our target semver and then run `gorelease`.  This creates a version.go file and will automatically tag and push to github.

## Testing Approach

go test ./... -cover
go test ./... -coverprofile=coverage/c.out
go tool cover -html=coverage/c.out
go test -run ' '
golangci-lint run

## References

1. https://endler.dev/2019/tinysearch
1. https://github.com/tinysearch/tinysearch
1. https://golangbot.com/webassembly-using-go/
1. https://www.stavros.io/posts/bloom-filter-search-engine/