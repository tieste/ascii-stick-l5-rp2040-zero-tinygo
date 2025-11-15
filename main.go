package main

import (
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
				{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // X
				{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // Y
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
	if js == nil {
		panic("Failed to configure joystick")
	}
	hat1 := machine.D3
	hat1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// hat1.Configure(machine.PinConfig{Mode: machine.PinInput})

	hat2 := machine.D4
	hat2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// hat2.Configure(machine.PinConfig{Mode: machine.PinInput})

	hat3 := machine.D5
	hat3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// hat3.Configure(machine.PinConfig{Mode: machine.PinInput})

	hat4 := machine.D6
	hat4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// hat4.Configure(machine.PinConfig{Mode: machine.PinInput})

	button1 := machine.D7
	// button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button1.Configure(machine.PinConfig{Mode: machine.PinInput})

	button2 := machine.D8
	// button2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button2.Configure(machine.PinConfig{Mode: machine.PinInput})

	button3 := machine.D9
	// button3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button3.Configure(machine.PinConfig{Mode: machine.PinInput})

	button4 := machine.D10
	// button4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button4.Configure(machine.PinConfig{Mode: machine.PinInput})

	machine.InitADC()
	analogX := machine.ADC{Pin: machine.D27}
	analogX.Configure(machine.ADCConfig{})
	analogY := machine.ADC{Pin: machine.D28}
	analogY.Configure(machine.ADCConfig{})

	ticker := time.NewTicker(100 * time.Millisecond)

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
		js.SetAxis(0, int(x)-32767)
		y := analogY.Get()
		js.SetAxis(1, 32767-int(y))

		js.SendState()
	}
}
