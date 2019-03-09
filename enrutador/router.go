package enrutador

import (
	"net/http"
)

func GetRouter() (router *http.ServeMux) {
	mux := http.NewServeMux()
	return mux
	//return mux.NewRouter()
}
