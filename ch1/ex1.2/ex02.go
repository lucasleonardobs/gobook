// Modify the 'echo' program to show the index and value of each of its arguments, one per line.
package echo

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
