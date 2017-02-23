package main

import (
	"fmt"
)

func main()  {
	var p *int
	fmt.Printf("%v\n", p)

	var  i int
	p = &i
	fmt.Printf("%v\n", p)

	*p = 9
	*p++
	fmt.Printf("%v\n", *p)
	*p++
	fmt.Printf("%v\n", i)

	type SyncdBuffer struct {
		a int
		b string
	}

	//var a *SyncdBuffer = new(SyncdBuffer)
	v := SyncdBuffer{1, "ssss"}
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", v.a)
	fmt.Printf("%s\n", v.b)

}