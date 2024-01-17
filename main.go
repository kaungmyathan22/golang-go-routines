package main

import producer_consumer "github.com/kaungmyathan22/golang-go-routine/producer-consumer"

// func printSomething(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(s)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	words := []string{
// 		"alpha",
// 		"beta",
// 		"delta",
// 		"gamma",
// 		"pi",
// 		"zeta",
// 		"eta",
// 		"theta",
// 		"epsilon",
// 	}
// 	wg.Add(len(words))
// 	for i, word := range words {
// 		go printSomething(fmt.Sprintf("%d : %s", i, word), &wg)
// 	}
// 	wg.Wait()
// 	wg.Add(1)
// 	printSomething("This is the second thing to be printed.", &wg)
// }

// var wg sync.WaitGroup

// func main() {

// 	// challenge: modify this code so that the calls to updateMessage() on lines
// 	// 28, 30, and 33 run as goroutines, and implement wait groups so that
// 	// the program runs properly, and prints out three different messages.
// 	// Then, write a test for all three functions in this program: updateMessage(),
// 	// printMessage(), and main().

// 	msg = "Hello, world!"

// 	wg.Add(1)
// 	go updateMessage("Hello, universe!", &wg)
// 	wg.Wait()
// 	printMessage()

// 	wg.Add(1)
// 	go updateMessage("Hello, cosmos!", &wg)
// 	wg.Wait()
// 	printMessage()

//		wg.Add(1)
//		go updateMessage("Hello, world!", &wg)
//		wg.Wait()
//		printMessage()
//	}

// func main() {
// 	// racecondition.RaceConditionMain()
// }

// func main() {
// 	mutex.RunMutextSample()
// }

func main() {
	producer_consumer.RunProducerConsumerSample()
}
