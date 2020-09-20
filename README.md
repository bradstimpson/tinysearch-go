# Tinysearch-Go

## Extract Data from Wordpress API

```bash
curl http://192.168.0.28:8009?rest_route=/wp/v2/posts\&page\=2\&per_page\=90 | jq '.[] |{title:.title.rendered,slug:.slug,content:.content.rendered}'
```

this put the data in a 'near' JSON form - I simply had to go through it and add commas plus array.

The resulting fixtures file is 2.1MB.

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

An 11.8MB wasm file certainly doesn't constitute as a 'tiny search'.  We can do better.  Also to make matters worse, there is a major delay at first load to build the index.

Let's try separating the index building from the wasm build.

## Second Pass

In this refactoring we will simplify the directory structure, adopt the use of viper and cobra to build a functional parsing tool and finally output the built fixtures into a separate file that can be packed into the final wasm without bundling the large fixture file.

## Third Pass

Finally we optimize

### References

1. https://golangbot.com/webassembly-using-go/