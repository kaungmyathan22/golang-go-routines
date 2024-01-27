package sleepingbarbar

import (
	"fmt"
	"strings"
	"time"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s !!!", strings.ToUpper(s))
	}
}

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}
func RunSleepingBarBar() {
	// ping := make(chan string)
	// pong := make(chan string)

	// go shout(ping, pong)
	// fmt.Println("Type something and press ENTER (enter Q to quit)")
	// for {
	// 	fmt.Print("-> ")
	// 	var userInput string
	// 	_, _ = fmt.Scanln(&userInput)
	// 	if userInput == strings.ToLower("q") {
	// 		break
	// 	}
	// 	ping <- userInput
	// 	response := <-pong
	// 	fmt.Println("Response : ", response)
	// }
	// fmt.Println("All done")
	// close(ping)
	// close(pong)
	fmt.Println("Select with channels")
	fmt.Println("--------------------")
	channel1 := make(chan string)
	channel2 := make(chan string)
	go server1(channel1)
	go server2(channel2)
	for {
		select {
		case s1 := <-channel1:
			fmt.Println("case 1: ", s1)
		case s2 := <-channel1:
			fmt.Println("case 2: ", s2)
		case s3 := <-channel2:
			fmt.Println("case 3: ", s3)
		case s4 := <-channel2:
			fmt.Println("case 4: ", s4)
		}
	}
}
