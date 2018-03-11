package retrieve

import (
	entities "prismapp-test/src/chat/entities"
	gateway "prismapp-test/src/chat/gateway"
)

// Service is an interface for domain sending
type Service interface {
	RetrievePreviousMessage(Page, Limit int) (*[]entities.Room, int, error)
}

type retrieve struct {
	roomGateway gateway.Room
}

// NewService like a constructor for gate to access this service
func NewService(roomGateway gateway.Room) Service {
	return &retrieve{
		roomGateway: roomGateway,
	}
}

func (r *retrieve) RetrievePreviousMessage(Page, Limit int) (*[]entities.Room, int, error) {

	messages, err := r.roomGateway.RetrieveAllPreviousMessages(Page, Limit)
	if err != nil {
		return nil, 0, err
	}

	count, err := r.roomGateway.RetrieveAllPreviousMessagesCount()
	if err != nil {
		return nil, 0, err
	}

	return messages, count, nil
}
