package epd

type device interface {
	init()
	clear()
	display(bytes []byte)
	sleep()
}
