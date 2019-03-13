package servicios

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashear(valor string) string {
	contraseñaPlanaComoByte := []byte(valor)
	hash, err := bcrypt.GenerateFromPassword(contraseñaPlanaComoByte, bcrypt.DefaultCost) //DefaultCost es 10
	if err != nil {
		fmt.Println(err)
	}
	hashComoCadena := string(hash)
	return hashComoCadena
}
func VerificarHash(textPlano string, hash string) bool {
	hashComoByte := []byte(hash)
	textPlanoComoByte := []byte(textPlano)
	error := bcrypt.CompareHashAndPassword(hashComoByte, textPlanoComoByte)
	if error == nil {
		return true
	} else {
		return false
	}
}
