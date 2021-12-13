package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	const name, age = "Johne Doe", 34

	n, err := fmt.Fprintln(f, name, "is", age, "years old.")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(n, "bytes written")
	fmt.Println("done")
}
