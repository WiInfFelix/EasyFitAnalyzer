package logger

import (
	"log"
	"net/http"
	"time"
)

type HTTPReqInfo struct {
	method    string
	uri       string
	referer   string
	ipaddr    string
	code      int
	size      int64
	duration  time.Duration
	userAgent string
}

func LogRequestHandler(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		ri := &HTTPReqInfo{
			method:    r.Method,
			uri:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			ipaddr:    r.RemoteAddr,
			userAgent: r.Header.Get("User-Agent"),
		}

		h.ServeHTTP(w, r)

		log.Printf("%s: %s \t %s \t %s \t from %s %s", time.Now().Format(time.RFC822), ri.method, ri.uri, ri.referer, ri.ipaddr, ri.userAgent)

	}

	return http.HandlerFunc(fn)
}
