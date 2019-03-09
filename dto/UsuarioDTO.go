package dto

import (
	conn "../conn"
	entidades "../entidades"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsuarios() (*[]entidades.Usuarios, error) {

	// run your query, fill in &u...
	db := conn.Conexion()
	result, err := db.Query("SELECT CodUsuario, NombreUsuario FROM usuarios")
	if err != nil {
		panic(err.Error())
	}
	usuario := entidades.Usuarios{}
	usuarios := []entidades.Usuarios{}
	for result.Next() {
		err = result.Scan(&usuario.CodUsuarios, &usuario.NombreUsuario)
		if err != nil {
			panic(err.Error())
		}
		usuarios = append(usuarios, usuario)
	}
	defer db.Close()

	return &usuarios, nil
}
func GetUsuario(id int) (*entidades.Usuarios, *entidades.Errores) {

	// run your query, fill in &u...
	db := conn.Conexion()
	usuario := entidades.Usuarios{}
	var error2 entidades.Errores
	println(id)
	err := db.QueryRow("SELECT CodUsuario, NombreUsuario, Correo, Password, Estado FROM usuarios where CodUsuario=?", id).Scan(&usuario.CodUsuarios, &usuario.NombreUsuario, &usuario.Correo, &usuario.Password, &usuario.Estado)
	if err != nil {
		error2 = entidades.Errores{Error: err.Error(), Descripcion: "no encontrado"}
		return nil, &error2
	}
	return &usuario, nil
}
func CreateUsuario(usuario entidades.Usuarios) (id int64) {
	db := conn.Conexion()

	query, _ := InsertAll(&usuario)

	insForm, err := db.Prepare(*query)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//res, err := insForm.Exec(usuario.NombreUsuario, usuario.Correo, usuario.Password, usuario.Estado)

	res, err := insForm.Exec()
	if err != nil {
		panic(err.Error())
	} else {
		var id int64
		id, err = res.LastInsertId()
		if err != nil {
			panic(err.Error())
		} else {
			return id
		}
	}
}
