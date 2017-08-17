package main

import (
	"fmt"
	"math/bits"
)

//START OMIT
func main() {
	a := uint8(0x88)
	b := bits.RotateLeft8(a, 2)
	fmt.Printf("a: %8.b\nb: %8.b", a, b)
}

//END OMIT
