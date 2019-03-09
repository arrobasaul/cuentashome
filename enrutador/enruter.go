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
	router.HandleFunc("/usuarios", controller.GetUsuarios)
	router.HandleFunc("/usuario", controller.CreateUsuario)
	router.HandleFunc("/usuario", controller.GetUsuario)

	router.HandleFunc("/login", lt(controller.GetUsuario))
}
