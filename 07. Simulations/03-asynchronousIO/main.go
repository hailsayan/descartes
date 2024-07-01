package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
	// TODO
}

type Task func() string

func Async(t Task) *FutureResult {
	result := &FutureResult{
		ResultChan: make(chan string),
	}

	go func() {
		funcResult := t()
		result.Done.Store(true)
		result.ResultChan <- funcResult
	}()

	return result
}

func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	result := &FutureResult{
		ResultChan: make(chan string),
	}

	go func() {
		funcResult := t()
		result.Done.Store(true)
		result.ResultChan <- funcResult
	}()

	go func() {

		select {

		case <-time.After(timeout):
			result.ResultChan <- "timeout"
		}
	}()

	return result
}

func (fResult *FutureResult) Await() string {
	return <-fResult.ResultChan
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	all := &FutureResult{
		ResultChan: make(chan string, len(fResults)),
	}

	go func() {
		for _, result := range fResults {
			response := result.Await()
			all.ResultChan <- response
		}
		all.Done.Store(true)

	}()
	return all
}
