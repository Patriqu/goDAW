package audio

import (
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"math"
	"os"
)

func Waveforms() {
	sine, square, triangle := generateWaveforms()
	sinEncoder, squareEncoder, triangleEncoder := createFileEncoders()
	generateWaveformFiles(sine, sinEncoder, square, squareEncoder, triangle, triangleEncoder)
}

func generateWaveforms() ([]int, []int, []int) {
	sine := generateSinWave(5)
	square := generateSquareWave(5)
	triangle := generateTriangleWave(5)
	return sine, square, triangle
}

func createFileEncoders() (*wav.Encoder, *wav.Encoder, *wav.Encoder) {
	sinEncoder := createFileEncoder("sin.wav")
	squareEncoder := createFileEncoder("square.wav")
	triangleEncoder := createFileEncoder("triangle.wav")
	return sinEncoder, squareEncoder, triangleEncoder
}

func generateWaveformFiles(sine []int, sinEncoder *wav.Encoder, square []int, squareEncoder *wav.Encoder, triangle []int, triangleEncoder *wav.Encoder) {
	format := audio.Format{
		NumChannels: 2,
		SampleRate:  44100,
	}

	createWaveformAudioFile(sine, format, *sinEncoder)
	createWaveformAudioFile(square, format, *squareEncoder)
	createWaveformAudioFile(triangle, format, *triangleEncoder)
}

func generateSinWave(seconds int) []int {
	sine := make([]int, 44100*seconds)
	for i := range sine {
		sine[i] = int(32767.0 * math.Sin(float64(i)*2*math.Pi*440/44100))
	}
	return sine
}

func generateSquareWave(seconds int) []int {
	square := make([]int, 44100*seconds)
	for i := range square {
		if i < 44100*seconds/2 {
			square[i] = int(32767.0 * math.Sin(float64(i)*2*math.Pi*440/44100))
		} else {
			square[i] = int(-32767.0 * math.Sin(float64(i)*2*math.Pi*440/44100))
		}
	}
	return square
}

func generateTriangleWave(seconds int) []int {
	triangle := make([]int, 44100*seconds)
	for i := range triangle {
		if i < 44100*seconds/4 {
			triangle[i] = int(32767.0 * (2.0 / float64(44100*seconds/4)) * float64(i))
		} else if i < 44100*seconds/2 {
			triangle[i] = int(32767.0 * (2.0 - (2.0/float64(44100*seconds/4))*float64(i)))
		} else if i < 44100*seconds*3/4 {
			triangle[i] = int(-32767.0 * (2.0 / float64(44100*seconds/4)) * float64(i))
		} else {
			triangle[i] = int(-32767.0 * (2.0 - (2.0/float64(44100*seconds/4))*float64(i)))
		}
	}
	return triangle
}

func createFileEncoder(filename string) *wav.Encoder {
	file, err := os.Create("output/" + filename)
	if err != nil {
		panic(err)
	}
	return wav.NewEncoder(file, 44100, 16, 1, 1)
}

func createWaveformAudioFile(waveform []int, format audio.Format, encoder wav.Encoder) {
	buffer := audio.IntBuffer{Data: waveform, Format: &format}
	writeAudioToFile(encoder, buffer)
}
func writeAudioToFile(encoder wav.Encoder, buffer audio.IntBuffer) {
	if err := encoder.Write(&buffer); err != nil {
		panic(err)
	}
	if err := encoder.Close(); err != nil {
		panic(err)
	}
}
