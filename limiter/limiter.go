package limiter

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	Limit       int
	Burst       int
	BurstRepeat int
	BurstChan   chan time.Time

	burstRepeatChan <-chan time.Time
	limitChan       <-chan time.Time

	wg sync.WaitGroup
}

// InitSetup sets up the limiter channels providding the limiting
// functionality
func (l *Limiter) initSetup() {
	l.BurstChan = make(chan time.Time, l.Burst)
	l.fillInBurst()

	l.burstRepeatChan = time.Tick(time.Duration(l.BurstRepeat) * time.Second)
	l.limitChan = time.Tick(time.Duration(l.Limit) * time.Millisecond)
}

// fillInBurst fills the the burst channel
func (l *Limiter) fillInBurst() {
	//fmt.Printf("\nbefore: length of BurstChan = %d ... %s\n", len(l.BurstChan), time.Now())
	for i := 0; i < l.Burst; i++ {
		l.BurstChan <- time.Now()
	}
	//fmt.Printf("after: length of BurstChan = %d ... %s\n\n", len(l.BurstChan), time.Now())
}

// Run starts the limiting by sending into the burst channel.
// Limit channel and repeat burst channels are ticker channels
// that cause the limiting and the bursts.
func (l *Limiter) Run() {

	l.initSetup()

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for t := range l.limitChan {
			select {
			case <-l.burstRepeatChan:
				fmt.Printf("\n\tBURST %s\n\n", t)
				l.fillInBurst()
			default:
				l.BurstChan <- t
				//fmt.Printf("\tLIMITED %s\n", t)
			}
		}
	}()
}
