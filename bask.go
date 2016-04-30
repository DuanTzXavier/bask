//bask is a tool for developing android application
package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

const version = "0.1"

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		//		help(args[1:])
		return
	}

	fmt.Fprintf(os.Stderr, "help\n")
	//	usage()

	os.Exit(2)
}

var usageTemplate = `Bask is a tool for managing android application.

Usage:

	bask command [arguments]
	
The commands are:

Use "bask help [command]" for more information about a command.

Additional help topics:

Use "bask help [topic]" for more information about that topic.

`

func usage() {
	tmpl(os.Stdout, usageTemplate)
	os.Exit(2)
}

func tmpl(w io.Writer, text string) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": func(s template.HTML) template.HTML {
		return template.HTML(strings.TrimSpace(string(s)))
	}})
	template.Must(t.Parse(text))
	if err := t.Execute(w, nil); err != nil {
		panic(err)
	}
}
