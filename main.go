package main

import (
	"flag"
	"fmt"

	"github.com/showylee/gunzip/lib"
)

var (
	d string
)

func main() {
	fmt.Print("init pkg")
	flag.StringVar(&d, "d", "", "destination directory")
	flag.Parse()

	gunzip := lib.Gunzip{}
	gunzip.Src = flag.Arg(0)
	if d == "" {
		gunzip.D = false
	} else {
		gunzip.D = true
	}

	gunzip.Dest = d

	gunzip.Unzip()

}
