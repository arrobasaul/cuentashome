package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Pagos struct {
	CodPagos        int     `json:"CodCuenta" llave:"SI" schema:"si"`
	DescripionPagos string  `json:"DescripionCuenta" schema:"si"`
	Valor           float32 `json:"Valor" schema:"si"`
	Estado          int     `json:"Estado" schema:"si"`
	CodUsuarios     int     `json:"CodUsuario" FK:"SI" schema:"si"`
	Usuarios        `json:"usuario"`
}
