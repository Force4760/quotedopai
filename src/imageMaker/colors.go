package imageMaker

import "github.com/fogleman/gg"

type Color struct {
	H int
	S int
	B int
	A int
}

func makeColor(h, s, b, a int) Color {
	return Color{
		h % 360,
		s % 100,
		b % 100,
		a % 100,
	}
}

func randColor() Color {
	return Color{
		randRange(0, 360),
		randRange(30, 70),
		randRange(30, 70),
		100,
	}
}

func (c Color) SetColor(dc *gg.Context) {
	r, g, b, _ := c.ToRGBA()

	dc.SetRGB255(r, g, b)
}

func (c Color) Analogous() (Color, Color, Color, Color) {
	h2 := c.H - 16
	h3 := c.H + 16
	h4 := c.H - 32
	h5 := c.H + 32

	s2 := c.S + 5

	b2 := c.B + 9
	b3 := c.B + 5

	color2 := makeColor(h2, s2, b2, c.A)
	color3 := makeColor(h3, s2, b2, c.A)
	color4 := makeColor(h4, s2, b3, c.A)
	color5 := makeColor(h5, s2, b3, c.A)

	return color2, color3, color4, color5
}

func (c Color) Triad() (Color, Color) {
	h2 := c.H + 120
	h3 := c.H + 240

	color2 := makeColor(h2, c.S, c.B, c.A)
	color3 := makeColor(h3, c.S, c.B, c.A)

	return color2, color3
}

func (c Color) Complement() (Color, Color, Color) {
	h2 := c.H + 90
	s2 := c.S / 2

	color2 := makeColor(h2, c.S, c.B, c.A)
	color3 := makeColor(h2, s2, c.B, c.A)
	color4 := makeColor(c.H, s2, c.B, c.A)

	return color2, color3, color4
}

func (c Color) Square() (Color, Color, Color) {
	h2 := c.H + 90
	h3 := c.H + 180
	h4 := c.H + 270

	color2 := makeColor(h2, c.S, c.B, c.A)
	color3 := makeColor(h3, c.S, c.B, c.A)
	color4 := makeColor(h4, c.S, c.B, c.A)

	return color2, color3, color4
}

func (c Color) Shades() (Color, Color, Color, Color) {
	b2 := c.B - 25
	b3 := c.B - 50
	b4 := c.B - 75
	b5 := c.B - 10

	color2 := makeColor(c.H, c.S, b2, c.A)
	color3 := makeColor(c.H, c.S, b3, c.A)
	color4 := makeColor(c.H, c.S, b4, c.A)
	color5 := makeColor(c.H, c.S, b5, c.A)

	return color2, color3, color4, color5
}

func (c Color) Mono() (Color, Color, Color, Color) {
	s2 := c.S - 25
	s3 := c.S - 50
	s4 := c.S - 75
	s5 := c.S - 10

	color2 := makeColor(c.H, s2, c.B, c.A)
	color3 := makeColor(c.H, s3, c.B, c.A)
	color4 := makeColor(c.H, s4, c.B, c.A)
	color5 := makeColor(c.H, s5, c.B, c.A)

	return color2, color3, color4, color5
}

func (c Color) Div() (Color, Color, Color) {
	s2 := c.S / 2
	b2 := c.B / 2

	color2 := makeColor(c.H, s2, b2, c.A)
	color3 := makeColor(c.H, s2, c.B, c.A)
	color4 := makeColor(c.H, c.S, b2, c.A)

	return color2, color3, color4
}

func (c Color) Alpha() (Color, Color, Color, Color) {
	color2 := makeColor(c.H, c.S, c.B, 100)
	color3 := makeColor(c.H, c.S, c.B, 75)
	color4 := makeColor(c.H, c.S, c.B, 50)
	color5 := makeColor(c.H, c.S, c.B, 25)

	return color2, color3, color4, color5
}

func (c *Color) SetH(val int) {
	c.H = val % 360
}

func (c *Color) SetS(val int) {
	c.S = val % 100
}

func (c *Color) SetB(val int) {
	c.B = val % 100
}

func (c *Color) SetA(val int) {
	c.A = val % 100
}

func (c Color) ToRGBA() (r, g, b, a int) {
	// normalize H S V to [0, 1]
	h := float64(c.H) / 360.0
	s := float64(c.S) / 100.0
	v := float64(c.B) / 100.0

	i := int(h * 6.0)
	f := h*6 - float64(i)

	p := int((v * (1 - s) * 255) + 1)
	q := int((v * (1 - f*s) * 255) + 1)
	t := int((v * (1 - (1-f)*s) * 255) + 1)
	v2 := int(v*255 + 1)

	switch i % 6 {
	case 0:
		r = v2
		g = t
		b = p
	case 1:
		r = q
		g = v2
		b = p

	case 2:
		r = p
		g = v2
		b = t

	case 3:
		r = p
		g = q
		b = v2

	case 4:
		r = t
		g = p
		b = v2

	case 5:
		r = v2
		g = p
		b = q

	}

	// c.A -> 100
	// a   -> 255
	a = c.A * 255 / 100

	return
}
