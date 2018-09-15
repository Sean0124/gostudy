package main

import (
	"awesomeProject/test/wightconv"
	"os"
	"strconv"
	"fmt"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}
		g := wightconv.G(t)
		kg := wightconv.Kg(t)
		fmt.Printf("%s = %s, %s = %s\n",
			g, wightconv.GToK(g), kg, wightconv.KToG(kg))
	}
}
