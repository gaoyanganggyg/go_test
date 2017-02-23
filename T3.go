package main

import(
	"os"
	"fmt"
	"bufio"
)

func useBuff()  {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		s, _ := r.Read(buf)
		if 0 == s {break}
		w.Write(buf[:s])
	}
}

func unBuff()  {
	buf := make([]byte, 1024)
	f, err := os.Open("/etc/passwd")
	fmt.Printf("%s", err)
	defer f.Close()
	for{
		n, _ := f.Read(buf)
		if 0 == n {break}
		os.Stdout.Write(buf[:n])

	}
}


func main()  {
	useBuff()
}
