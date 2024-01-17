package mutex

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func RunMutextSample() {
	var bankBalance int
	var balance sync.Mutex
	startTime := time.Now()
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(ii int, income Income) {
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
			wg.Done()
		}(i, income)
	}
	wg.Wait()
	duration := time.Since(startTime)
	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
	fmt.Printf("Execution Duration: %f seconds\n", duration.Seconds())
	fmt.Println()
}
