package main

import (
	// "fmt"
	"machine"
	"machine/usb/hid/joystick"
	"time"

	"github.com/tieste/ascii-stick-l5-rp2040-zero-tinygo/pid"
)

var js = joystick.Joystick

func init() {
	js = joystick.UseSettings(
		joystick.Definitions{
			ReportID:     1,
			ButtonCnt:    4,
			HatSwitchCnt: 1,
			AxisDefs: []joystick.Constraint{
				// {MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // X
				// {MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // Y
				{MinIn: -21000, MaxIn: 21000, MinOut: -21000, MaxOut: 21000}, // X
				{MinIn: -29000, MaxIn: 29000, MinOut: -29000, MaxOut: 29000}, // Y
			},
		},
		nil,
		nil,
		pid.HidDescriptor,
	)
	if js == nil {
		panic("UseSettings failed")
	}
}

func main() {
	hat1 := machine.D2
	hat1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	hat2 := machine.D4
	hat2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	hat3 := machine.D5
	hat3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	hat4 := machine.D6
	hat4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	button1 := machine.D7
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	button2 := machine.D8
	button2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	button3 := machine.D9
	button3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	button4 := machine.D11
	button4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	machine.InitADC()
	analogX := machine.ADC{Pin: machine.D27}
	analogX.Configure(machine.ADCConfig{})
	analogY := machine.ADC{Pin: machine.D28}
	analogY.Configure(machine.ADCConfig{})

	ticker := time.NewTicker(100 * time.Millisecond)

	// minX := 0
	// maxX := 0
	// minY := 0
	// maxY := 0

	for range ticker.C {
		if !hat1.Get() {
			js.SetHat(0, joystick.HatUp)
		} else if !hat2.Get() {
			js.SetHat(0, joystick.HatLeft)
		} else if !hat3.Get() {
			js.SetHat(0, joystick.HatRight)
		} else if !hat4.Get() {
			js.SetHat(0, joystick.HatDown)
		} else {
			js.SetHat(0, joystick.HatCenter)
		}

		js.SetButton(0, !button1.Get())
		js.SetButton(1, !button2.Get())
		js.SetButton(2, !button3.Get())
		js.SetButton(3, !button4.Get())

		x := analogX.Get()
		tmpX := int(x) - 32767 - 15000
		js.SetAxis(0, tmpX)

		y := analogY.Get()
		tmpY := 32767 - int(y) + 4600
		js.SetAxis(1, tmpY)

		// if minX > tmpX {
		// 	minX = tmpX
		// }
		// if maxX < tmpX {
		// 	maxX = tmpX
		// }
		// if minY > tmpY {
		// 	minY = tmpY
		// }
		// if maxY < tmpY {
		// 	maxY = tmpY
		// }
		// centerX := (minX + maxX) / 2
		// centerY := (minY + maxY) / 2

		// fmt.Printf("%d %d %d %d : %d %d\n", minX, maxX, minY, maxY, centerX, centerY)

		js.SendState()
	}
}
