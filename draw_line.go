package gameutil

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"math"
)

var Pixel *ebiten.Image


func InitPixel(){
	//Create Pixel
	var err error
	//Simple white 1x1 pixel image for manipulation
	Pixel, err := ebiten.NewImage(1, 1, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	_ = Pixel.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}


// DrawLine draws a line segment on the given destination dst.
//
// DrawLine is intended to be used mainly for debugging or prototyping purpose.
func DrawLine(dst *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color, camera *Camera) {
	ew, eh := Pixel.Size()
	length := math.Hypot(x2-x1, y2-y1)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(length/float64(ew), 1/float64(eh))
	op.GeoM.Rotate(math.Atan2(y2-y1, x2-x1))
	op.GeoM.Translate(x1, y1)
	op.ColorM.Scale(colorScale(clr))
	camera.ApplyCameraTransform(op, true)
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	_ = dst.DrawImage(Pixel, op)
}