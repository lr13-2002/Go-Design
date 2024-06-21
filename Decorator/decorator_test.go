package decorator

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestColorSquare(t *testing.T) {
	sq := Sqare{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}
