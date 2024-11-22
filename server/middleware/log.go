package middleware

import (
  "log"
  "net/http"
)

func Log(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Method, r.URL)
    next.ServeHTTP(w, r)
  })
}
