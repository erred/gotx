package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type options struct {
	NoEcho bool
	SrcDir string
	OutDir string
}

func main() {
	opts := options{}
	flag.BoolVar(&opts.NoEcho, "noecho", false, "do not echo input first in STDIN/STDOUT mode")
	flag.StringVar(&opts.SrcDir, "src", "", "source dir to walk recursively")
	flag.StringVar(&opts.OutDir, "out", "", "output dir")
	flag.Parse()

	if opts.SrcDir == "" && opts.OutDir == "" {
		dirmode(opts)
		return
	}
	stdiomode(opts)

}

func stdiomode(opts options) {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error reading input: ", err)
	}

	t, err := template.New("").Funcs(
		template.FuncMap(
			map[string]interface{}{
				"slice": func(args ...interface{}) []interface{} {
					return args
				},
			})).Parse(string(b))
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	if !opts.NoEcho {
		_, err = fmt.Print(string(b))
		if err != nil {
			log.Fatal("Error echoing input")
		}
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}

func dirmode(opts options) {
	tmpls := template.New("").Funcs(
		template.FuncMap(
			map[string]interface{}{
				"slice": func(args ...interface{}) []interface{} {
					return args
				},
			}))
	dirs := []string{}

	filepath.Walk(opts.SrcDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		name, err := filepath.Rel(opts.SrcDir, p)
		if err != nil {
			log.Fatal("error converting path to relative: ", err)
		}
		name = path.Join(opts.OutDir, name)

		if info.IsDir() {
			dirs = append(dirs, name)
			return nil
		}

		b, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatal("error reading file: ", err)
		}

		_, err = tmpls.New(name).Parse(string(b))
		if err != nil {
			log.Fatal("error parsing template: ", err)
		}
		return nil
	})

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal("error creating dirs: ", err)
		}
	}

	for _, tmpl := range tmpls.Templates() {
		if d, _ := filepath.Split(tmpl.Name()); d == "" {
			// ignore paths without dir
			continue
		}

		b := bytes.NewBuffer([]byte{})
		err := tmpl.Execute(b, nil)
		if err != nil {
			log.Fatal("error executing template: ", err)
		}

		if strings.TrimSpace(b.String()) == "" {
			continue
		}
		err = ioutil.WriteFile(tmpl.Name(), b.Bytes(), 0644)
		if err != nil {
			log.Fatal("error writing file", err)
		}
	}
}
