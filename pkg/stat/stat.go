package main

import (
	"flag"
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	var file string

	flag.StringVar(&file, "f", "", "file name")
	flag.Parse()

	var stat unix.Stat_t
	if err := unix.Stat(file, &stat); err != nil {
		panic(err)
	}

	typeMapping := map[uint32]string{
		unix.S_IFBLK:  "block",
		unix.S_IFCHR:  "char",
		unix.S_IFDIR:  "dir",
		unix.S_IFIFO:  "fifo",
		unix.S_IFLNK:  "link",
		unix.S_IFSOCK: "sock",
		unix.S_IFREG:  "file",
	}
	mode := stat.Mode & unix.S_IFMT

	fmt.Printf("file %s type is %s(%d)\n", file, typeMapping[mode], mode)
	fmt.Printf("file setuid %v\n", (stat.Mode&unix.S_ISUID) != 0)
	fmt.Printf("file setgid %v\n", (stat.Mode&unix.S_ISGID) != 0)
	fmt.Printf("file device number %d/%d\n", unix.Major(stat.Dev), unix.Minor(stat.Dev))
	fmt.Printf("file stat %+v\n", stat)
}
