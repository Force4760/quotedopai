package imageMaker

import (
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

func MakeImage() {
	const w = 1000
	dc := gg.NewContext(w, w)
	dc.Translate(w/2, w/2)

	colors, bg := makeAllColors()
	background(dc, bg, w)

	drawAllPlanets(dc, colors, bg)
	foregroundShadow(dc, w)

	drawQuote(dc, "É bastante improvável ir num avião com uma bomba. É por isso que eu levo sempre uma bomba comigo, porque ir num avião com duas bombas é ainda mais improvável.", w)

	dc.SavePNG("out.png")
}

func drawQuote(dc *gg.Context, quote string, w float64) {
	err := dc.LoadFontFace("./font.ttf", 70)
	if err != nil {
		panic(err)
	}

	// Draw a shadow
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(quote, 3, 3, 0.5, 0.5, w*0.9, 1.7, gg.AlignCenter)

	// Draw the "real" text
	dc.SetRGB(1, 1, 1)
	dc.DrawStringWrapped(quote, 0, 0, 0.5, 0.5, w*0.9, 1.7, gg.AlignCenter)

}

func drawAllPlanets(dc *gg.Context, colors [6]Color, bg Color) {
	// Draw the center planet
	dc.DrawCircle(0, 0, 50)
	colors[0].SetColor(dc)
	dc.Fill()

	// Draw some orbital planets
	k := 1
	for i := 100; i < 500; i += 80 {
		drawOrbit(dc, float64(i), colors[k], bg)
		k++
	}
}
func makeAllColors() ([6]Color, Color) {
	// Get the original Colors
	c1 := randColor()

	// Get 6 more colors (4 Analogous, 2 Triads)
	c2, c3, c4, c5 := c1.Analogous()
	c6, c7 := c1.Triad()

	colors := [6]Color{c2, c3, c4, c5, c6, c7}

	// Default background color
	bg := makeColor(224, 30, 25, 100)

	return colors, bg
}

func background(dc *gg.Context, bg Color, w float64) {
	// Draw a background by drawing a rectangle over the canvas
	dc.DrawRectangle(-w/2, -w/2, w, w)
	bg.SetColor(dc)
	dc.Fill()
}

func foregroundShadow(dc *gg.Context, w float64) {
	// Draw a background by drawing a rectangle over the canvas
	dc.DrawRectangle(-w/2, -w/2, w, w)
	dc.SetRGBA(0, 0, 0, 0.3)
	dc.Fill()
}

func drawOrbit(dc *gg.Context, radius float64, c1, bg Color) {
	c1.SetColor(dc)

	// Draw the orbit line
	dc.DrawCircle(0, 0, radius)
	dc.SetLineWidth(6)
	dc.Stroke()

	// Coordinates of the planet in polar coordinates
	angle := rand.Float64() * 2 * math.Pi
	x := radius * math.Cos(angle)
	y := radius * math.Sin(angle)
	radius2 := float64(randRange(20, 30))

	// Draw the planet (inside)
	dc.DrawCircle(x, y, radius2)
	dc.Fill()

	// Draw the planet (border)
	dc.DrawCircle(x, y, radius2)
	bg.SetColor(dc)
	dc.Stroke()
}
