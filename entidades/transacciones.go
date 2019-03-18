package entidades

import (
	"fmt"

	. "../sorm"
)

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Transacciones struct {
	Sorm             `json:"-" llave:"si" schema:"si"`
	CodTransacciones int     `json:"CodTransacciones" llave:"si" schema:"si"`
	CodCuentas       int     `json:"CodCuenta" schema:"si" FK:"si"`
	CodPagos         int     `json:"CodPagos" schema:"si" FK:"si"`
	Valor            float32 `json:"Valor" schema:"si"`
	FechaCuenta      string  `json:"FechaDeuda" schema:"si" isFecha:"si"`
	Estado           int     `json:"Estado" schema:"si"`
	//Pagos            []Pagos `json:"usuario" Map:"NO"`
}

func Init2() {
	trans := &Transacciones{}
	trans.B = trans

	sormvar := trans.Select().Where("CodTransacciones", 1.5, "FechaCuenta", "fecha")
	println(sormvar.Tabla)

	println()
	println("Select")
	println("Campos")
	for _, valo := range sormvar.Campos {
		println(valo)
	}
	println("Valores")
	for _, valo := range sormvar.Valores {
		fmt.Printf("%v", valo)
	}

	sormvar2 := trans.Insert("Valor", 2.6, "Estado", 0)
	println()
	println("Insert")
	println("Campos")
	for _, valo := range sormvar2.Campos {
		println(valo)
	}
	println("Valores")
	for _, valo := range sormvar2.Valores {
		fmt.Printf("%v", valo)
	}

}
