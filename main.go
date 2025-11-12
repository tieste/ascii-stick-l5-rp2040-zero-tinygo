package main

import (
	"fmt"
	"machine"
	"machine/usb/hid/joystick"
	"time"
)

var def = joystick.Definitions{
	ReportID:     1,
	ButtonCnt:    4,
	HatSwitchCnt: 1,
	AxisDefs: []joystick.Constraint{
		{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
		{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
	},
}

// var js = joystick.UseSettings(def, nil, nil, nil)
var js = joystick.Port()

func main() {

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
	cnt := 0
	const f = 3.0

	for range ticker.C {
		if !button1.Get() {
			// println("button up is pressed!!")
			js.SetHat(0, joystick.HatUp)
		} else if !button2.Get() {
			// println("button left is pressed!!")
			js.SetHat(0, joystick.HatLeft)
		} else if !button3.Get() {
			// println("button right is pressed!!")
			js.SetHat(0, joystick.HatRight)
		} else if !button4.Get() {
			// println("button down is pressed!!")
			js.SetHat(0, joystick.HatDown)
		} else {
			js.SetHat(0, joystick.HatCenter)
		}

		// if !button5.Get() {
		// 	// println("button a is pressed!!")
		// 	js.SetButton(0, true)
		// } else {
		// 	js.SetButton(0, false)
		// }
		js.SetButton(0, !button5.Get())
		// if !button6.Get() {
		// 	// println("button b is pressed!!")
		// 	js.SetButton(1, true)
		// } else {
		// 	js.SetButton(1, false)
		// }
		js.SetButton(1, !button6.Get())

		x := analogX.Get()

		// t := float64(cnt) * 0.01
		// x := 32767 * math.Sin(2*math.Pi*f*t)
		y := analogY.Get()
		fmt.Printf("%0d %0d\n", int(x), int(y))
		js.SetAxis(0, int(x)-32767)
		js.SetAxis(1, 32767-int(y))

		js.SendState()

		// time.Sleep(time.Millisecond * 100)
		cnt++
	}
}
