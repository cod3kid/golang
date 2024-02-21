package main
import "fmt"

func main() {
	mychannel := make(chan int, 2)
	mychannel <- 10
	fmt.Println(<-mychannel)
}