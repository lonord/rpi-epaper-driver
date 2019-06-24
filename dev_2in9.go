package epd

var (
	dev2in9LutFullUpdate = []byte{
		0x02, 0x02, 0x01, 0x11, 0x12, 0x12, 0x22, 0x22,
		0x66, 0x69, 0x69, 0x59, 0x58, 0x99, 0x99, 0x88,
		0x00, 0x00, 0x00, 0x00, 0xF8, 0xB4, 0x13, 0x51,
		0x35, 0x51, 0x51, 0x19, 0x01, 0x00,
	}
	dev2in9LutPartialUpdate = []byte{
		0x10, 0x18, 0x18, 0x08, 0x18, 0x18, 0x08, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x13, 0x14, 0x44, 0x12,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	dev2in9Width  = 128
	dev2in9Height = 296
)

type dev2in9 struct {
	board *board
}

func newDev2in9(board *board) device {
	return &dev2in9{board}
}

/***************************************** interface functions ****************************************/

func (d *dev2in9) init() {
	// TODO full update and partial update
	d.initLut(dev2in9LutFullUpdate)
}

func (d *dev2in9) clear() {
	d.setWindow(0, 0, uint(dev2in9Width-1), uint(dev2in9Height-1))
	for j := 0; j < dev2in9Height; j++ {
		d.setCursor(0, uint(j))
		d.sendCmd(0x24)
		for i := 0; i < dev2in9Width/8; i++ {
			d.sendData(0xff)
		}
	}
	d.turnOnDisplay()
}

func (d *dev2in9) display(bytes []byte) {
	d.setWindow(0, 0, uint(dev2in9Width-1), uint(dev2in9Height-1))
	for j := 0; j < dev2in9Height; j++ {
		d.setCursor(0, uint(j))
		d.sendCmd(0x24)
		for i := 0; i < dev2in9Width/8; i++ {
			d.sendData(bytes[i+j*(dev2in9Width/8)])
		}
	}
	d.turnOnDisplay()
}

func (d *dev2in9) sleep() {
	d.sendCmd(0x10)
	d.sendData(0x01)
}

/***************************************** private functions ****************************************/

func (d *dev2in9) initLut(lut []byte) {
	d.reset()

	d.sendCmd(0x01)
	d.sendData(byte((dev2in9Height - 1) & 0xff))
	d.sendData(byte(((dev2in9Height - 1) >> 8) & 0xff))
	d.sendData(0x00)

	d.sendCmd(0x0c)
	d.sendData(0xd7)
	d.sendData(0xd6)
	d.sendData(0x9d)

	d.sendCmd(0x2c)
	d.sendData(0xa8)

	d.sendCmd(0x3a)
	d.sendData(0x1a)

	d.sendCmd(0x3b)
	d.sendData(0x08)

	d.sendCmd(0x11)
	d.sendData(0x03)

	d.sendCmd(0x32)
	d.sendData(lut...)
}

func (d *dev2in9) reset() {
	d.board.writeRST(true)
	d.board.delayms(200)
	d.board.writeRST(false)
	d.board.delayms(10)
	d.board.writeRST(true)
	d.board.delayms(200)
}

func (d *dev2in9) sendCmd(b byte) {
	d.board.writeDC(false)
	d.board.writeCS(false)
	d.board.writeSPI(b)
	d.board.writeCS(true)
}

func (d *dev2in9) sendData(b ...byte) {
	d.board.writeDC(true)
	d.board.writeCS(false)
	d.board.writeSPI(b...)
	d.board.writeCS(true)
}

func (d *dev2in9) turnOnDisplay() {
	d.sendCmd(0x22)
	d.sendData(0xc4)
	d.sendCmd(0x20)
	d.sendCmd(0xff)
	d.board.readBUSY()
}

func (d *dev2in9) setWindow(xStart, yStart, xEnd, yEnd uint) {
	d.sendCmd(0x44)
	// x point must be the multiple of 8 or the last 3 bits will be ignored
	d.sendData(byte((xStart >> 3) & 0xff))
	d.sendData(byte((xEnd >> 3) & 0xff))
	d.sendCmd(0x45)
	d.sendData(byte(yStart & 0xff))
	d.sendData(byte((yStart >> 8) & 0xff))
	d.sendData(byte(yEnd & 0xff))
	d.sendData(byte((yEnd >> 8) & 0xff))
}

func (d *dev2in9) setCursor(x, y uint) {
	d.sendCmd(0x4e)
	// x point must be the multiple of 8 or the last 3 bits will be ignored
	d.sendData(byte((x >> 3) & 0xff))
	d.sendCmd(0x4f)
	d.sendData(byte(y & 0xff))
	d.sendData(byte((y >> 8) & 0xff))
	d.board.readBUSY()
}
