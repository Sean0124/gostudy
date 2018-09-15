package echo

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
