package middelware

import (
	"log"
	"net"
	"net/http"
	"sync"
)

type Limiter struct {
	ipCount map[string]int
	sync.Mutex
}

var limiter Limiter

func init() {
	limiter.ipCount = make(map[string]int)
}

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP address for the current user.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Get the # of times the visitor has visited in the last 60 seconds
		limiter.Lock()
		count, ok := limiter.ipCount[ip]
		if !ok {
			limiter.ipCount[ip] = 0
		}
		if count > 10 {
			limiter.Unlock()
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		limiter.Unlock()
		next.ServeHTTP(w, r)
	})
}
