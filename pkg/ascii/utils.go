package ascii

import "github.com/Simo672K/go-ascii/pkg/process"

func CalculateAvrg(density, row, col int, img process.Image) int {
	avrg := 0
	for i := 0; i < density && (row+i) < img.Height; i++ {
		for j := 0; j < density && (col+j) < img.Width; j++ {
			avrg += img.GrayScaledPixelArray[row+i][col+j].R
		}
	}

	return avrg / density
}
