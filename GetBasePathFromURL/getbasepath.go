package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	u, err := url.Parse("https://some.url.com/v1/id/f753e11b76fae9fd66984cf745e4383e/")
	if err != nil {
		fmt.Printf("error:%v", err)
	} else {
		p := path.Base(u.Path)
		fmt.Println(p)
	}
}
