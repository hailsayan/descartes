package main

// DO NOT USE ANY IMPORT

type Qutex struct {
	ch chan bool
}

func NewQutex() *Qutex {
	return &Qutex{ch: make(chan bool, 1)}
}

func (q *Qutex) Lock() {
	q.ch <- true
}

func (q *Qutex) Unlock() {
	select {
	case <-q.ch:
	default:
		panic("Error in unlocking")
	}
}
