package entidades

//Grupos todo los usuarios que conectaran con el sistema
type Grupos struct {
	CodCrupo    int       `json:"CodCrupo" llave:"SI"`
	NombreGrupo string    `json:"NombreGrupo"`
	Estado      int       `json:"Estado"`
	Usuario     []Usuario `json:"usuario" Map:"NO"`
}
