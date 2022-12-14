package helper

import (
    "strings"
    )

    const validNameLen int = 2
    const validEmailLen = 5

//Export with function name start with Captial letter
func ValidateUserInput(firstName string, lastName string, email string ,userTickets uint, remainingTickets uint) (bool, bool, bool) { 
    isValidName := len(firstName) >= validNameLen && len(lastName) >= validNameLen
    isValidEmail := strings.Contains(email, "@") && len(email) >= validEmailLen
    isValidTicket := userTickets > 0 && userTickets <= remainingTickets

    return isValidName, isValidEmail, isValidTicket
}