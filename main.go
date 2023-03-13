package main

import (
	"fmt"
	"log"

	"github.com/taigrr/temper"
)

var tempers []*temper.Temper

func main() {
	tempers, err := temper.FindTempers()
	if err != nil {
		panic(err)
	}
	if len(tempers) == 0 {
		log.Fatal("no tempers found\n")
	}
	for _, t := range tempers {
		defer t.Close()
	}
	for _, t := range tempers {
		c, err := t.ReadC()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%.2f\n", c)
	}
}
