package main

import "fmt"

func mergeArrays(leftArr []int,rightArr []int)[]int{
	merged := make([] int, 0, len(leftArr) + len(rightArr))

	for len(leftArr) > 0 || len(rightArr)>0 {
		if len(leftArr) == 0{
			return append(merged,rightArr...)
		}else if len(rightArr) == 0{
			return append(merged,leftArr...)
		}else if leftArr[0] > rightArr[0]{
			merged = append(merged,rightArr[0])
			rightArr = rightArr[1:]
		}else{
			merged = append(merged,leftArr[0])
			leftArr = leftArr[1:]
		}

	}
	return merged
}

func mergeSort(arr []int) []int{
if(len(arr)<=1){
	return arr
}

done := make(chan bool)
doneRight := make(chan bool)

mid:= len(arr)/2

var left [] int 
var right [] int

go func(){
	left = mergeSort(arr[:mid])
	done <- true
}()

go func(){
	right = mergeSort(arr[mid:])
	doneRight <- true
}()



<- done
<- doneRight

return mergeArrays(left,right)
}


func main(){
data := [] int{9,4,3,6,1,2,10,5,7,8}
fmt.Printf("Intial Array: %v\nMerged Array: %v\n", data, mergeSort(data))
}