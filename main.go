package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/taigrr/temper"
)

func main() {
	fahrenheit := flag.Bool("f", false, "output temperature in Fahrenheit")
	flag.Parse()

	tempers, err := temper.FindTempers()
	if err != nil {
		log.Fatalf("error finding temper devices: %v", err)
	}
	if len(tempers) == 0 {
		fmt.Fprintln(os.Stderr, "no temper devices found")
		os.Exit(1)
	}
	defer func() {
		for _, t := range tempers {
			t.Close()
		}
	}()

	for _, t := range tempers {
		c, err := t.ReadC()
		if err != nil {
			log.Fatalf("error reading temperature: %v", err)
		}
		if *fahrenheit {
			fmt.Printf("%.2f\n", c*9.0/5.0+32.0)
		} else {
			fmt.Printf("%.2f\n", c)
		}
	}
}
