package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")
	assert.Nil(t, err)
}
