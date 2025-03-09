package main

import "fmt"

func contains(arr []int,num int)bool{
	for _,val:=range arr{
		if val==num{
			return true
		}
	}

	return false
}

func main(){
arr:=[]int{1,2,3}

result:=[][]int{}

var backtrack func([]int)

backtrack=func(current []int){
if len(current)==len(arr){
	currentCopy:=make([]int,len(current))
	copy(currentCopy,current)
	result = append(result,currentCopy)
	return
}


for _,val:=range arr{
	if !contains(current,val){
		current=append(current, val)
		backtrack(current)
		current = current[:len(current)-1]
	}
}

}


backtrack([]int{})

fmt.Println(result)
}