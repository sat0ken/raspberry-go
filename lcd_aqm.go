package main

import (
	"log"
	"time"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func writeCmd(dev *i2c.Dev, cmd []byte) {

	code, err := dev.Write(cmd)
	time.Sleep(time.Millisecond * 100)
	if err != nil {
		log.Println("i2c write error : ", err)
	}
	log.Println("i2c write code : ", code)
}

func main() {
	initlcd := []byte{0x38, 0x39, 0x14, 0x73, 0x56, 0x6c, 0x38, 0x01, 0x0f}
	clear := []byte{0x01}
    on := []byte{0x0f}

	if _, err := host.Init(); err != nil {
		log.Println("host init error : ", err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Println("i2c open err : ", err)
	}
	defer bus.Close()

	dev := &i2c.Dev{Addr: 0x3e, Bus: bus}
	writeCmd(dev, initlcd)
	writeCmd(dev, clear)
	writeCmd(dev, on)

    str := "HHello!!"
	var bytestr []byte
	for i := range str {
		bytestr = append(bytestr, str[i])
	}
    log.Println(len(bytestr), bytestr)
	writeCmd(dev, bytestr)
}
