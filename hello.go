package main

import (
	"fmt"
)

func main() {
	const organizer = "Usman Khalil"
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to our booking application conference called", conferenceName, ". Be sure to enjoy")
	fmt.Printf("%v is responsible for managing all sort of stuff \n", organizer)
	fmt.Println("We have total of", conferenceTickets, "tickets and", "only", remainingTickets, "seats are available")

}
