package main

import (
    "net/http"
)

// application/json.
func ContentTypeJson(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}
