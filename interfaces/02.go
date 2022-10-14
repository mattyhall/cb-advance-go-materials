package main

import "fmt"

type AudioSource interface {
	Volume() float32
}

type Microphone struct{}
func (m *Microphone) Volume() float32 { return 70 }

type Webcam struct{}
func (w *Webcam) Volume() float32 { return 30 }

type Soundtrack struct{}
func (s *Soundtrack) Volume() float32 { return 55.5 }

func BoostAudioSourceVolume(a AudioSource, gain float32) float32 {
	return a.Volume() * gain
}

func main() {
	m := Microphone{}
	w := Webcam{}
	s := Soundtrack{}

	fmt.Println(BoostAudioSourceVolume(&m, 5.3))
	fmt.Println(BoostAudioSourceVolume(&w, 5.3))
	fmt.Println(BoostAudioSourceVolume(&s, 5.3))
}
