package main
import "fmt"

func main() {
    number := 0;
    
    go func(){
      number++ //reading and modifying the value of 'number'
    }()

    fmt.Println(number) //reading the value of 'number'


}