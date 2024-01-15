package main

import (
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
)

// ticket manager
type TicketManager struct {
	tickets []pb.Ticket
}
var ticketsLock = &sync.Mutex{}
var lastTicketId int64 = 2

func newTicketId() int64 {
	return atomic.AddInt64(&lastTicketId, 1)
}


// returns a new ticket manager
func NewTicketManager() *TicketManager {
	return &TicketManager{}
}

func (TicketManager *TicketManager) addTicket(ticket pb.Ticket) {
	ticketsLock.Lock()
	defer ticketsLock.Unlock()
	ticket.TicketId = newTicketId()
	TicketManager.tickets = append(TicketManager.tickets, ticket)
}

func (TicketManager *TicketManager) deleteTicket(Id int64) {
	ticketsLock.Lock()
	defer ticketsLock.Unlock()
	for i, ticket := range TicketManager.tickets {
		if ticket.TicketId == Id {
			TicketManager.tickets = append(TicketManager.tickets[:i], TicketManager.tickets[i+1:]...)
			break
		}
	}
}

func (TicketManager *TicketManager) updateTicket(Id int64, Seat int32) {
	ticketsLock.Lock()
	defer ticketsLock.Unlock()
	for _, ticket := range TicketManager.tickets {
		if ticket.TicketId == Id {
			ticket.Seat = Seat
		}
	}
}

// return ticket by Id
func (TicketManager *TicketManager) ticketById(Id int64) (*pb.Ticket, error) {

	for _, ticket := range TicketManager.tickets {
		if ticket.TicketId == Id {
			return &ticket, nil
		}
	}
	return nil, errors.NotFound("not found: ticket %d", Id)
}

// return tickets by Section
func (TicketManager *TicketManager) seatDetailsBySection(ticketSection string) ([]*pb.PurchasedTicket, error) {
	var detailsBySection []*pb.PurchasedTicket

	for _, ticket := range TicketManager.tickets {
		if ticket.Section == ticketSection {
			purchasedTicket := &pb.PurchasedTicket{}

			user, _ := userManager.userById(ticket.UserId)
			purchasedTicket.Ticket = &ticket
			purchasedTicket.User = user			
			detailsBySection = append(detailsBySection, purchasedTicket)
		
		}
	}
	return detailsBySection, errors.NotFound("not found: tickets for section %s", ticketSection)
}