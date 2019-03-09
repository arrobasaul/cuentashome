package entidades

//Usuario todo los usuarios que conectaran con el sistema
type Usuarios struct {
	CodUsuarios   int    `json:"CodUsuario" llave:"SI" schema:"si"`
	NombreUsuario string `json:"NombreUsuario" schema:"si"`
	Correo        string `json:"Correo" schema:"si"`
	Password      string `json:"Password" schema:"si"`
	Estado        int    `json:"Estado" schema:"si"`
}
