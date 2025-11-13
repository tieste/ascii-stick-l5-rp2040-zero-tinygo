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
	button1 := machine.D3
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button2 := machine.D4
	button2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button3 := machine.D5
	button3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button4 := machine.D6
	button4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button5 := machine.D15
	button5.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button6 := machine.D26
	button6.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	machine.InitADC()
	analogX := machine.ADC{Pin: machine.D28}
	analogX.Configure(machine.ADCConfig{})
	analogY := machine.ADC{Pin: machine.D27}
	analogY.Configure(machine.ADCConfig{})

	ticker := time.NewTicker(100 * time.Millisecond)

	for range ticker.C {
		if !button1.Get() {
			js.SetHat(0, joystick.HatUp)
		} else if !button2.Get() {
			js.SetHat(0, joystick.HatLeft)
		} else if !button3.Get() {
			js.SetHat(0, joystick.HatRight)
		} else if !button4.Get() {
			js.SetHat(0, joystick.HatDown)
		} else {
			js.SetHat(0, joystick.HatCenter)
		}

		js.SetButton(0, !button5.Get())
		js.SetButton(1, !button6.Get())

		x := analogX.Get()
		js.SetAxis(0, int(x)-32767)
		y := analogY.Get()
		js.SetAxis(1, 32767-int(y))

		js.SendState()
	}
}
