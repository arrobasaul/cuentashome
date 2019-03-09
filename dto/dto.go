package dto

import (
	"fmt"
	"reflect"
	"strings"
)

type todo interface{}

func GetAll(a interface{}) (*string, error) {

	elementos := reflect.ValueOf(a).Elem()
	nombre := strings.ToLower(reflect.TypeOf(a).Elem().Name())

	var query string
	query = "select "
	for j := 0; j < elementos.NumField(); j++ {
		nombreAtributo := strings.ToLower(elementos.Type().Field(j).Name)
		query += fmt.Sprintf(", %s ", nombreAtributo)
	}

	query = strings.Replace(query, "select ,", "select ", -1)
	query += fmt.Sprintf(" from %s ", nombre)
	return &query, nil
}

func GetById(a interface{}, id int) (*string, error) {

	elementos := reflect.ValueOf(a).Elem()
	nombre := strings.ToLower(reflect.TypeOf(a).Elem().Name())
	var Millave string
	var query string
	query = "select "
	for j := 0; j < elementos.NumField(); j++ {

		nombreAtributo := strings.ToLower(elementos.Type().Field(j).Name)

		llave := elementos.Type().Field(j).Tag.Get("llave")
		if llave == "SI" {
			Millave = nombreAtributo
		}
		query += fmt.Sprintf(", %s ", nombreAtributo)
	}
	query = strings.Replace(query, "select ,", "select ", -1)
	query += fmt.Sprintf(" from %s where %s = %d", nombre, Millave, id)

	return &query, nil
}
func InsertAll(a interface{}) (*string, error) {

	elementos := reflect.ValueOf(a).Elem()
	nombre := strings.ToLower(reflect.TypeOf(a).Elem().Name())

	var query string
	query = "insert into " + nombre + " set "
	for j := 0; j < elementos.NumField(); j++ {
		//tag := typeField.Tag
		//tag.Get("tag_name")

		nombreAtributo := strings.ToLower(elementos.Type().Field(j).Name)

		switch elementos.Field(j).Kind() {
		case reflect.Int:
			valor := elementos.Field(j).Int()
			query += fmt.Sprintf(", %s = %v", nombreAtributo, valor)
		case reflect.String:
			valor := elementos.Field(j).String()
			query += fmt.Sprintf(", %s = \"%v\"", nombreAtributo, valor)
		default:

		}
	}
	query = strings.Replace(query, "set ,", " set ", -1)
	println(query)
	return &query, nil
}
func UpdateById(a interface{}, id int) (*string, error) {

	elementos := reflect.ValueOf(a).Elem()
	nombre := strings.ToLower(reflect.TypeOf(a).Elem().Name())
	var Millave string
	var query string
	query = "update " + nombre + " set "
	for j := 0; j < elementos.NumField(); j++ {
		//tag := typeField.Tag
		//tag.Get("tag_name")

		nombreAtributo := strings.ToLower(elementos.Type().Field(j).Name)
		llave := elementos.Type().Field(j).Tag.Get("llave")
		if llave == "SI" {
			Millave = nombreAtributo
			continue
		}
		switch elementos.Field(j).Kind() {
		case reflect.Int:
			valor := elementos.Field(j).Int()
			if valor != 0 {
				query += fmt.Sprintf(", %s = %v", nombreAtributo, valor)
			}
		case reflect.String:
			valor := elementos.Field(j).String()
			if valor != "" {
				query += fmt.Sprintf(", %s = \"%v\"", nombreAtributo, valor)
			}
		default:

		}
	}
	query = strings.Replace(query, "set ,", " set ", -1)
	query += fmt.Sprintf(" where %s = %v", Millave, id)
	return &query, nil
}
