package main

import (
	"fmt"

	ent "./entidades"
)

func main() {
	d := ent.Cuenta{
		CodCuenta: 1,
	}
	fmt.Println(d)
}
