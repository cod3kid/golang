package main

import (
	"fmt"
)

func main() {
	fmt.Println(0x1p-2)
}


/*
0x: Indicates the number is in hexadecimal format.
1: The significand (or mantissa) in base 16 (hexadecimal). This is equivalent to 1 in decimal.
p: Indicates the power of 2 (exponent) in scientific notation for floating-point numbers.
-2: The exponent

1*2^-2 = 0.25
*/