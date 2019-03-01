package main

import (
	"log"
	"net/http"
	"time"

	enrutador "./enrutador"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := enrutador.GetRouter()
	enrutador.RauterUsuario(r)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
