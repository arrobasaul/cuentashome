package dto

import (
	conn "../conn"
	dto "../entidades"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsuarios() (*[]dto.Usuario, error) {

	// run your query, fill in &u...
	db := conn.Conexion()
	result, err := db.Query("SELECT CodUsuario, NombreUsuario FROM usuarios")
	if err != nil {
		panic(err.Error())
	}
	usuario := dto.Usuario{}
	usuarios := []dto.Usuario{}
	for result.Next() {
		err = result.Scan(&usuario.CodUsuario, &usuario.NombreUsuario)
		if err != nil {
			panic(err.Error())
		}
		usuarios = append(usuarios, usuario)
	}
	defer db.Close()

	return &usuarios, nil
}
