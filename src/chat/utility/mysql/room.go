package mysql

import (
	entities "prismapp-test/src/chat/entities"
	gateway "prismapp-test/src/chat/gateway"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type room struct {
	db *gorm.DB
}

func NewRoomGateway(db *gorm.DB) gateway.Room {
	return &room{db: db}
}

// Save message to db
func (m *room) Save(room entities.Room) error {
	uid, _ := uuid.NewV4()
	err := m.db.Exec("INSERT room(id ,message) VALUES(?,?) ", uid.String(), room.Message).Error
	return err
}

func (m *room) RetrieveAllPreviousMessages(page, limit int) (*[]entities.Room, error) {

	var id, message string
	var createdAt time.Time

	offset := (limit * (page - 1))
	rows, err := m.db.Raw("select id, message, created_at from room order by created_at limit ?,?", offset, limit).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []entities.Room{}
	for rows.Next() {
		rows.Scan(&id, &message, &createdAt)

		room := entities.Room{
			ID:        id,
			Message:   message,
			CreatedAt: createdAt,
		}

		rooms = append(rooms, room)
	}

	return &rooms, nil
}

func (m *room) RetrieveAllPreviousMessagesCount() (int, error) {
	var result int
	err := m.db.Raw("select count(*) as result from room").Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
