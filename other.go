package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("./main.go", []byte(`package main

func main() {
	
}`), 0755)

	if err != nil {
		log.Fatal(err)
	}
}
