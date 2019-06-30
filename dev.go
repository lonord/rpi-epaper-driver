package epd

import "image"

type device interface {
	init() error
	clear()
	display(img image.Image)
	sleep()
}
