package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var mutex sync.Mutex

	fmt.Printf("Initial balance %d.00", bankBalance)

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Smuggling", Amount: 20000},
		{Source: "Drug Dealing", Amount: 50000},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income){
			defer wg.Done()
			for week:=1;week<=52;week++{
				mutex.Lock()
				temp:=bankBalance
				temp+=income.Amount
				bankBalance=temp
				mutex.Unlock()

				fmt.Printf("On week %d, you earned %d",week,income.Amount)
			}
		}(i,income)
	}


	wg.Wait()

	fmt.Printf("Final balance %d.00", bankBalance)
}
