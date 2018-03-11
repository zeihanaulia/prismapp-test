package send

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_send_message_cannot_be_empty(t *testing.T) {
	s := NewService()
	err := s.Send("")
	assert.EqualError(t, err, "message cannot be empty")
}
