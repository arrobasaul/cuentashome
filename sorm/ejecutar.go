package sorm

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (o *Sorm) Ejecutar() {
	if 'S' == 'S' {
		o.ejecutar()
	}
}

// Ejecutar ejecuta la sentencia sql.
func (o *Sorm) ejecutar() (int, []map[string]interface{}, error) {
	// verificar que la sentencia pueda ser generada.
	if err := o.validarSql(); err != nil {
		return 0, nil, err
	}

	// obtener sentencia 'select...'.
	sentencia := o.Tabla
	/*if err != nil {
		return 0, nil, err
	}*/

	// slice de parámetros de la sentencia sql a ejecutar.
	var parametros []interface{}

	// where.
	if o.where != "" {
		parametros = append(parametros, o.Valores...)
	}

	// having.
	/*
		if o.teniendoCondicion != "" {
			parametros = append(parametros, o.teniendoValores...)
		}*/

	var filas *sql.Rows
	if o.db != nil {
		filas, _ = o.db.Query(sentencia, parametros...)
	} else {
		filas, _ = o.tx.Query(sentencia, parametros...)
	}
	/*if err != nil {
		return 0, nil, envolverError(err, errorBD, mensajeErrorSeleccionar+mensajeErrorSeleccionarEjecutar)
	}*/
	defer filas.Close()

	var cant int
	var datos []map[string]interface{}
	if o.B != nil {
		cant, _ = asignarAObjeto(filas, o.B)
	} else {
		cant, datos, _ = asignarAMapa(filas)
	}
	/*if err != nil {
		return 0, nil, envolverError(err, errorDeUso, mensajeErrorSeleccionar+mensajeErrorSeleccionarRetorno)
	}*/

	return cant, datos, nil
}

// validarSql valida que la sentencia sql pueda ser generada.
func (o *Sorm) validarSql() error {
	var msjError string
	if o.Tabla == "" {
		msjError = "" //mensajeErrorSeleccionar + mensajeErrorTablaVacia
	}
	if len(o.Campos) == 0 {
		if msjError != "" {
			msjError += ", "
		} else {
			msjError = "" //mensajeErrorSeleccionar
		}
		msjError += "" //mensajeErrorCamposVacios
	}

	/*if msjError != "" {
		return nuevoError(errorDeUso, msjError)
	}*/

	return nil
}

// asignarAMapa asigna las filas (*sql.Rows) a un slice de mapas de campos.
// Devuelve la cantidad de elementos recorridos, los datos y el error en
// caso de suceder.
func asignarAMapa(filas *sql.Rows) (int, []map[string]interface{}, error) {
	var datos []map[string]interface{}

	// Nombres de campos.
	campos, err := filas.Columns()
	if err != nil {
		return 0, nil, err
	}

	// Obtener los valores de los campos de la base de datos.
	var valores = make([]interface{}, len(campos))
	for i := range valores {
		var ii interface{}
		valores[i] = &ii
	}

	var cant int
	for filas.Next() {
		cant++

		if err := filas.Scan(valores...); err != nil {
			return 0, nil, nil
		}

		// Recibir los valores de los campos de la fila en el mapa.
		mapaFila := make(map[string]interface{})
		for i, campo := range campos {
			var valorCrudo = *(valores[i].(*interface{}))
			var tipoCrudo = reflect.TypeOf(valorCrudo)

			switch {
			case tipoCrudo == nil:
				// Si el campo es nulleable y el valor es null...
				mapaFila[campo] = nil
			case tipoCrudo.String() == "[]uint8":
				if s, ok := valorCrudo.([]uint8); ok {
					mapaFila[campo] = fmt.Sprintf("%s", s)
				}
			default:
				mapaFila[campo] = valorCrudo
			}
		}

		// Agregar los datos de la fila en el slice de mapas de datos de retorno.
		datos = append(datos, mapaFila)
	}

	return cant, datos, nil
}

// asignarAObjeto asigna las filas (*sql.Rows) al objeto recibido.
// Devuelve la cantidad de elementos recorridos y el error en caso
// de suceder.
func asignarAObjeto(filas *sql.Rows, objeto interface{}) (int, error) {
	// Nombres de campos del resultado obtenido de la base de datos.
	camposFila, err := filas.Columns()
	if err != nil {
		return 0, err
	}

	// Sabiendo que el objeto es un puntero de slice de una estructura se
	// asigna un elemento para ir incorporando las estructuras.
	sliceDeObjetos := reflect.ValueOf(objeto).Elem()

	// Obtener la lista de campos de la estructura que podrían ser asignados
	// desde la consulta obtenida.
	estructura := reflect.TypeOf(objeto).Elem().Elem()
	camposEstructura := make(map[string]string)
	var cantCampos int

	for i := 0; i < estructura.NumField(); i++ {
		campo := estructura.Field(i)
		campoTabla := campo.Tag.Get("bdsql")

		switch campoTabla {
		case "-":
			// No asignar el valor.
		case "":
			// En caso que el campo de la estructura no tenga asignado el
			// tag "bdsql", se asume que el nombre del campo de la la consulta
			// sql, es el mismo (en minúsculas) que el nombre de campo de la
			// estructura.
			cantCampos++
			camposEstructura[campo.Name] = strings.ToLower(campo.Name)
		default:
			// En caso que el campo de la estructura tenga asignado el
			// tag "bdsql", se toma el tag para obtener el dato de la
			// consulta sql.
			cantCampos++
			camposEstructura[campo.Name] = campoTabla
		}
	}
	// Podría pasar (raramente) que todos los campos de la estructura contengan
	// el tag `bdsql:"-"`. En ese caso, la estrucutra no permite que ninguno de
	// sus campos sean asignables por los valores recibidos de la consulta sql.
	if cantCampos == 0 {
		return 0, err //nuevoError(errorDeUso, mensajeErrorSeleccionar+mensajeErrorSeleccionarReflexionSinAsignacion)
	}

	// Mapa de funciones:
	// Se crea un slice el cuál mantiene la función que completa cada campo
	// de la estructura.
	funciones := make(map[string]func(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error), estructura.NumField())

	// Verificar qe todos los campos de la consulta obtenida, puedan ser
	// ingresados en el objeto (slice de la estructura) recibido.
	var camposFaltantes string
	estructuraNueva := reflect.New(estructura).Elem()
	for _, campoFila := range camposFila {
		var encontrado bool
		for ce, cf := range camposEstructura {
			if campoFila == cf {
				encontrado = true

				// Llenar el slice de funciones según cada tipo de campo
				// de la estructura.
				switch estructuraNueva.FieldByName(ce).Kind() {
				case reflect.Int:
					funciones[ce] = valori
				case reflect.Int8:
					funciones[ce] = valori8
				case reflect.Int16:
					funciones[ce] = valori16
				case reflect.Int32:
					funciones[ce] = valori32
				case reflect.Int64:
					funciones[ce] = valori64
				case reflect.Uint:
					funciones[ce] = valorui
				case reflect.Uint8:
					funciones[ce] = valorui8
				case reflect.Uint16:
					funciones[ce] = valorui16
				case reflect.Uint32:
					funciones[ce] = valorui32
				case reflect.Uint64:
					funciones[ce] = valorui64
				case reflect.Float32:
					funciones[ce] = valorf32
				case reflect.Float64:
					funciones[ce] = valorf64
				case reflect.Bool:
					funciones[ce] = valorbool
				case reflect.String:
					funciones[ce] = valorstring
				case reflect.Struct:
					if estructuraNueva.FieldByName(ce).Type().String() == "time.Time" {
						funciones[ce] = valorFecha
					} else {
						return 0, err //nuevoError(errorDeUso, mensajeErrorSeleccionar+mensajeErrorSeleccionarReflexionFecha)
					}
				default:
					return 0, err //nuevoError(errorDeUso, mensajeErrorSeleccionar+mensajeErrorSeleccionarReflexionCampo)
				}

				break
			}
		}
		if !encontrado {
			if camposFaltantes != "" {
				camposFaltantes += ", "
			}
			camposFaltantes += campoFila
			continue
		}
	}
	if camposFaltantes != "" {
		return 0, err //nuevoError(errorDeUso, fmt.Sprintf(mensajeErrorSeleccionar+"los campos obtenidos de la consulta, no existen en su totalidad en la estructura. Faltantes: %v", camposFaltantes))
	}

	// Obtener los valores de los campos de la base de datos.
	var valores = make([]interface{}, len(camposFila))
	for i := range valores {
		var ii interface{}
		valores[i] = &ii
	}

	// Recorrer todas las filas e insertarlas en el puntero de slice de objetos.
	var cant int
	for filas.Next() {
		cant++

		if err := filas.Scan(valores...); err != nil {
			return 0, err //envolverError(err, errorBD, mensajeErrorSeleccionar+"no es posible obtener los datos (scan)")
		}

		for i := 0; i < len(camposFila); i++ {
			var valorCrudo = *(valores[i].(*interface{}))
			var tipoCrudo = reflect.TypeOf(valorCrudo)

			for campoEstructura, campoFila := range camposEstructura {
				if camposFila[i] == campoFila {
					x, err := funciones[campoEstructura](valorCrudo, tipoCrudo)
					if err != nil {
						return 0, err // nuevoError(errorDeUso, fmt.Sprintf("%v, campo: %v", err, campoFila))
					}
					estructuraNueva.FieldByName(campoEstructura).Set(x)
				}
			}
		}

		sliceDeObjetos.Set(reflect.Append(sliceDeObjetos, estructuraNueva))
	}

	return cant, nil
}

// ---- Funciones de asignación de campos de la estructura ---------------------

func valori(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v int

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.Atoi(string(s))
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = vTemp
		} else {
			vTemp, ok := valorCrudo.(int64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = int(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valori8(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v int8

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseInt(string(s), 10, 8)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = int8(vTemp)
		} else {
			vTemp, ok := valorCrudo.(int64)
			if !ok {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = int8(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valori16(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v int16

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseInt(string(s), 10, 16)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = int16(vTemp)
		} else {
			vTemp, ok := valorCrudo.(int64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = int16(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valori32(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v int32

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseInt(string(s), 10, 32)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = int32(vTemp)
		} else {
			vTemp, ok := valorCrudo.(int64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = int32(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valori64(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v int64

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseInt(string(s), 10, 64)
			if err != nil {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = vTemp
		} else {
			vTemp, ok := valorCrudo.(int64)
			if !ok {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorui(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v uint

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseUint(string(s), 10, 64)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = uint(vTemp)
		} else {
			vTemp, ok := valorCrudo.(uint)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = uint(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valorui8(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v uint8

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseUint(string(s), 10, 8)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = uint8(vTemp)
		} else {
			vTemp, ok := valorCrudo.(uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorui16(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v uint16

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseUint(string(s), 10, 16)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = uint16(vTemp)
		} else {
			vTemp, ok := valorCrudo.(uint16)
			if !ok {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorui32(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v uint32

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseUint(string(s), 10, 32)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = uint32(vTemp)
		} else {
			vTemp, ok := valorCrudo.(uint32)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorui64(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v uint64

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil // nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseUint(string(s), 10, 64)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = vTemp
		} else {
			vTemp, ok := valorCrudo.(uint64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorf32(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v float32

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseFloat(string(s), 32)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = float32(vTemp)
		} else {
			vTemp, ok := valorCrudo.(float64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = float32(vTemp)
		}
	}

	return reflect.ValueOf(v), nil
}

func valorf64(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v float64

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseFloat(string(s), 64)
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = vTemp
		} else {
			vTemp, ok := valorCrudo.(float64)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
			v = vTemp
		}
	}

	return reflect.ValueOf(v), nil
}

func valorbool(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v bool

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			vTemp, err := strconv.ParseBool(string(s))
			if err != nil {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionAsignacion)
			}
			v = vTemp
		} else {
			switch valorCrudo.(type) {
			case bool:
				v = valorCrudo.(bool)
			case int64:
				v = valorCrudo == int64(1)
			default:
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
			}
		}
	}

	return reflect.ValueOf(v), nil
}

func valorstring(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v string

	if tipoCrudo != nil {
		if tipoCrudo.String() == "[]uint8" {
			s, ok := valorCrudo.([]uint8)
			if !ok {
				return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionUint8)
			}
			v = fmt.Sprintf("%s", s)
		} else {
			v = fmt.Sprintf("%s", valorCrudo)
		}
	}

	return reflect.ValueOf(v), nil
}

func valorFecha(valorCrudo interface{}, tipoCrudo reflect.Type) (reflect.Value, error) {
	var vacio reflect.Value
	var v time.Time = time.Time{}

	if valorCrudo != nil {
		fh, ok := valorCrudo.(time.Time)
		if !ok {
			return vacio, nil //nuevoError(errorDeUso, mensajeErrorSeleccionarReflexionTipo)
		}
		v = fh
	}

	return reflect.ValueOf(v), nil
}
