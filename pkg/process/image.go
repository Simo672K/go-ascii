package process

import (
	"image"
	_ "image/png"
	"io"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

type Image struct {
	ImgReader            io.Reader
	NumberOfPixels       int
	PixelArray           [][]Pixel
	GrayScaledPixelArray [][]Pixel
}

func (img *Image) GetPixelArray() error {
	var pixels [][]Pixel

	decodedImg, _, err := image.Decode(img.ImgReader)
	if err != nil {
		return err
	}

	bounds := decodedImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	img.NumberOfPixels = width * height

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			pixelsData := *RGBA2Pixel(decodedImg.At(x, y).RGBA())
			row = append(row, pixelsData)
		}
		pixels = append(pixels, row)
	}

	img.PixelArray = pixels
	return nil
}

func (img *Image) ToGrayScale() {
	for _, row := range img.PixelArray {
		var gsRow []Pixel
		for _, pixel := range row {
			grayPixel := Pixel{R: CalculatePixelAvrgRGB(pixel), G: CalculatePixelAvrgRGB(pixel), B: CalculatePixelAvrgRGB(pixel), A: pixel.A}
			gsRow = append(gsRow, grayPixel)
		}
		img.GrayScaledPixelArray = append(img.GrayScaledPixelArray, gsRow)
	}
}
