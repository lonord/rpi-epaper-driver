package epd

import "image"

// Epaper is a e-paper device
type Epaper struct {
	board  *board
	device device
}

// New create a new e-paper device
func New() *Epaper {
	b := &board{}
	d := newDev2in9(b)
	return &Epaper{
		board:  b,
		device: d,
	}
}

// Init setup gpio and e-paper board
func (p *Epaper) Init() error {
	if err := p.device.init(); err != nil {
		return err
	}
	return nil
}

// Display display img on e-paper
func (p *Epaper) Display(img image.Image) {
	p.device.display(img)
}

// Clear clear the e-paper screen
func (p *Epaper) Clear() {
	p.device.clear()
}

// Close teardown gpio and e-paper board
func (p *Epaper) Close() {
	p.device.sleep()
}
