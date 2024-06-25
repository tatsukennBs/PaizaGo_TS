//go:build ignore

package main

import (
	"fmt"
)

func main() {

	kvs := map[string]string{"b": "banana", "c": "cheese", "g": "grape"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}
}
