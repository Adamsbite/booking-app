package main

import "fmt"

func main(){ 
	fmt.Printf("Welcome to our conference booking application\n")
	fmt.Printf("Get your ticket here to attend\n")

   var name = "Go conference"
   var conferenceName = "Go conference"

   const conferenceTickets uint = 50 

   var remainingTickets uint = 50

   var firstName string
   var lastName string 
   var email string
   var userTickets uint

   fmt.Printf("%s\n", name)
   fmt.Printf("Welcome to %v booking application.\n" +
   "We have total of %v tickets and %v are still available.\n" +
   "Get your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
   


   fmt.Println( "please insert your name\n")
   fmt.Scanln(&firstName)

   fmt.Println( "please insert your lastName\n")
   fmt.Scanln(&lastName)

   fmt.Println( "please insert your email\n")
   fmt.Scanln(&email)

   fmt.Println( "please insert your userTicket\n")
   fmt.Scanln(&userTickets)

   fmt.Printf( "Thank you %v %v for booking %v tickets. you will receive a mail in %v\n", firstName, lastName, userTickets, email )
   

   remainingTickets = remainingTickets - userTickets
   fmt.Printf( "we now have %v remaining for %v", remainingTickets, name )
}

