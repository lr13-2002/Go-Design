package singleton_test

import (
	singleton "go_design/Singleton"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetHungryInstance(), singleton.GetHungryInstance())
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}
