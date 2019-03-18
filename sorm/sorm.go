package sorm

//type Todo interface{}
/*import (
	"fmt"

	entidades "../entidades"
	. "github.com/ahmetb/go-linq"
)
*/
type Generador interface {
	Select(from ...string) *ISelect
	Insert(from ...string) *IInsert
	/*Update(tabla interface{}) *[]interface{}
	Insert(tabla interface{}) *[]interface{}
	Ejecutar() []interface{}

	From(tabla interface{}) *[]interface{}
	Replace(tabla interface{}) *[]interface{}*/
	/*InsertarEn(tabla interface{}) *[]ObgetoGenerador
	ModificarEn(tabla interface{}) *[]ObgetoGenerador
	EliminarEn(tabla string) *eliminarEn
	SeleccionarDe(tabla string) *seleccionarDe
	SeleccionarSql(sentencia string, valores ...interface{}) *seleccionarSql
	TxIniciar() (transaccionador, error)*/
}
type Sorm struct {
	queryIn
	B interface{}
}
type queryIn struct {
	/*db *sql.DB
	tx *sql.Tx*/

	Tabla   string
	where   string
	Campos  []string
	Valores []interface{}
	From    string

	//para insert
	Set string

	/*condicionValores []interface{}
	ordenadoPor      []string

	agruparPor        []string
	teniendoCondicion string
	teniendoValores   []interface{}

	limite int
	salto  int

	objeto interface{}

	juntaInternaTabla       string
	juntaInternaCondicion   string
	juntaIzquierdaTabla     string
	juntaIzquierdaCondicion string
	juntaDerechaTabla       string
	juntaDerechaCondicion   string
	juntaExternaTabla       string
	juntaExternaCondicion   string

	ejecutorSeleccionar*/
}

/*
type ObgetoGenerador interface {
	Select(tabla interface{}) *[]interface{}
	From(tabla interface{}) *[]interface{}
	Replace(tabla interface{}) *[]interface{}
}

func Pruebaslinq() {
	cuentasvarias := llenar()
	/*author := From(cuentasvarias).SelectMany( // make a flat array of authors
	func(Cuentasmias interface{}) Query {
		return From(Cuentasmias.(entidades.Cuentas).Transacciones)
	}).GroupBy( // group by author
	func(trasn interface{}) interface{} {
		return trasn // author as key
	}, func(trasn interface{}) interface{} {
		return trasn // author as value
	}).OrderByDescending( // sort groups by its length
	func(group interface{}) interface{} {
		return len(group.(Group).Group)
	}).Select( // get authors out of groups
	func(group interface{}) interface{} {
		return group.(Group).Key
	}) //.First() //dv take the first author
*/ /*
	var results []string
	type result2 struct {
		cuenta int
		trans  int
	}
	resultadostodos := []result2{}

	From(cuentasvarias).
		SelectManyBy(
			func(cuentas interface{}) Query { return From(cuentas.(entidades.Cuentas).Transacciones) },
			func(transacciones, cuentas interface{}) interface{} {
				val := result2{
					cuenta: cuentas.(entidades.Cuentas).CodCuentas,
					trans:  transacciones.(entidades.Transacciones).CodTransacciones,
				}
				return val
				//return fmt.Sprintf("{ Cuenta: %d, Transaccion: %d }", cuentas.(entidades.Cuentas).CodCuentas, transacciones.(entidades.Transacciones).CodTransacciones)
			},
		).
		ToSlice(&resultadostodos)

	for _, result := range results {
		fmt.Println(result)
	}

	for _, result2 := range resultadostodos {
		fmt.Println("cuenta")
		fmt.Println(result2.cuenta)
		fmt.Println("trans")
		fmt.Println(result2.trans)
	}
	/*author.ForEach(func(fruit interface{}) {
		fmt.Print("codtransaccione: ")
		fmt.Println(fruit.(entidades.Transacciones).CodTransacciones)
	})*/
/*
	From(cuentasvarias).Where(func(c interface{}) bool {
		return c.(Cuentas).CodCuentas >= 1
	}).Select(func(c interface{}) interface{} {
		return c.(Car).owner
	}).ToSlice(&owners)*/
/*}*/
/*
func llenar() []entidades.Cuentas {
	cuenta1 := entidades.Cuentas{
		CodCuentas:  1,
		FechaCuenta: "fecha",
		Estado:      2,
	}
	cuenta2 := entidades.Cuentas{
		CodCuentas:  2,
		FechaCuenta: "fecha 2",
		Estado:      2,
	}
	cuenta3 := entidades.Cuentas{
		CodCuentas:  3,
		FechaCuenta: "fecha",
		Estado:      0,
	}
	cuentas := []entidades.Cuentas{}
	transaccion1 := entidades.Transacciones{
		CodTransacciones: 1,
		CodCuentas:       2,
		FechaCuenta:      "fecha 1",
	}
	transaccion2 := entidades.Transacciones{
		CodTransacciones: 2,
		CodCuentas:       2,
		FechaCuenta:      "fecha 2",
	}
	transaccion3 := entidades.Transacciones{
		CodTransacciones: 3,
		CodCuentas:       2,
		FechaCuenta:      "fecha 3",
	}
	transaccion4 := entidades.Transacciones{
		CodTransacciones: 4,
		CodCuentas:       2,
		FechaCuenta:      "fecha 4",
	}
	transacciones1 := []entidades.Transacciones{}
	transacciones1 = append(transacciones1, transaccion2)
	transacciones1 = append(transacciones1, transaccion4)

	cuenta1.Transacciones = transacciones1

	cuentas = append(cuentas, cuenta1)

	transacciones2 := []entidades.Transacciones{}
	transacciones2 = append(transacciones2, transaccion1)
	transacciones2 = append(transacciones2, transaccion3)

	cuenta2.Transacciones = transacciones2

	cuentas = append(cuentas, cuenta2)

	transacciones3 := []entidades.Transacciones{}
	transacciones3 = append(transacciones3, transaccion1)
	transacciones3 = append(transacciones3, transaccion2)
	transacciones3 = append(transacciones3, transaccion4)

	cuenta3.Transacciones = transacciones3

	cuentas = append(cuentas, cuenta3)
	return cuentas
}
*/
