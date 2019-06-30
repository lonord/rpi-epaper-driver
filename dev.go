package epd

import "image"

type device interface {
	init(particalUpdate bool) error
	clear()
	display(img image.Image)
	sleep()
}
