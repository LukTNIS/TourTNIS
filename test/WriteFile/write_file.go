package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	val := "old falcon\n"
	data := []byte(val)

	err := ioutil.WriteFile("data.txt", data, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
