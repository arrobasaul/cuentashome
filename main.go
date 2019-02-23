package main

import (
	"encoding/json"
	"fmt"

	ent "./entidades"
)

func main() {
	fmt.Println("Contabilidad de la casa")
	m := ent.Cuenta{
		CodCuenta:    1,
		NombreCuenta: "",
		Valor:        1.2,
		Estado:       1,
	}
	b, _ := json.Marshal(m)
	println(b)
}
