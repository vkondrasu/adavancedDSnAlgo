package tokenBucket

import (
	"sync"
	"time"
)

type RateLimiter interface {
	HaveTokens(int) bool
}

type TokenBucket struct {
	bucketCapacity  int
	availableTokens int
	refill          int
	rate            time.Duration
	lastRefillTime  time.Time
	sync.Mutex
}

func NewTokenBucket(bucketCapacity int, refill int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		bucketCapacity:  bucketCapacity,
		availableTokens: bucketCapacity,
		refill:          refill,
		rate:            rate,
		lastRefillTime:  time.Now(),
	}

}

func (t *TokenBucket) HaveTokens(n int) bool {
	t.refillTokens()
	t.Lock()
	defer t.Unlock()
	if t.availableTokens < n {
		return false
	}

	t.availableTokens -= n
	return true
}

func (t *TokenBucket) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(t.lastRefillTime)
	if elapsed < t.rate {
		return
	}
	t.lastRefillTime = now
	if t.availableTokens+t.refill >= t.bucketCapacity {
		t.availableTokens = t.bucketCapacity
	} else {
		t.availableTokens += t.refill
	}
}

/*
func maint() {
	rateLimterMap := make(map[string]RateLimiter)

	rlVenkat := NewTokenBucket(10, 10, time.Second)
	rateLimterMap["venkat"] = rlVenkat
	rlWSWD := NewTokenBucket(5, 3, time.Minute)
	rateLimterMap["wswd"] = rlWSWD

	for i := 0; i < 20; i++ {
		fmt.Println("venkat", rateLimterMap["venkat"].HaveTokens(2))
		fmt.Println("wswd", rateLimterMap["wswd"].HaveTokens(2))
		time.Sleep(100 * time.Millisecond)
	}

}
*/

