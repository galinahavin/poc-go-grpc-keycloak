package main

import (
	"context"
	"log"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
)

var ticketManager = NewTicketManager()

func (s *Server) SubmitPuchase(ctx context.Context, req *pb.SubmitPuchaseRequest) (*pb.SubmitPuchaseResponse, error) {
	log.Printf("SubmitPuchase function was invoked with %v\n", req)

	UserId := req.UserId
	user, _ := userManager.userById(UserId)
	ticket := req.GetTicket()
	ticketManager.addTicket(ticket)

	return &pb.SubmitPuchaseResponse{
		User: &pb.User{
			UserId: 			 user.GetUserId(),
			FirstName:    	  user.GetFirstName(),
			LastName:  	       user.GetLastName(),
			EmailAddress:  user.GetEmailAddress(),
		},
		Ticket: &pb.Ticket{

			TicketId:         ticket.GetTicketId(),
			UserId: 	      user.GetUserId(),
			From:            ticket.GetFrom(),
			To:  		       ticket.GetTo(),
			Price:			ticket.GetPrice(),
			Section:	  ticket.GetSection(),
			Seat:		  ticket.GetSeat(),
		},
	}, nil
}

func (s *Server) GetTicketDetailsById(ctx context.Context, req *pb.GetTicketDetailsRequest) (*pb.GetTicketDetailsResponse, error) {
	log.Printf("GetTicketDetailsById function was invoked with %v\n", req)

	TicketId := req.TicketId
	ticket, _ := ticketManager.ticketById(TicketId)
	return &pb.GetTicketDetailsResponse{
		Ticket: ticket,
	}, nil
}

func (s *Server) GetUserSeatDetailsBySection(ctx context.Context, req *pb.GetUserSeatDetailsBySectionRequest) (*pb.GetUserSeatDetailsBySectionResponse, error) {
	log.Printf("GetUserSeatDetailsBySection function was invoked with %v\n", req)

	Section := req.Section
	details, _ := ticketManager.seatDetailsBySection(Section)
	return &pb.GetUserSeatDetailsBySectionResponse{PurchasedTickets: details}, nil
}
func (s *Server) DeleteUserPurchase(ctx context.Context, req *pb.DeleteUserPurchaseRequest) (*pb.DeleteUserPurchaseResponse, error) {
	log.Printf("DelDeleteUserPurchase function was invoked %v\n", req)
	id := req.TicketId
	ticketManager.deleteTicket(id)
	return &pb.DeleteUserPurchaseResponse{TicketId: id}, nil
}

func (s *Server) UpdateUserPuchase(ctx context.Context, req *pb.UpdateUserPuchaseRequest) (*pb.UpdateUserPuchaseResponse, error) {
	log.Printf("UpdateUserPuchase function was invoked %v\n", req)
	id := req.TicketId
	seat := req.Seat
	ticketManager.updateTicket(id, seat)
	ticket, _ := ticketManager.ticketById(id)
	return &pb.UpdateUserPuchaseResponse{
		Ticket: ticket,
	}, nil	
}

