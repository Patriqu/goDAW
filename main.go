package main

import (
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
	"os"
)

func main() {
	fmt.Println("GoDAW")

	generateSinwave()
}

func generateSinwave() {
	// Create a new file
	file, err := os.Create("output.wav")
	if err != nil {
		panic(err)
	}

	enc := wav.NewEncoder(file, 44100, 16, 1, 1)

	// Create a 440Hz sine wave for 5 seconds
	sine := make([]int, 44100*5)
	for i := range sine {
		sine[i] = int(32767.0 * math.Sin(float64(i)*2*math.Pi*440/44100))
	}

	format := audio.Format{
		NumChannels: 2,
		SampleRate:  44100,
	}

	buffer := audio.IntBuffer{Data: sine, Format: &format}

	// Write to the .wav file
	if err := enc.Write(&buffer); err != nil {
		panic(err)
	}

	if err := enc.Close(); err != nil {
		panic(err)
	}
}
