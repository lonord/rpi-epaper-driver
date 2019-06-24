package epd

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	// default gpio pins
	defaultPinRST  = 17
	defaultPinDC   = 25
	defaultPinCS   = 8
	defaultPinBUSY = 24
)

type board struct {
	// gpio pins
	pinRST  rpio.Pin
	pinDC   rpio.Pin
	pinCS   rpio.Pin
	pinBUSY rpio.Pin
}

func (p *board) init() error {
	// init gpio
	if err := rpio.Open(); err != nil {
		return err
	}
	// RST
	p.pinRST = rpio.Pin(defaultPinRST)
	p.pinRST.Output()
	// DC
	p.pinDC = rpio.Pin(defaultPinDC)
	p.pinDC.Output()
	// CS
	p.pinCS = rpio.Pin(defaultPinCS)
	p.pinCS.Output()
	// BUSY
	p.pinBUSY = rpio.Pin(defaultPinBUSY)
	p.pinBUSY.Input()
	p.pinBUSY.PullDown()
	// init spi
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		return err
	}
	rpio.SpiSpeed(2000000)
	return nil
}

func (p *board) writeRST(v bool) {
	if v {
		p.pinRST.High()
	} else {
		p.pinRST.Low()
	}
}

func (p *board) writeDC(v bool) {
	if v {
		p.pinDC.High()
	} else {
		p.pinDC.Low()
	}
}

func (p *board) writeCS(v bool) {
	if v {
		p.pinCS.High()
	} else {
		p.pinCS.Low()
	}
}

func (p *board) readBUSY() {
	for p.pinBUSY.Read() == rpio.High {
		p.delayms(100)
	}
}

func (p *board) writeSPI(b ...byte) {
	rpio.SpiTransmit(b...)
}

func (p *board) cleanup() {
	// stop spi
	rpio.SpiEnd(rpio.Spi0)
	// stop gpio
	rpio.Close()
}

func (p *board) delayms(ms uint) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}
