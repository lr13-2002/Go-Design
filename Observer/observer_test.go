package observer

import "testing"

func TestSubjectNotify(t *testing.T) {
	sub := &SubJect{}
	sub.Register(&Observer1{})
	sub.Register(&Observer2{})
	sub.Notify("hi")
}
