package enrutador

import (
	"fmt"
	"net/http"

	dto "../dto"
)

func Manejador(w http.ResponseWriter, r *http.Request) {
	usuarios, error := dto.GetUsuarios()
	if error != nil {
		fmt.Println("errores")
	} else {
		fmt.Println(usuarios)
		fmt.Println("correcto")
	}
	/*d := ent.Cuenta{
		CodCuenta: 1,
	}
	fmt.Println(d)

	m := image.NewRGBA(image.Rect(0, 0, 10, 10))
	m.Set(0, 0, color.RGBA{0, 0, 0, 255})
	m.Set(1, 1, color.RGBA{255, 255, 255, 255})
	m.Set(2, 2, color.RGBA{0, 0, 0, 255})
	m.Set(3, 3, color.RGBA{255, 255, 255, 255})
	m.Set(4, 4, color.RGBA{0, 0, 0, 255})
	m.Set(5, 5, color.RGBA{255, 255, 255, 255})
	m.Set(6, 6, color.RGBA{0, 0, 0, 255})
	m.Set(7, 7, color.RGBA{255, 255, 255, 255})
	m.Set(8, 8, color.RGBA{0, 0, 0, 255})
	m.Set(9, 9, color.RGBA{255, 255, 255, 255})

	pic.ShowImage(m)
	fmt.Fprintf(w, "algo", m)*/
	fmt.Fprintf(w, "Hola, %s, ¡este es un servidor!", r.URL.Path)
}
