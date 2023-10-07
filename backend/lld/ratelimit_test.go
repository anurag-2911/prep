package lld

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRL(t *testing.T){
	fmt.Println("test")
	rateLimiter := NewRateLimiter(5, 1)

	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		if rateLimiter.Allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Not Allowed\n", i+1)
		}
	}
}

type RateLimiter struct {
	tokens int
	maxTokens int
	tokensPerSecond int
	lastRefillTime time.Time
	lock sync.Mutex
}

func NewRateLimiter(maxTokens int, tokensPerSecond int) *RateLimiter{
	return &RateLimiter{
		tokens: maxTokens,
		maxTokens: maxTokens,
		tokensPerSecond: tokensPerSecond,
		lastRefillTime: time.Now(),
	}
}

func (rl *RateLimiter)refillTokens(){
	now := time.Now()
	elapsed := now.Sub(rl.lastRefillTime)
	newTokens := int(elapsed.Seconds()) * rl.tokensPerSecond
	if newTokens > 0 {
		if rl.tokens+newTokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		} else {
			rl.tokens += newTokens
		}
		rl.lastRefillTime = now
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	rl.refillTokens()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}