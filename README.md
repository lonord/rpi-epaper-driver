# rpi-epaper-driver

[![GoDoc](https://godoc.org/github.com/lonord/rpi-epaper-driver?status.svg)](https://godoc.org/github.com/lonord/rpi-epaper-driver)

The driver of waveshare e-paper for praspberry pi

## Usage

```go
import epd "github.com/lonord/rpi-epaper-driver"

// create a new e-paper device
paper := epd.New()
// display an image
paper.Display(img)
// clear the screen
paper.Clear()
```

## Supported devices

- [x] [2.9inch e-Paper Module (black and white)](http://www.waveshare.net/wiki/2.9inch_e-Paper_Module)

## License

MIT
