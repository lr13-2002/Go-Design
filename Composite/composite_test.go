package composite

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	assert.Equal(t, 20, got)
}
