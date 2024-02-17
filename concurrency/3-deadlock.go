package main
import "fmt"

func main() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

// we can avoid this deadlock if we put the send operation in a goroutine
// such that both the send/receive operations are ready for each other simultaneously.

// go func(){
// 	mychannel <- 10
// }()