package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	f = flag.Bool("f", false, "quick create all sort file by filename")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: create [-f filename]")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		if *f {
			fmt.Fprintf(os.Stderr, "error: need at least one filename")
			os.Exit(2)
		}
	}

	for i := 0; i < flag.NArg(); i++ {
		name := flag.Arg(i)
		os.Mkdir(name, 0666)
		os.Chdir(name)
		os.Create(name + ".go")
		os.Create(name + "_test.go")
		os.Create("README.md")
	}
}
