package entidades

//Grupos todo los usuarios que conectaran con el sistema
type Grupos struct {
	CodGrupos     int        `json:"CodGrupo" llave:"SI" schema:"si"`
	NombreGrupo   string     `json:"NombreGrupo" schema:"si" langth:"100"`
	Estado        int        `json:"Estado" schema:"si"`
	FechaCreacion string     `json:"FechaCreacion" schema:"si" isFecha:"si"`
	Usuarios      []Usuarios `json:"usuario" Map:"NO"`
}
