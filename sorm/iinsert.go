package sorm

import (
	"fmt"
	"reflect"
)

type IInsert interface {
	//Set(set ...interface{}) *Sorm
	//Values(values ...interface{}) *Sorm
}

func (objeto *Sorm) Insert(set ...interface{}) *Sorm {
	objeto.queryIn = queryIn{}
	numero := len(set)
	contador := 0
	if numero%2 == 0 {
		for _, campo := range set {
			/*objeto.Campos = nil
			objeto.Valores = nil*/
			if contador%2 == 0 {
				if reflect.ValueOf(campo).Kind() == reflect.String {
					objeto.Campos = append(objeto.Campos, reflect.ValueOf(campo).String())
				}
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
			}
			contador++
		}
	} else {
		println(numero)
	}
	return objeto
}
