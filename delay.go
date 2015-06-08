package delay

import (
	"net/http"
	"time"
)

const DELAY_HEADER_KEY = "X-Add-Delay"

type Middleware struct{}

func (m Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	delayHeaderValue := request.Header.Get(DELAY_HEADER_KEY)

	if delayHeaderValue != "" {
		delayDuration, err := time.ParseDuration(delayHeaderValue)

		if err == nil {
			time.Sleep(delayDuration)
		}
	}

	next(response, request)
}
