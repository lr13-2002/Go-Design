package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}
type IObserver interface {
	Update(msg string)
}

type SubJect struct {
	opservers []IObserver
}

func (sub *SubJect) Register(observer IObserver) {
	sub.opservers = append(sub.opservers, observer)
}

func (sub *SubJect) Remove(observer IObserver) {
	for i, ob := range sub.opservers {
		if ob == observer {
			sub.opservers = append(sub.opservers[:i], sub.opservers[i+1:]...)
			return
		}
	}
}

func (sub *SubJect) Notify(msg string) {
	for _, o := range sub.opservers {
		o.Update(msg)
	}
}

type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s", msg)
}
