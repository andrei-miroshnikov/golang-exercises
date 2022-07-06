// print the name of a program and then all the arguments given

package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	for _, s := range os.Args[1:] {
		str += s + " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(str)
}
