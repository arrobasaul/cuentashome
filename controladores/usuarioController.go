package controladores

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dto "../dto"
	entidades "../entidades"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, error := dto.GetUsuarios()
	if error != nil {
		fmt.Println("errores")
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("correcto")
		json.NewEncoder(w).Encode(&usuarios)
	}
	//fmt.Fprintf(w, "Hola, %s, ¡este es un servidor!", r.URL.Path)
}
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//id := params["CodUsuario"]
	keys := r.URL.Query().Get("CodUsuario")
	i, err := strconv.Atoi(keys)
	if err != nil {
		// handle error
		fmt.Println(err)
		decoder := json.NewDecoder(r.Body)
		var usuario entidades.Usuario
		err = decoder.Decode(&usuario)
		if err != nil {
			json.NewEncoder(w).Encode(&entidades.Errores{Error: err.Error(), Descripcion: "no se envian datos"})
		} else {
			usua, error2 := dto.GetUsuario(usuario.CodUsuario)
			if error2 != nil {
				json.NewEncoder(w).Encode(&error2)
			} else {
				json.NewEncoder(w).Encode(&usua)
			}

		}
	} else {
		usua, error2 := dto.GetUsuario(i)
		if error2 != nil {
			json.NewEncoder(w).Encode(&error2)
		} else {
			json.NewEncoder(w).Encode(&usua)
		}
	}

}
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var usuario entidades.Usuario
	err := decoder.Decode(&usuario)
	fmt.Println(usuario)
	if err != nil {
		panic(err)
	} else {
		id := dto.CreateUsuario(usuario)
		json.NewEncoder(w).Encode(&id)
		fmt.Println("correcto")
	}
}
