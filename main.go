package main

import (
	"GoDAW/audio"
	"GoDAW/view"
	"fmt"
)

func main() {
	fmt.Println("GoDAW")
	audio.Waveforms()
	view.CreateKeyboard()
}
