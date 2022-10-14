package main

import "fmt"

type Microphone struct { }

func (m *Microphone) Volume() float32 {
	// It's a directional mic so picks up your voice really well
	return 70.0 
}

// Make the mic louder
func BoostMicVolume(m *Microphone, gain float32) float32 {
	return m.Volume() * gain
}

func main() {
	m := Microphone{}
	fmt.Println(BoostMicVolume(&m, 2.5))
}
