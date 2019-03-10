package enrutador

import (
	"net/http"

	controller "../controladores"
	middle "../middlewares"
)

/*

 */
func RauterUsuario(router *http.ServeMux) {
	lt := middle.ChainMiddleware(middle.WithLogging, middle.WithTracing)
	router.HandleFunc("/usuarios", controller.GetUsuarios)
	router.HandleFunc("/usuario", controller.CreateUsuario)
	router.HandleFunc("/getusuario", controller.GetUsuario)
	router.HandleFunc("/misdeudas", controller.GetDeudas)
	router.HandleFunc("/login", lt(controller.GetUsuario))
}
