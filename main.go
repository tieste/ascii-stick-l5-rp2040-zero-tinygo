package main

import (
	"machine"
	"time"
)

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

	for {
		if !button1.Get() {
			println("button up is pressed!!")
		}
		if !button2.Get() {
			println("button left is pressed!!")
		}
		if !button3.Get() {
			println("button right is pressed!!")
		}
		if !button4.Get() {
			println("button down is pressed!!")
		}
		if !button5.Get() {
			println("button a is pressed!!")
		}
		if !button6.Get() {
			println("button b is pressed!!")
		}

		time.Sleep(time.Millisecond * 100)
	}
}
