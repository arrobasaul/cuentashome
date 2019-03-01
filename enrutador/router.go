package enrutador

import "github.com/gorilla/mux"

func GetRouter() (router *mux.Router) {
	return mux.NewRouter()
}
