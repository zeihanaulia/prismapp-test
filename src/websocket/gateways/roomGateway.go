package gateway

import entities "prismapp-test/src/websocket/entities"

// Room gateway
type Room interface {
	Save(message entities.Room) error
	RetrieveAllPreviousMessages(Page, Limit int) (*[]entities.Room, error)
	RetrieveAllPreviousMessagesCount() (int, error)
}
