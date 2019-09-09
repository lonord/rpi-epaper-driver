package epd

import "image/color"

func convertColorToBlackWhite(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	grey := (r*299 + g*587 + b*114 + 500) / 1000
	return grey >= 0x7fff
}
