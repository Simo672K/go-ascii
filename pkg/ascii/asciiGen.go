package ascii

import (
	"github.com/Simo672K/go-ascii/pkg/process"
)

const AsciiChars = ".,-_:;^*=#&@"

func GenerateASCII(img process.Image) string {
	generatedAscii := ""
	density := 5

	for row := 0; row < img.Height; row += density {
		for col := 0; col < img.Width; col += density {
			avrg := 0
			for i := 0; i < density && (row+i) < img.Height; i++ {
				for j := 0; j < density && (col+j) < img.Width; j++ {
					avrg += img.GrayScaledPixelArray[row+i][col+j].R
				}
			}
			avrg /= density

			switch {
			case avrg > 220:
				generatedAscii += "@"
			case avrg > 190:
				generatedAscii += "#"
			case avrg > 170:
				generatedAscii += "&"
			case avrg > 140:
				generatedAscii += "="
			case avrg > 110:
				generatedAscii += "+"
			case avrg > 80:
				generatedAscii += ";"
			case avrg > 50:
				generatedAscii += "-"
			case avrg > 20:
				generatedAscii += "."
			default:
				generatedAscii += " "
			}
		}
		generatedAscii += "\n"
	}
	return generatedAscii
}
