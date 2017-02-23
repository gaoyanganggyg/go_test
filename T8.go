package main

import (
	"reflect"
	"fmt"
)

func A()  {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(8.88)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type T struct {
	A int
	B string
}

func C()  {
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func B()  {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	fmt.Println(reflect.TypeOf(x))
}

func main()  {
	C()
}
