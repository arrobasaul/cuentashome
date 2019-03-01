package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query().Get("nombre")
		//i, err := strconv.Atoi(keys)
		println(keys)
		if keys == "saul" {
			next.ServeHTTP(w, r)
		} else {
			var mensaje string
			mensaje = "no autorizado"
			json.NewEncoder(w).Encode(&mensaje)
		}

	}
}
func WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
