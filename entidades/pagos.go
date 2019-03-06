package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Pagos struct {
	CodPagos        int     `json:"CodCuenta" llave:"SI"`
	DescripionPagos string  `json:"DescripionCuenta"`
	Valor           float32 `json:"Valor"`
	Estado          int     `json:"Estado"`
	CodUsuario      int     `json:"CodUsuario" FK:"SI"`
	Usuario         `json:"usuario"`
}
