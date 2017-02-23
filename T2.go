package main

import(
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int)  {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf(w, "is ready\n")
	c <- 1
}

func main()  {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("i am waiting")
	<-c
	<-c
	//time.Sleep(5 * time.Second)

	fmt.Printf("aaaa")
}
