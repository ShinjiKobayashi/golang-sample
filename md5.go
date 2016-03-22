package main

import (
	"crypto/md5"
	"io/ioutil"
	"fmt"
)

func main() {
	in := "0e8cb623373fc99f5a92531a2be5668b"
	b, _ := ioutil.ReadFile("exec.go")
	md := md5.Sum(b)

	if fmt.Sprintf("%x", md) == in {
		fmt.Print("equal")
	} else {
		fmt.Print("Different")
	}
}
