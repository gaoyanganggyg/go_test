package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
)

func main()  {
	TestString := "aaaaaa"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%d\n", len(string(Result)))
	fmt.Printf("%x\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%d\n", len(string(Result)))
	fmt.Printf("%x\n", Result)
}
