package main

import (
    "fmt"
    "os"
)

func main() {
    for i := 0; i < len(os.Args); i++ {
	fmt.Printf("Index: %v, Value: %v\n", i, os.Args[i])
    }
}
