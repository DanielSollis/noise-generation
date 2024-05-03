package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func generateSin(duration, sampleRate, frequency int) []float64 {
	sampleNumber := duration * sampleRate
	tau := math.Pi * 2
	angle := tau / float64(sampleRate)

	var sinWave []float64
	for i := 0; i < sampleNumber; i++ {
		xRadian := angle * float64(frequency) * float64(i)
		sinePoint := math.Sin(xRadian)
		sinWave = append(sinWave, sinePoint)
	}
	return sinWave
}

func writeSin(fileName string, sinWave []float64) (err error) {
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		return err
	}

	for i := 0; i < len(sinWave); i++ {
		var buffer [8]byte
		sample := float32(sinWave[i])
		binaryFloat := math.Float32bits(sample)
		binary.LittleEndian.PutUint32(buffer[:], binaryFloat)

		if _, err = file.Write(buffer[:]); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	duration := 2       // in seconds
	sampleRate := 44100 // CD quality sample rate
	frequency := 329    // A sub 4

	sinWave := generateSin(duration, sampleRate, frequency)

	err := writeSin("out.bin", sinWave)
	if err != nil {
		fmt.Println("Error writting file: ", err)
	}
}
