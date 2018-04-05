# gotx [![Build Status][1]][2] [![Go Report Card][3]][4] [![License: MIT][5]][6]
[1]: https://img.shields.io/travis/seankhliao/gotx.svg?style=flat-square
[2]: https://travis-ci.org/seankhliao/gotx
[3]: https://goreportcard.com/badge/github.com/seankhliao/gotx?style=flat-square
[4]: https://goreportcard.com/report/github.com/seankhliao/gotx
[5]: https://img.shields.io/badge/License-MIT-blue.svg?longCache=true&style=flat-square
[6]: LICENSE

go-template-exec

provides the function `slice` to create slices

read STDIN -> parse and execute (text/template) -> print to STDOUT

or walk srcdir and output to destdir (ignores empty files!)


## Install
```sh
go get github.com/seankhliao/gotx
```

## Usage
#### Input 
```sh
gotx
{{ range slice "a" "b" "c" }}
    {{ . }}
{{ end }}
```
#### Output
```
{{ range slice "a" "b" "c" }}
    {{ . }}
{{ end }}

    a


    b


    c

```

#### Input
```sh
gotx -src src -out dest 
```
