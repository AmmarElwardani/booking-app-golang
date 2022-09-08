package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Weclome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v Tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var bookings = []string{}

	for {

		var firstName string
		var lasttName string
		var email string
		var userTicket uint

		// ask for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lasttName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTicket)

		if userTicket <= remainingTickets {
		
			remainingTickets = remainingTickets - userTicket
			
			bookings = append(bookings, firstName + " " + lasttName)

			fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve a confirmation email at %v \n",firstName, lasttName, userTicket, email)

			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames :=  [] string{}

			for _, booking := range(bookings){
				var names = strings.Fields(booking)
				
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The firstNames of bookings: %v \n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		}else {	
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n",remainingTickets, userTicket)
			continue
		}

	}
}