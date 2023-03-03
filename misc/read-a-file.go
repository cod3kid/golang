package main

import (
	"os"
    "fmt"
	"io/ioutil"
)

func main2(){

	file, err := os.Open("dummy.txt")
	if err != nil{
		return
	}

	defer file.Close()

	// Reading File Stats
	stats, err := file.Stat()
	if err != nil{
		return 
	}

	fmt.Printf("%+v\n",stats)

	contents, err := ioutil.ReadFile("dummy.txt")

	if err!= nil{
		return
	}

	fmt.Println("\n",string(contents))



}