package entidades

//Usuario todo los usuarios que conectaran con el sistema
type Usuario struct {
	CodUsuario    int
	NombreUsuario string
	Correo        string
	Password      string
	Estado        int
}
