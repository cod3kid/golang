package main
import ( "fmt"
        "sync")

func deposit(balance *int,amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup){
    myMutex.Lock()
    *balance += amount
    myMutex.Unlock()
    myWaitGroup.Done()
}

func withdraw(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup){
    myMutex.Lock()
    *balance -= amount 
    myMutex.Unlock()
    myWaitGroup.Done()
}

func main() {
  
    balance := 100 
    var myWaitGroup sync.WaitGroup
    var myMutex sync.Mutex

    myWaitGroup.Add(2)
    go deposit(&balance,10, &myMutex, &myWaitGroup) //depositing 10
    withdraw(&balance, 50,&myMutex, &myWaitGroup) //withdrawing 50

    myWaitGroup.Wait()
    fmt.Println(balance) 


}