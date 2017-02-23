package main

import(
	"fmt"
	"os/exec"
)

func t1()  {
	cmd := exec.Command("/bin/ls", "-la")
	err := cmd.Run()
	fmt.Printf("%s", err)
}

func t2()  {
	cmd := exec.Command("/bin/ls", "-la")
	buf, err := cmd.Output()
	fmt.Printf("%s", buf)
	fmt.Printf("%s", err)
}

func main()  {
	t2()
	
}
