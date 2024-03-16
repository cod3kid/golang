package main

import "fmt"

func main() {
	        // Go accepts Unicode characters as variable names
			// The Go type system will not allow dividing (or any other mathematical operation) between an integer (22) and a float (7.0).
			// But if it's 22/7.0, it'll allow
        var a, b = 22, 7.0
        var π = a / b
        fmt.Println(π)
}