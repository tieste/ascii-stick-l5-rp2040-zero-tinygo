package main

import (
	"fmt"
	"machine"
	"machine/usb/hid/joystick"
	"time"
)

//	var def = joystick.Definitions{
//		ReportID:     1,
//		ButtonCnt:    4,
//		HatSwitchCnt: 1,
//		AxisDefs: []joystick.Constraint{
//			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
//			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
//		},
//	}
//
// js holds the joystick instance. We cannot name the joystick type here
// because the type in the package is unexported; use type inference
// from joystick.Port() so the variable has the correct pointer type.
var js = joystick.Joystick

func init() {
	// Definitions: 4 buttons, 1 hat switch, 2 axes (X, Y)
	// def := joystick.Definitions{
	// 	ReportID:     1,
	// 	ButtonCnt:    4,
	// 	HatSwitchCnt: 1,
	// 	AxisDefs: []joystick.Constraint{
	// 		{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // X
	// 		{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // Y
	// 	},
	// }

	// HID report descriptor must match the layout produced by State.MarshalBinary():
	// [ReportID (1 byte)] [Buttons (1 byte: 4 bits used)] [Hat (1 byte: 4 bits used)] [Axes (2 * 16-bit)]
	// hidDesc := descriptor.Append([][]byte{
	// 	descriptor.HIDUsagePageGenericDesktop,
	// 	descriptor.HIDUsageDesktopJoystick,
	// 	descriptor.HIDCollectionApplication,
	// 	descriptor.HIDReportID(1),

	// 	// Buttons (4 bits used, packed into a full byte)
	// 	descriptor.HIDUsagePageButton,
	// 	descriptor.HIDUsageMinimum(1),
	// 	descriptor.HIDUsageMaximum(4),
	// 	descriptor.HIDLogicalMinimum(0),
	// 	descriptor.HIDLogicalMaximum(1),
	// 	descriptor.HIDReportSize(1),
	// 	descriptor.HIDReportCount(4),
	// 	descriptor.HIDInputDataVarAbs,
	// 	// Padding to fill the rest of the byte (4 bits)
	// 	descriptor.HIDReportSize(4),
	// 	descriptor.HIDReportCount(1),
	// 	descriptor.HIDInputConstVarAbs,

	// 	// Hat switch (4 bits), followed by 4 bits padding to make a full byte
	// 	descriptor.HIDUsagePageGenericDesktop,
	// 	descriptor.HIDUsageDesktopHatSwitch,
	// 	descriptor.HIDLogicalMinimum(0),
	// 	descriptor.HIDLogicalMaximum(7),
	// 	descriptor.HIDPhysicalMinimum(0),
	// 	descriptor.HIDPhysicalMaximum(315),
	// 	descriptor.HIDUnit(0x14),
	// 	descriptor.HIDReportCount(1),
	// 	descriptor.HIDReportSize(4),
	// 	descriptor.HIDInputDataVarAbs,
	// 	// Padding nibble after hat to complete the byte
	// 	descriptor.HIDReportSize(4),
	// 	descriptor.HIDReportCount(1),
	// 	descriptor.HIDInputConstVarAbs,

	// 	// Axes: X, Y (each 16 bits, signed)
	// 	descriptor.HIDUsagePageGenericDesktop,
	// 	descriptor.HIDUsageDesktopPointer,
	// 	descriptor.HIDLogicalMinimum(-32767),
	// 	descriptor.HIDLogicalMaximum(32767),
	// 	descriptor.HIDReportCount(2),
	// 	descriptor.HIDReportSize(16),
	// 	descriptor.HIDCollectionPhysical,
	// 	descriptor.HIDUsageDesktopX,
	// 	descriptor.HIDUsageDesktopY,
	// 	descriptor.HIDInputDataVarAbs,
	// 	descriptor.HIDCollectionEnd,
	// 	descriptor.HIDCollectionEnd,
	// })
	// hidDesc := HidDescriptor

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
		hidDescriptor,
	)
	if js == nil {
		panic("UseSettings failed")
	}

}

// var js = joystick.Port()

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
