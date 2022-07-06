// print the index and value of each argument, 1 arg per line

package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, s := range os.Args[1:] {
		fmt.Println(ind, s)
	}
}
