// Echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		s = " "
	}

	fmt.Println(s)
}
