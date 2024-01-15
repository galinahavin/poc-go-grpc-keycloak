package main

import (
	"context"
	"log"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
)
func doGetTicketDetailsById(c pb.EcommerceServiceClient, ctx context.Context) {
	log.Println("doGetTicketDetailsById was invoked")
	res, err := c.GetTicketDetailsById(ctx, &pb.GetTicketDetailsRequest{TicketId: 3,})
	if err != nil {
		log.Fatal("Could not get ticket details: \n", err)
	}

	log.Printf("Response: %v\n", res.Ticket)	
}
func doSubmitPuchase(c pb.EcommerceServiceClient, ctx context.Context) {
	log.Println("doSubmitPuchase was invoked")

	ticket:= &pb.Ticket{

		UserId: 						3,		
		From:            		 "London",
		To:  		              "Paris",
		Price:						 20.0,
		Section:	  				  "A",
		Seat:		  					7,
	}

	res, err := c.SubmitPuchase(ctx, &pb.SubmitPuchaseRequest{UserId: 3, Ticket: ticket,})	
	if err != nil {
		log.Fatal("Could not place an order: \n", err)
	}

	log.Printf("Response: %s %s\n", res.User, res.Ticket)

}

func doGetUserSeatDetailsBySection(c pb.EcommerceServiceClient, ctx context.Context) {
	log.Println("doGetUserSeatDetailsBySection was invoked")
	res, err := c.GetUserSeatDetailsBySection(ctx, &pb.GetUserSeatDetailsBySectionRequest{Section: "A",})

	if err != nil {
		log.Fatal("Could not get details on the purchased tickets: \n", err)
	}

	log.Printf("Response: %s \n", res.PurchasedTickets)	
}

func doDeleteUserPurchase(c pb.EcommerceServiceClient, ctx context.Context) {

	log.Println("doDeleteUserPurchase was invoked")

	res, err := c.DeleteUserPurchase(ctx, &pb.DeleteUserPurchaseRequest{TicketId: 3,})
	if err != nil {
		log.Fatal("Could not delete purchased ticket: \n", err)
	}

	log.Printf("Response: %d\n", res.TicketId)

}

func doUpdateUserPuchase(c pb.EcommerceServiceClient, ctx context.Context) {
	log.Println("doUpdateUserPuchase was invoked")

	res, err := c.UpdateUserPuchase(ctx, &pb.UpdateUserPuchaseRequest{TicketId: 3, Seat: 3,})

	if err != nil {
		log.Fatal("Could not update the seat for the purchased ticket: \n", err)
	}

	log.Printf("Response: %v\n", res.Ticket)	
}
