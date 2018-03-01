package main // import "github.com/teamwork/kommentaar"

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kr/pretty"
	"github.com/teamwork/kommentaar/docparse"
	"github.com/teamwork/kommentaar/finder"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: kommentaar [pkg pkg...]\n\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}

	docparse.InitProgram()

	//err := finder.FindComments(paths, docparse.Parse, openapi3.Write)
	err := finder.FindComments(paths, docparse.Parse, dumpOut)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}

func dumpOut(_ io.Writer, prog docparse.Program) error {
	_, err := pretty.Print(prog)
	fmt.Println("")
	return err
}
