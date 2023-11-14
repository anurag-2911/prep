package requestratelimit

import 
(
	"time"
)

func HandleRateLimit(url string) {
	time.Sleep(10*time.Millisecond) //todo: added a delay for now, replace it with rate limit logic
}