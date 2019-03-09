package dto

import (
	"fmt"
	"testing"

	entidades "../entidades"
)

func TestInsertAll(t *testing.T) {
	a := &entidades.Usuario{
		CodUsuario:    5,
		NombreUsuario: "ljfkpqe",
		Correo:        "aeij 3rij1o3",
		Password:      "3sflgmwop",
		Estado:        2,
	}
	val, err := InsertAll(a)
	if err == nil {
		fmt.Println(*val)
		t.Log(val)
	}
}

func TestGetAll(t *testing.T) {
	a := &entidades.Usuario{
		CodUsuario:    5,
		NombreUsuario: "ljfkpqe",
		Correo:        "aeij 3rij1o3",
		Password:      "3sflgmwop",
		Estado:        2,
	}
	val, err := GetAll(a)
	if err == nil {
		fmt.Println(*val)
		t.Log(val)
	}
}
func TestGetById(t *testing.T) {
	a := &entidades.Usuario{
		CodUsuario:    5,
		NombreUsuario: "ljfkpqe",
		Correo:        "aeij 3rij1o3",
		Password:      "3sflgmwop",
		Estado:        2,
	}
	val, err := GetById(a, 5)
	if err == nil {
		fmt.Println(*val)
		t.Log(val)
	}
}
func TestUpdateById(t *testing.T) {
	a := &entidades.Usuario{
		CodUsuario:    5,
		NombreUsuario: "ljfkpqe",
		Correo:        "aeij 3rij1o3",
		Password:      "3sflgmwop",
		Estado:        2,
	}
	val, err := UpdateById(a, 5)
	if err == nil {
		fmt.Println(*val)
		t.Log(val)
	}
}
