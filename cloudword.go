package main

import (
	"fmt"
	"math"
	"os"

	"math/rand"
)

func getColor(size, minSize, maxSize int) string {
	scale := float64(size-minSize) / float64(maxSize-minSize)
	r := int(255 * scale)
	g := int(100 * (1 - scale))
	b := int(255 * (1 - scale) * 0.7)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func checkCollision(x, y, size int, positions []struct{ x, y, size int }) bool {
	for _, pos := range positions {
		dist := math.Sqrt(float64((x-pos.x)*(x-pos.x) + (y-pos.y)*(y-pos.y)))
		if dist < float64(size+pos.size)/2 {
			return true
		}
	}
	return false
}

func getRandomPosition(size, width, height int, positions []struct{ x, y, size int }) (int, int) {
	x, y := rand.Intn(width-size*2), rand.Intn(height-size*2)
	for checkCollision(x, y, size, positions) {
		x, y = rand.Intn(width-size*2), rand.Intn(height-size*2)
	}
	return x, y
}

func getMinMaxSize(words map[string]int) (int, int) {
	minSize, maxSize := 100000, 0
	for _, size := range words {
		if size < minSize {
			minSize = size
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return minSize, maxSize
}

func writeSVGSart(file *os.File, width, height int) {
	file.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">`, width, height))
}

func writeSVGEnd(file *os.File) {
	file.WriteString("</svg>")
}

func generateCloudWords(file *os.File, words map[string]int, width, height int) {
	writeSVGSart(file, width, height)

	var positions []struct{ x, y, size int }

	for word, size := range words {
		x, y := getRandomPosition(size, width, height, positions)

		minSize, maxSize := getMinMaxSize(words)

		color := getColor(size, minSize, maxSize)
		positions = append(positions, struct{ x, y, size int }{x, y, size})

		file.WriteString(fmt.Sprintf(
			`<text x="%d" y="%d" font-size="%d" fill="%s" style="font-family: Arial, sans-serif;">%s</text>`,
			x, y, size, color, word))
	}
	writeSVGEnd(file)
}
