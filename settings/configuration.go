package settings

import (
	"encoding/json"
	"fmt"
	"os"

	entidades "../entidades"
)

func Setting() {
	file, _ := os.Open("../conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Conf := entidades.Conf{}
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(Conf.ConnMYSQL.User) // output: [UserA, UserB]
}
