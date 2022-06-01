// Modify the 'echo' program to show os.Args[0], which is the name of the command that called it
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}
