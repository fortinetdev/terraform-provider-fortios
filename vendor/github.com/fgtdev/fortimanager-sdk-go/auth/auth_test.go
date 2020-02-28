package auth

import (
	"testing"
)

func subRoutine(index int, authBlob *AuthBlob, ch chan int) {
	authBlob.EnterAuthCriticalSection()
	if authBlob.LeaveAuthCriticalSection() != nil {
		panic("LeaveAuthCriticalSection error")
	}
	ch <- index
}

func TestLock(t *testing.T) {
	authBlob := FmgAuthInit(nil, nil)
	ch := make(chan int)
	for idx := 0; idx < 100; idx++ {
		go subRoutine(idx, authBlob, ch)
	}

	counter := 0
	channelCounter := 0
	for idx := 0; idx < 100; idx++ {
		counter += idx
		channelCounter += <-ch
	}
	if counter != channelCounter {
		t.Errorf("TestLock fails")
	}
}
