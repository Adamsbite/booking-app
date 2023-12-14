package main

import ("fmt"
		"strings"
)
func main(){ 
	fmt.Printf("Welcome to our conference booking application\n")
	fmt.Printf("Get your ticket here to attend\n")

   var name = "Go conference"
   var conferenceName = "Go conference"
   var remainingTickets uint = 50

   const conferenceTickets uint = 50 
   fmt.Printf("%s\n", name)
	fmt.Printf("Welcome to  %v booking application.\n" +
	"We have total of %v tickets and %v are still available.\n" +
	"Get your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
	



   var bookings []string 


   for {
	var firstName string
	var lastName string 
	var email string 
	var userTickets uint 
 
	
	fmt.Println( "please insert your name\n")
	fmt.Scanln(&firstName)
 
	fmt.Println( "please insert your lastName\n")
	fmt.Scanln(&lastName)
 
	fmt.Println( "please insert your email\n")
	fmt.Scanln(&email)
 
	fmt.Println( "please insert your userTicket\n")
	fmt.Scanln(&userTickets)
 
	
	remainingTickets = remainingTickets - userTickets
 
 //    bookings [0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
 
 
	fmt.Printf("the whole array: %v \n", bookings)
	fmt.Printf("the first  value: %v \n", bookings[0])
	fmt.Printf("the array array: %T \n", bookings)
 
	fmt.Printf( "Thank you %v %v for booking %v tickets. you will receive a mail in %v\n", firstName, lastName, userTickets, email )
 
	fmt.Printf( "we now have %v remaining for %v \n", remainingTickets, name )
	
	  firstNames := []string {} 
	for _, booking := range bookings{
		var names = strings.Fields(booking)
		// var firstName = names [0] 
		firstNames = append(firstNames, names[0])
	}
	fmt.Print("The first names of bookings are %v", firstNames)
   }
  
}

