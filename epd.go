package epd

import (
	"image"
	"sync"
)

const particalUpdateCount = 10

// Epaper is a e-paper device
type Epaper struct {
	board  *board
	device device
	mu     sync.Mutex
	puc    int
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

// Display display img on e-paper
func (p *Epaper) Display(img image.Image) error {
	return p.wrap(func() {
		p.device.display(img)
	})
}

// Clear clear the e-paper screen
func (p *Epaper) Clear() error {
	if err := p.device.init(false); err != nil {
		return err
	}
	defer p.device.sleep()
	p.device.clear()
	return nil
}

func (p *Epaper) wrap(fn func()) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	var pu bool
	if p.puc > 0 {
		pu = true
		p.puc = p.puc - 1
	} else {
		pu = false
		p.puc = particalUpdateCount
	}
	if err := p.device.init(pu); err != nil {
		return err
	}
	defer p.device.sleep()
	fn()
	return nil
}
