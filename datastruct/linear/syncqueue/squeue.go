package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass         int
	waitTicket       int
	playPass         bool
	playTicket       bool
	queuePass        chan int
	queueTicket      chan int
	message          chan int
	handle           func() error
	ticketsSold      int
	passengersServed int
}

func (q *Queue) New() {
	q.message = make(chan int)
	q.queuePass = make(chan int)
	q.queueTicket = make(chan int)
	q.handle = func() error {
		var message int
		for {
			if q.passengersServed >= 15 {
				fmt.Println("passengers served", q.passengersServed)
				fmt.Println("END")
				return fmt.Errorf("Bus is full")
			}
			fmt.Println("tickets", q.waitTicket, q.ticketsSold, "passengers", q.waitPass, q.passengersServed)
			select {
			case message = <-q.message:
				switch message {
				case messagePassStart:
					q.waitPass++
				case messagePassEnd:
					q.playPass = false
				case messageTicketStart:
					q.waitTicket++
				case messageTicketEnd:
					q.playTicket = false
				}
				if q.waitPass > 0 && q.waitTicket > 0 &&
					!q.playPass && !q.playTicket {
					q.playPass = true
					q.playTicket = true
					q.waitTicket--
					q.waitPass--
					q.queuePass <- 1
					q.queueTicket <- 1
				}
			}
		}
	}
}

func (q *Queue) StartTicketIssue() {
	q.message <- messageTicketStart
	<-q.queueTicket
}

func (q *Queue) EndTicketIssue() {
	q.message <- messageTicketEnd
}

func ticketIssue(q *Queue) {
	for {
		fmt.Println("starting Ticket Queue")
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		q.StartTicketIssue()
		fmt.Println("Ticket Issue Starts")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Ticket Issue Ends")
		q.EndTicketIssue()
		q.ticketsSold++
	}
}

func (q *Queue) StartPass() {
	q.message <- messagePassStart
	<-q.queuePass
}

func (q *Queue) EndPass() {
	q.message <- messagePassEnd
}

func passenger(q *Queue) {
	fmt.Println("starting passenger Queue")
	for {
		fmt.Println("starting the processing")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		q.StartPass()
		fmt.Println("Passenger starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Passenger Ends")
		q.EndPass()
		q.passengersServed++
	}
}

func main() {
	var q *Queue = &Queue{}
	q.New()
	fmt.Println(q)
	// var i int
	// for i = 0; i < 10; i++ {
	// fmt.Println(i, "passenger issued in the Queue")

	go passenger(q)
	// }
	// var j int
	// for j = 0; j < 5; j++ {
	// fmt.Println(i, "ticket issued in the Queue")
	go ticketIssue(q)
	// }
	err := q.handle()
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
