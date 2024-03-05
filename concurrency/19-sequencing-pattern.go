package main

import ( "fmt"
		  "math/rand"
		  "time")


type CookInfo struct {
	foodCooked  string
	waitForPartner chan bool 
}



func cookFood(name string) <-chan CookInfo { 
	
	cookChannel := make(chan CookInfo)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			cookChannel<- CookInfo{fmt.Sprintf("%s %s", name,"Done") , wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-wait
		}
	}()

	return cookChannel
}

func fanIn(mychannel1, mychannel2 <-chan CookInfo) <-chan CookInfo {
	mychannel := make(chan CookInfo)

	go func() { 
		for {
			mychannel <- <-mychannel1 
		}
	}()

	go func() { 
		for {
			mychannel <- <-mychannel2
		}
	}()

	return mychannel
}


func main() {
	gameChannel := fanIn(cookFood("Player 1 : "), cookFood("Player 2 :"))
	

	for round := 0; round < 3; round++ {
		food1 := <-gameChannel
		fmt.Println(food1.foodCooked)

		food2 := <-gameChannel
		fmt.Println(food2.foodCooked)

		food1.waitForPartner <- true
		food2.waitForPartner <- true

		fmt.Printf("Done with round %d\n", round+1)
	}

	fmt.Println("Done with the competition")
}

