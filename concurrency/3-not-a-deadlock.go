package main
import "fmt"

func main() {
	mychannel1 := make(chan int)
	mychannel2 := make(chan int)
	mychannel3 := make(chan int)
	go func(){
		<-mychannel1
	}()

	go func(){
		mychannel2 <- 20
	}()

	go func(){
		<-mychannel3 
	}()


	fmt.Println(<-mychannel2)
}


// The program needs to be stuck as a whole for a deadlock to happen
// Here, first and third go routine are blocked, but the second one is not