package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tamarakaufler/go-limiter/limiter"
)

var (
	requests    int
	limit       int
	burst       int
	burstrepeat int
)

func init() {
	flag.IntVar(&requests, "requests", 50, "Number of requests")
	flag.IntVar(&limit, "limit", 300, "Number of limiting miliseconds")
	flag.IntVar(&burst, "burst", 3, "Value of the burst (how many requests can go through without limiting")
	flag.IntVar(&burstrepeat, "burstrepeat", 3, "Number of seconds, after which the burst will be repeated")
}

func main() {
	flag.Parse()

	rc := requests
	requests := make(chan int, rc)
	for i := 1; i <= rc; i++ {
		requests <- i
	}
	close(requests)

	l := &limiter.Limiter{
		Limit:       limit,
		Burst:       burst,
		BurstRepeat: burstrepeat,
	}
	l.Run()

	for req := range requests {
		<-l.BurstChan
		fmt.Printf("==> R E Q U E S T [%d] ... %s\n", req, time.Now())
		time.Sleep(10 * time.Millisecond)
	}

}
