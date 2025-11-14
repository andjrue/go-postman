package main

import (
	"fmt"

	"github.com/andjrue/go-postman/internal/collections"
)

func main() {
	coll, err := collections.LoadFile()
	if err != nil {
		panic(err)
	}
	
	pp, err := collections.PrettyPrint(coll)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("config:\n%v", pp)
}