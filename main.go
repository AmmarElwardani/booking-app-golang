package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numbaerOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){


	greetingUsers()



	

		firstName, lasttName, email, userTicket := getUserInput()
		

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lasttName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
		
			bookTicket(userTicket, firstName, lasttName, email)
			wg.Add(1)
			go sendTicket(userTicket, firstName, lasttName, email)

			// displaying first names

			firstNames := getFirstNames()

			fmt.Printf("The firstNames of bookings: %v \n", firstNames)
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year!")
				//break
			}
		}else {	
			if !isValidName {
				fmt.Println("first name or last name might be too short")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number")
			}
		
		}

		wg.Wait()
	
}

func greetingUsers(){
	fmt.Printf("Weclome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v Tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}
func getFirstNames() []string{
	firstNames :=  [] string{}

	for _, booking := range(bookings){
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}



func getUserInput() (string, string, string, uint){

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
	
	return firstName, lasttName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lasttName string, email string){
	remainingTickets = remainingTickets - userTicket
	
	var userData = UserData {
		firstName: firstName,
		lastName: lasttName,
		email: email,
		numbaerOfTickets: userTicket,
	}
	
	
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve a confirmation email at %v \n",firstName, lasttName, userTicket, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("###############")

	wg.Done()
}