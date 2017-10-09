package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print on newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()

	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}

		s += flag.Arg(i)
	}

	if *NewLine {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
