package sorm

import (
	"fmt"
	"reflect"
)

type ISelect interface {
	Where(where ...interface{}) *Sorm
	Values(values ...interface{}) *Sorm
}

func (objeto *Sorm) Where(where ...interface{}) *Sorm {
	numero := len(where)
	//var queryWhere = ""
	contador := 0
	/*var parametros []interface{}
	parametros = append(parametros, where...)*/
	if numero%2 == 0 {
		for _, campo := range where {
			/*objeto.Campos = nil
			objeto.Valores = nil*/
			if contador%2 == 0 {
				if reflect.ValueOf(campo).Kind() == reflect.String {
					objeto.Campos = append(objeto.Campos, reflect.ValueOf(campo).String())
				}
				/*nombre := reflect.TypeOf(campo).Field(contador).Name
				nombre2 := reflect.TypeOf(campo)
				fmt.Printf("%s, %v ", nombre, nombre)
				fmt.Printf("%s, %v ", nombre2, nombre2)
				val := reflect.ValueOf(campo).Elem()
				valueField := val.Field(0).Interface()
				typeField := val.Type().Field(0).Name
				println(valueField)
				println(typeField)
				elementos := reflect.ValueOf(parametros[contador]).Elem()

				nombreAtributo2 := elementos.Type().Field(0).Name
				println(nombreAtributo2)*/

				//objeto.Campos = append(objeto.Campos, reflect.TypeOf(campo).Name())
			} else {
				//reflect.ValueOf(b).String()
				switch reflect.ValueOf(campo).Kind() {
				case reflect.Int:
					valor := fmt.Sprintf("%v", reflect.ValueOf(campo))
					objeto.Valores = append(objeto.Valores, valor)
				case reflect.String:
					valor := fmt.Sprintf("\"%v\"", reflect.ValueOf(campo))
					objeto.Valores = append(objeto.Valores, valor)
				case reflect.Float64:
					valor := fmt.Sprintf("%v", reflect.ValueOf(campo))
					objeto.Valores = append(objeto.Valores, valor)
				case reflect.Float32:
					valor := fmt.Sprintf("%v", reflect.ValueOf(campo))
					objeto.Valores = append(objeto.Valores, valor)
				default:

				}
				//fmt.Printf("%s, %v ", reflect.ValueOf(campo).Interface(), reflect.ValueOf(campo).Interface())
				//fmt.Printf("%s, %v ", reflect.ValueOf(campo))
				//objeto.Valores = append(objeto.Valores, reflect.ValueOf(campo).Interface())
			}
			contador++
		}
		println(numero)
	} else {
		println(numero)
	}
	return objeto
}
func (objeto *Sorm) Values(values ...interface{}) *Sorm {
	return objeto
}
func (objeto *Sorm) Select(from ...string) *Sorm {
	objeto.queryIn = queryIn{}
	total := len(from)
	if total == 0 {
		objeto.From = " * "
		return objeto
	}
	for i := 0; i < total; i++ {
		if i == 0 {
			objeto.From = from[i]
		} else {
			objeto.From += ", " + from[i]
		}
	}
	return objeto
}
