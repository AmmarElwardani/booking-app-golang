package helper

import "strings"

func ValidateUserInput(firstName string, lasttName string, email string, userTicket uint,remainingTickets uint)(bool, bool, bool){

	isValidName := len(firstName) >= 2 && len(lasttName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets 

	return isValidName, isValidEmail, isValidTicketNumber
}