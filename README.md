# Tinysearch-Go

![GitHub Workflow Status](https://github.com/bradstimpson/tinysearch-go/workflows/CI/badge.svg)
[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bradstimpson/tinysearch-go?tab=doc)
[![codecov](https://codecov.io/gh/bradstimpson/tinysearch-go/branch/master/graph/badge.svg)](https://codecov.io/gh/bradstimpson/tinysearch-go)
[![Go Report Card](https://goreportcard.com/badge/bradstimpson/tinysearch-go)](https://goreportcard.com/report/bradstimpson/tinysearch-go)

**Tinysearch-go** is a set of command-line tools that can create a **tiny webassembly module** of your site's posts.  Once embedded in your static website users will be able to search your posts quickly and easily.  This is heavily inspired by tinysearch developed at endler.dev but instead of using Rust, we use Golang.

- Website: tbd
- Slack: [gophers.slack.com](https://gophers.slack.com) **#tinygo** ([invite](https://gophersinvite.herokuapp.com/))

## Sponsors

Click on sponsor, above, for more information on sponsorship.

## Goals

- webassembly module that is as small as possible (ideally less than 150kB)
- set of tools to build the wasm module from querying your static site's posts to the final output of `tiny.wasm`
- leverage all learnings from Endler.dev's approach in Rust but using Golang
- build a complete set of unit and integration tests to automate the development approach

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md).
Thank you, [contributors](https://github.com/bradstimpson/tinysearch-go/graphs/contributors)!

## Related Projects

1. [Rust tinysearch github](https://github.com/tinysearch/tinysearch)
2. [Jekyll REST API](https://github.com/riichard/jekyll-rest-api)

## Additional Reading

1. [Endler.dev's write-up](https://endler.dev/2019/tinysearch)

