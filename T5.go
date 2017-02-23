package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	r, err := http.Get("http://www.baidu.com")
	if nil != err {fmt.Printf("aaaaaa"); return }
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if nil == err {fmt.Printf("%s", string(b))}
}