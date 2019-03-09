package servicios

import (
	"fmt"

	conn "../conn"
	entidades "../entidades"
	_ "github.com/go-sql-driver/mysql"
)

func GetDeudas(id int) (*[]entidades.Cuentas, error) {

	// run your query, fill in &u...
	db := conn.Conexion()
	query := fmt.Sprintf("SELECT * FROM homedb.Cuentas c where c.CodUsuarios = %d and c.Estado = 1;", id)
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	cuenta := entidades.Cuentas{}
	cuentas := []entidades.Cuentas{}
	for result.Next() {
		err = result.Scan(&cuenta.CodCuentas, &cuenta.CodDeudas, &cuenta.CodUsuarios, &cuenta.ValorInicial, &cuenta.ValorActual, &cuenta.FechaCuenta, &cuenta.Estado)
		if err != nil {
			panic(err.Error())
		}
		query2 := fmt.Sprintf("SELECT * FROM homedb.Transacciones t where t.CodCuentas = %d ;", cuenta.CodCuentas)
		result2, err := db.Query(query2)
		transaccion := entidades.Transacciones{}
		transacciones := []entidades.Transacciones{}
		if err != nil {
			panic(err.Error())
		}
		for result2.Next() {
			err = result2.Scan(&transaccion.CodTransacciones, &transaccion.CodCuentas, &transaccion.CodPagos, &transaccion.Valor, &transaccion.FechaCuenta)
			if err != nil {
				panic(err.Error())
			}
			transacciones = append(transacciones, transaccion)
		}

		cuenta.Transacciones = transacciones
		cuentas = append(cuentas, cuenta)
	}
	defer db.Close()

	return &cuentas, nil
}
