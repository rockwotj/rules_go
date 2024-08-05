package main

import (
	"fmt"
	"log"

	"example.com/lib"
)

type basicStruct struct {
	foo int
}

var _ = basicStruct{}

func main() {
	if main != nil {
		fmt.Printf("Hello %s!", lib.Name())
	}
	_ = !(true && false)
	maps := make(map[string]string)
	for k, _ := range maps {
		log.Println(k)
	}
	maps = maps
}
