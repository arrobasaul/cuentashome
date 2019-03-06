package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Cuentas struct {
	CodCuenta    int     `json:"CodDeuda"`
	CodDeuda     int     `json:"CodDeuda"`
	ValorInicial float32 `json:"Valor"`
	ValorActual  float32 `json:"Estado"`
	FechaCuenta  string  `json:"FechaDeuda"`

	Pagos []Pagos `json:"usuario" Map:"NO"`
}
