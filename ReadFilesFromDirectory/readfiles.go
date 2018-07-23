package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("C:\\ReadFiles")
	if err == nil {
		for i, file := range files {
			fmt.Printf("%v file: %v \n", i, file.Name())
		}
	} else {
		fmt.Printf("Error:%v", err)
	}
}
