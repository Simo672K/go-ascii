package ascii

import (
	"github.com/Simo672K/go-ascii/pkg/process"
)

// const AsciiChars = ".,-_:;^*=#&@"

func GenerateASCII(img process.Image, density int) string {
	generatedAscii := ""

	for row := 0; row < img.Height; row += density {
		for col := 0; col < img.Width; col += density {
			avrg := CalculateAvrg(density, row, col, img)

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
			case avrg > 30:
				generatedAscii += "¨"
			case avrg > 10:
				generatedAscii += "."
			default:
				generatedAscii += " "
			}
		}
		generatedAscii += "\n"
	}
	return generatedAscii
}
