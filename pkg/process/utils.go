package process

func RGBA2Pixel(r, g, b, a uint32) *Pixel {
	return &Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257), A: int(a / 257)}
}

func CalculatePixelAvrgRGB(pixel Pixel) int {
	return (pixel.R + pixel.G + pixel.B + pixel.A) / 4
}
