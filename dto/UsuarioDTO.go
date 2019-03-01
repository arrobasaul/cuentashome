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
func GetUsuario(id int) (*dto.Usuario, *dto.Errores) {

	// run your query, fill in &u...
	db := conn.Conexion()
	usuario := dto.Usuario{}
	var error2 dto.Errores
	println(id)
	err := db.QueryRow("SELECT CodUsuario, NombreUsuario, Correo, Password, Estado FROM usuarios where CodUsuario=?", id).Scan(&usuario.CodUsuario, &usuario.NombreUsuario, &usuario.Correo, &usuario.Password, &usuario.Estado)
	if err != nil {
		error2 = dto.Errores{Error: err.Error(), Descripcion: "no encontrado"}
		return nil, &error2
	}
	return &usuario, nil
}
func CreateUsuario(usuario dto.Usuario) (id int64) {
	//err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	// run your query, fill in &u...
	db := conn.Conexion()
	insForm, err := db.Prepare("INSERT INTO usuarios(NombreUsuario, Correo, Password, Estado) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	res, err := insForm.Exec(usuario.NombreUsuario, usuario.Correo, usuario.Password, usuario.Estado)
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
