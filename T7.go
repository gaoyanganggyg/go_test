package main

import (
	"os"
	"crypto/md5"
	"io"
	"fmt"
	"crypto/sha1"
)

func main()  {
	TestFile := "test.txt"
	f, err := os.Open(TestFile)
	if nil == err{
		md5h := md5.New()
		io.Copy(md5h, f)
		fmt.Printf("%x  %s\n", md5h.Sum([]byte("")), TestFile)

		sha1Inst := sha1.New()
		io.Copy(sha1Inst, f)
		fmt.Printf("%x  %s\n", sha1Inst.Sum([]byte("")), TestFile)
	}else {
		fmt.Println(err)
		os.Exit(1)
	}
}
