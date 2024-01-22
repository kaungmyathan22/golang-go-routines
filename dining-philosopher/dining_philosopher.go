package diningphilosopher

import (
	"fmt"
	"sync"
	"time"
)

const hunger = 3

var eatTime = 3 * time.Second
var philophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second

func diningProblem(philopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	fmt.Println(philopher, " is seated")
	time.Sleep(sleepTime)
	for i := hunger; i > 0; i-- {
		fmt.Println(philopher, " is hungry.")
		time.Sleep(sleepTime)
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philopher)
		fmt.Println(philopher, "has both forks, and is eating.")
		time.Sleep(eatTime)
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philopher)
		time.Sleep(sleepTime)
	}
	fmt.Printf("%s is satisfied.\n.", philopher)
	time.Sleep(sleepTime)
	fmt.Println(philopher, " has left the table.")

}

func DemoDiningProblem() {
	// print intro

	wg.Add(len(philophers))
	forkLeft := &sync.Mutex{}
	// spawn one goroutine for each philsopher
	for i := 0; i < len(philophers); i++ {
		forkRight := &sync.Mutex{}
		// call a goroutine
		go diningProblem(philophers[i], forkLeft, forkRight)
	}

	wg.Wait()
}
