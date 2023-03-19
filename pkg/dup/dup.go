package main

import (
	"flag"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	var dup2 string

	flag.StringVar(&dup2, "to", "stdout.out", "dup stdout to file.")
	flag.Parse()

	f, err := os.Create(dup2)
	if err != nil {
		panic(err)
	}

	unix.Dup2(int(f.Fd()), int(os.Stdout.Fd()))

	io.Copy(os.Stdout, os.Stdin)
}
