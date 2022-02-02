package imageMaker

import "github.com/fogleman/gg"

const FontSize = 70
const FontPath = "./font.ttf"
const Width = 1000
const ShadeAlpha = 0.6

func MakeImage(quote string) {
	dc := gg.NewContext(Width, Width)
	dc.Translate(Width/2, Width/2)

	colors, bg := makeAllColors()
	background(dc, bg, Width)

	drawAllPlanets(dc, colors, bg)
	foregroundShadow(dc, Width)

	drawQuote(dc, quote, Width)

	dc.SavePNG("quote.png")
}
