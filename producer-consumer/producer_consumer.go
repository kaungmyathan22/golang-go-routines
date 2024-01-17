package producer_consumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d secodns....\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			msg = fmt.Sprintf("**** We ran out of ingredients for pizza #%d!\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("**** The cook quit while making pizza #%d!\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func RunProducerConsumerSample() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("The pizzeria is open for business!")
	color.Cyan("----------------------------------")
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzeria(pizzaJob)
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}
	color.Cyan("----------------------------------")
	color.Cyan("Don for the day")
	color.Cyan(fmt.Sprintf("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzaFailed, total))
}
