package echo

import (
	"os"
	"fmt"
)

func main() {
	s, sep := "", ""
	for num, arg := range os.Args[1:] {
		fmt.Println(num, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
