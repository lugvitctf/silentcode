package main

import (
	"fmt"
	"net/http"
)

var API_KEY string = "keyatempctfapithisfor"

func newFlag(botUsername string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		apiKey := query.Get("api_key")
		if apiKey != API_KEY {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		flag := query.Get("flag")
		if flag == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Flag not found"))
			return
		}
		sessionId := NewSession(flag)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("https://t.me/%s?start=n_%s", botUsername, sessionId)))
	}
}
