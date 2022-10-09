package main
import (
    "strings")

func validateUserInput(firstName string, lastName string, email string ,userTickets uint, remainingTickets uint) (bool, bool, bool) { 
    isValidName := len(firstName) >= validNameLen && len(lastName) >= validNameLen
    isValidEmail := strings.Contains(email, "@") && len(email) >= validEmailLen
    isValidTicket := userTickets > 0 && userTickets <= remainingTickets

    return isValidName, isValidEmail, isValidTicket
}