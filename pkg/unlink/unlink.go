package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	f, err := os.Create("temp")
	if err != nil {
		panic(err)
	}

	if err := unix.Unlink(f.Name()); err != nil {  // 进程退出后才实际删除，在此之前依然可以通过 f 访问
		panic(err)
	}

	_, err = f.Write([]byte("1234567890"))
	if err != nil {
		panic(err)
	}

	f.Seek(0, os.SEEK_SET)
	bytes, _ := ioutil.ReadAll(f)

	fmt.Println(string(bytes))
}

