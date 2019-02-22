package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/
import (
	"image"
	"image/color"
)

type Cuenta struct {
	CodCuenta    int
	NombreCuenta string
	Valor        float32
	Estado       int
}

var ymax, ymin, xmax, xmin, screeny, screenx, MaxIterations int
var VisibleXmin, VisibleXmax, VisibleYmin, VisibleYmax, ImC, ReC float64

func AdjustAspect() {
	var want_aspect int
	var canvas_aspect int
	var hgt int
	var wid int
	var mid int

	want_aspect = (ymax - ymin) / (xmax - xmin)
	canvas_aspect = screeny / screenx
	if want_aspect > canvas_aspect {
		wid = (ymax - ymin) / canvas_aspect
		mid = (xmin + xmax) / 2
		VisibleXmin = (float64)(mid - wid/2)
		VisibleXmax = (float64)(mid + wid/2)
		VisibleYmin = (float64)(ymin)
		VisibleYmax = (float64)(ymax)
	} else {
		hgt = (xmax - xmin) * canvas_aspect
		mid = (ymin + ymax) / 2
		VisibleYmin = (float64)(mid - hgt/2)
		VisibleYmax = (float64)(mid + hgt/2)
		VisibleXmin = (float64)(xmin)
		VisibleXmax = (float64)(xmax)
	}
}

func Julia() {
	var ReZ, ImZ float64
	var Re2Z, Im2Z float64
	//var Red, Green, Blue float64
	var i, j int
	var N int
	AdjustAspect()
	m := image.NewRGBA(image.Rect(0, 0, screenx, screeny))
	for i = 0; i < screenx; i++ {
		for j = 0; j < screeny; j++ {
			ReZ = ((VisibleXmax - VisibleXmin) / (float64)(screenx-1))
			ImZ = ((VisibleYmax - VisibleYmin) / (float64)(screeny-1))
			ReZ = (float64(i) * ReZ)
			ImZ = (float64(j) * ImZ)

			ReZ = ReZ + VisibleXmin
			ImZ = ImZ + VisibleYmin
			Re2Z = ReZ * ReZ
			Im2Z = ImZ * ImZ
			N = 0
			for {
				Re2Z = ReZ * ReZ
				Im2Z = ImZ * ImZ
				ImZ = 2*ReZ*ImZ + ImC
				ReZ = (Re2Z - Im2Z) + ReC
				N = N + 1
				if (N < MaxIterations) && (Re2Z+Im2Z < 4) {
					break
				}
			}
			if N == MaxIterations {

				m.Set(5, 5, color.RGBA{0, 0, 0, 255})
				//Red = 0
				//Green = 0
				//Blue = 0
			} else {
				m.Set(5, 5, color.RGBA{0, 1, 2, 255})
				//Red = RGBColor(N, 0)
				//Green = RGBColor(N, 1)
				//Blue = RGBColor(N, 2)
			}
			//Picture1.PSet (i, screeny - j), RGB(Red, Green, Blue)
		}
	}
}
