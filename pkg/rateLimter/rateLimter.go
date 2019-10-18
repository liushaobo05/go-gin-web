package rateLimter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	Interval time.Duration
	MaxCount int
	Lock     sync.Mutex
	ReqCount int
}

func NewRateLimiter(interval time.Duration, maxCnt int) *RateLimiter {
	rateLimit := &RateLimiter{
		Interval: interval,
		MaxCount: maxCnt,
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			rateLimit.Lock.Lock()
			rateLimit.ReqCount = 0
			rateLimit.Lock.Unlock()
		}
	}()

	return rateLimit
}

func (rateLimit *RateLimiter) Incr() {
	rateLimit.Lock.Lock()
	defer rateLimit.Lock.Unlock()

	rateLimit.ReqCount += 1
}

func (rateLimit *RateLimiter) IsLimit() bool {
	rateLimit.Lock.Lock()
	defer rateLimit.Lock.Unlock()

	return rateLimit.ReqCount < rateLimit.MaxCount
}
