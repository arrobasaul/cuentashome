package enrutador

import (
	controller "../controladores"
	middle "../middlewares"
	"github.com/gorilla/mux"
)

/*

 */
func RauterUsuario(router *mux.Router) {
	lt := middle.ChainMiddleware(middle.WithLogging, middle.WithTracing)
	router.HandleFunc("/usuarios", controller.GetUsuarios).Methods("GET")
	router.HandleFunc("/usuario", controller.CreateUsuario).Methods("POST")
	router.HandleFunc("/usuario", controller.GetUsuario).Methods("GET")

	router.HandleFunc("/login", lt(controller.GetUsuario)).Methods("GET")
}
