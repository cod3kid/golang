package main

import (
        "errors"
        "fmt"
)

func main() {
        var errNotFound error = errors.New("Not found error")
        fmt.Println("Error:", errNotFound)
}

// error: Not found error