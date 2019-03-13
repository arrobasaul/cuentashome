package entidades

type Conf struct {
	ConnMYSQL ConnMYSQL `json:"ConnMYSQL"`
	Globales  Globales  `json:"Globales"`
}

type ConnMYSQL struct {
	Host string `json:"Host"`
	DB   string `json:"Db"`
	User string `json:"User"`
	Pass string `json:"Pass"`
}

type Globales struct {
	Env string `json:"Env"`
}
