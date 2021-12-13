package main

import (
	"log"
	"os"
)

func main() {
	var f *os.File
	var err error
	if _, err = os.Stat("data.txt"); err == nil {
		f, err = os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	} else {
		f, err = os.Create("data.txt")
	}
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}
