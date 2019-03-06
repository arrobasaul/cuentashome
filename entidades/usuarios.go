package entidades

//Usuario todo los usuarios que conectaran con el sistema
type Usuario struct {
	CodUsuario    int    `json:"CodUsuario" llave:"SI"`
	NombreUsuario string `json:"NombreUsuario"`
	Correo        string `json:"Correo"`
	Password      string `json:"Password"`
	Estado        int    `json:"Estado"`
}
