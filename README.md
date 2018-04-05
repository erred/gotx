# gotx [![Build Status][1]][2] [![Go Report Card][3]][4] [![License: MIT][5]][6]
[1]: https://img.shields.io/travis/seankhliao/gotx.svg?style=flat-square
[2]: https://travis-ci.org/seankhliao/gotx
[3]: https://goreportcard.com/badge/github.com/seankhliao/gotx?style=flat-square
[4]: https://goreportcard.com/report/github.com/seankhliao/gotx
[5]: https://img.shields.io/badge/License-MIT-blue.svg?longCache=true&style=flat-square
[6]: LICENSE

go-template-exec

read STDIN -> parse and execute (text/template) -> print to STDOUT

or walk srcdir and output to destdir (ignores empty files!)

## Extra functions
### slice
define a slice in your template
```go
{{ $i := slice "a" "b" "c" }}
```

### obj
define a json object
```go
{{ $i := obj `{"hello": "world"}` }}
{{ index $i "hello" }}
```
### objs
define a json array
```go
{{ $i := obj `[{"hello": "world"}, {"this": "is"}, {"a": "message"}]` }}
{{ index (index $i 2) "a" }}
```

## Install
```sh
go get github.com/seankhliao/gotx
```

## Usage
```
gotx [--noecho]
```

```
gotx -src srcdir -out -destdir
```
