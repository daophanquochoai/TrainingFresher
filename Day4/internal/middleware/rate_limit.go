package middleware

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"net"
	"net/http"
	"time"
)

var rateLimit int64 = 10

func RateLimit(redis *redis.Client) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// tao key rate limit
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}
			key := "rate:" + ip

			// dem request
			count, err := redis.Incr(r.Context(), key).Result()
			fmt.Printf("INFO : key : %s - count : %d\n", key, count)

			// check so request
			if err == nil {
				if count == 1 {
					redis.Expire(r.Context(), key, time.Minute)
				}

				if count > rateLimit {
					fmt.Println("Too many requests")
					http.Error(w, "Too many requests", 429)
					return
				}
			}

			// tiep bp loc
			next.ServeHTTP(w, r)
		})
	}
}
