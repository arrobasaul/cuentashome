package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Cuentas struct {
	CodCuentas    int             `json:"CodCuentas" llave:"si" schema:"si"`
	CodDeudas     int             `json:"CodDeudas" schema:"si" FK:"si"`
	CodUsuarios   int             `json:"CodUsuarios" schema:"si" FK:"si"`
	ValorInicial  float32         `json:"Valor" schema:"si"`
	ValorActual   float32         `json:"ValorActual" schema:"si"`
	FechaCuenta   string          `json:"FechaDeuda" schema:"si" isFecha:"si"`
	Estado        int             `json:"Estado" schema:"si"`
	Transacciones []Transacciones `json:"Transacciones" Map:"NO"`
}
