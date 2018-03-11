package send

import (
	"errors"
	entities "prismapp-test/src/chat/entities"
	gateway "prismapp-test/src/chat/gateway"
)

// Service is an interface for domain sending
type Service interface {
	SendNewMessage(Message string) error
}

type send struct {
	roomGateway      gateway.Room
	websocketGateway gateway.Websocket
}

// NewService like a constructor for gate to access this service
func NewService(roomGateway gateway.Room, websocketGateway gateway.Websocket) Service {
	return &send{
		roomGateway:      roomGateway,
		websocketGateway: websocketGateway,
	}
}

func (s *send) SendNewMessage(message string) error {

	if message == "" {
		return errors.New("message cannot be empty")
	}

	s.websocketGateway.Publish(message)

	room := entities.Room{}
	room.Message = message

	err := s.roomGateway.Save(room)
	if err != nil {
		return err
	}

	return nil
}
