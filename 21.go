package main

import (
	"fmt"
	"strconv"
)

// LightBulb class to be adapted to the Lighting interface (Adaptee)
// IsOn is a boolean property
// Assume it is not allowed to add more methods for other interfaces
type LightBulb struct {
	IsOn bool
}

func (lightBulb *LightBulb) GetIsOn() bool {
	return lightBulb.IsOn
}

func (lightBulb *LightBulb) SetIsOn(isOn bool) {
	lightBulb.IsOn = isOn
}

// string method to implement the Stringer interface
func (lightBulb *LightBulb) String() string {
	return "LightBulb: " + strconv.FormatBool(lightBulb.IsOn)
}

// Lighting interface
type Lighting interface {
	Light()
}

// Adapter class for the light switch
type LightSwitchAdapter struct {
	lightBulb LightBulb // adaptee object
}

// Adapter implementation of the Lighting interface
func (lightSwitchAdapter *LightSwitchAdapter) Light() {
	lightBulb := lightSwitchAdapter.lightBulb

	if lightBulb.GetIsOn() {
		fmt.Println("Light is on, do nothing")
	} else {
		lightBulb.SetIsOn(true)
		fmt.Println("Light is off, turning it on")
	}
}

func main() {
	// client code
	// create a Lighting interface object
	var lightSwitch Lighting = &LightSwitchAdapter{
		lightBulb: LightBulb{true},
	}

	// use light switch
	lightSwitch.Light()

}

//Lighting is the Target interface
//LightBulb class is the Adaptee
//LightSwitchAdapter is the Adapter that implements the Lighting interface
//main() is the Client
