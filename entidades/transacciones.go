package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Transacciones struct {
	CodTransacciones int     `json:"CodTransacciones" llave:"si" schema:"si"`
	CodCuentas       int     `json:"CodCuenta" schema:"si" FK:"si"`
	CodPagos         int     `json:"CodPagos" schema:"si" FK:"si"`
	Valor            float32 `json:"Valor" schema:"si"`
	FechaCuenta      string  `json:"FechaDeuda" schema:"si" isFecha:"si"`
	Estado           int     `json:"Estado" schema:"si"`
	//Pagos            []Pagos `json:"usuario" Map:"NO"`
}
