package entidades

/*
Cuenta estructura que manejara los valores de la cuenta
*/

type Deudas struct {
	CodDeudas       int        `json:"CodDeudas" llave:"SI" schema:"si"`
	DescripionDeuda string     `json:"DescripionDeuda" schema:"si"`
	Valor           float32    `json:"Valor" schema:"si"`
	Estado          int        `json:"Estado" schema:"si"`
	FechaDeuda      string     `json:"FechaDeuda" schema:"si" isFecha:"si"`
	Usuarios        []Usuarios `json:"usuario" Map:"NO"`
}

/*
func adjustAspect() {
	var wantAspect int
	var canvasAspect int
	var hgt int
	var wid int
	var mid int
	fmt.Println(global.Ymax)
	fmt.Println(global.Ymin)
	fmt.Println(global.Ymax)
	fmt.Println(global.Xmin)
	global.Main2()
	fmt.Println(global.Ymax)
	fmt.Println(global.Ymin)
	fmt.Println(global.Ymax)
	fmt.Println(global.Xmin)
	wantAspect = (global.Ymax - global.Ymin) / (global.Ymax - global.Xmin)
	canvasAspect = global.Screeny / global.Screenx
	if wantAspect > canvasAspect {
		wid = (global.Ymax - global.Ymin) / canvasAspect
		mid = (global.Xmin + global.Xmax) / 2
		global.VisibleXmin = (float64)(mid - wid/2)
		global.VisibleXmax = (float64)(mid + wid/2)
		global.VisibleYmin = (float64)(global.Ymin)
		global.VisibleYmax = (float64)(global.Ymax)
	} else {
		hgt = (global.Xmax - global.Xmin) * canvasAspect
		mid = (global.Ymin + global.Ymax) / 2
		global.VisibleYmin = (float64)(mid - hgt/2)
		global.VisibleYmax = (float64)(mid + hgt/2)
		global.VisibleXmin = (float64)(global.Xmin)
		global.VisibleXmax = (float64)(global.Xmax)
	}
}

func Julia() {
	var ReZ, ImZ float64
	var Re2Z, Im2Z float64
	//var Red, Green, Blue float64
	var i, j int
	var N int
	adjustAspect()
	m := image.NewRGBA(image.Rect(0, 0, global.Screenx, global.Screeny))
	for i = 0; i < global.Screenx; i++ {
		for j = 0; j < global.Screeny; j++ {
			ReZ = ((global.VisibleXmax - global.VisibleXmin) / (float64)(global.Screenx-1))
			ImZ = ((global.VisibleYmax - global.VisibleYmin) / (float64)(global.Screeny-1))
			ReZ = (float64(i) * ReZ)
			ImZ = (float64(j) * ImZ)

			ReZ = ReZ + global.VisibleXmin
			ImZ = ImZ + global.VisibleYmin
			Re2Z = ReZ * ReZ
			Im2Z = ImZ * ImZ
			N = 0
			for {
				Re2Z = ReZ * ReZ
				Im2Z = ImZ * ImZ
				ImZ = 2*ReZ*ImZ + global.ImC
				ReZ = (Re2Z - Im2Z) + global.ReC
				N = N + 1
				if (N < global.MaxIterations) && (Re2Z+Im2Z < 4) {
					break
				}
			}
			if N == global.MaxIterations {

				m.Set(i, global.Screeny-j, color.RGBA{0, 0, 0, 255})
				//Red = 0
				//Green = 0
				//Blue = 0
			} else {
				m.Set(i, global.Screeny-j, color.RGBA{0, 1, 2, 255})
				//Red = RGBColor(N, 0)
				//Green = RGBColor(N, 1)
				//Blue = RGBColor(N, 2)
			}
			//Picture1.PSet (i, screeny - j), RGB(Red, Green, Blue)
		}
	}
	outFile, err := os.Create("img2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, m, nil)
}
*/
