package main

import "fmt"

func main() {
    fmt.Println("Starting the program...")
    safeFunction()
    fmt.Println("Program continues after recovery.")
}

func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    // This will cause a panic
    panic("Something went wrong!")
}
