package main

import (
	"net/http"

	ruter "./enrutador"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*m := image.NewRGBA(image.Rect(0, 0, 10, 10))
	m.Set(0, 0, color.RGBA{0, 0, 0, 255})
	m.Set(1, 1, color.RGBA{255, 255, 255, 255})
	m.Set(2, 2, color.RGBA{0, 0, 0, 255})
	m.Set(3, 3, color.RGBA{255, 255, 255, 255})
	m.Set(4, 4, color.RGBA{0, 0, 0, 255})
	m.Set(5, 5, color.RGBA{255, 255, 255, 255})
	m.Set(6, 6, color.RGBA{0, 0, 0, 255})
	m.Set(7, 7, color.RGBA{255, 255, 255, 255})
	m.Set(8, 8, color.RGBA{0, 0, 0, 255})
	m.Set(9, 9, color.RGBA{255, 255, 255, 255})*/
	//ent.Julia()

	//pic.ShowImage(m)
	// Dx and Dy return a rectangle's width and height.
	//fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r))
	/*img, err := jpeg.Decode(m)
	if err != nil {
		return err
	}
	return png.Encode(w, img)*/
	/*r := mux.NewRouter()
	r.HandleFunc("/", ruter.HomeHandler)
	r.HandleFunc("/products", ruter.ProductsHandler)
	r.HandleFunc("/articles", ruter.ArticlesHandler)
	http.Handle("/", r)*/
	http.HandleFunc("/", ruter.Manejador)
	http.ListenAndServe(":8080", nil)

}
