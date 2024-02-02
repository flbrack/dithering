package main

import (
	"image"
	"image/color"
	"math/rand"
)

func floydSteinberg(grey *image.Gray) *image.Gray {
    bounds := grey.Bounds()
    width := bounds.Dx()
    height := bounds.Dy()
    output := copyGrey(grey)

    for j := 0; j < height; j++ {
	for i := 0; i < width; i++ {
	    oldPixel := output.GrayAt(i, j)
	    newPixel := blackOrWhite(oldPixel, 127)
	    output.SetGray(i, j, newPixel)

	    quantErr := ( int16(oldPixel.Y) - int16(newPixel.Y) )/ 16

	    output.SetGray(i+1, j, color.Gray{clamp(int16(output.GrayAt(i+1, j).Y) + quantErr*7)})
	    output.SetGray(i-1, j+1, color.Gray{clamp(int16(output.GrayAt(i-1, j+1).Y) + quantErr*3)})
	    output.SetGray(i, j+1, color.Gray{clamp(int16(output.GrayAt(i, j+1).Y) + quantErr*5)})
	    output.SetGray(i+1, j+1, color.Gray{clamp(int16(output.GrayAt(i+1, j+1).Y) + quantErr*1)})
	}
    }
    return output
}

func colourFloydSteinberg(img *image.RGBA) *image.RGBA {
    bounds := img.Bounds()
    width := bounds.Dx()
    height := bounds.Dy()
    output := copyRGBA(img)

    for j := 0; j < height; j++ {
	for i := 0; i < width; i++ {
	    oldPixel := output.RGBAAt(i, j)
	    newPixel := newColourPixel(oldPixel)
	    output.SetRGBA(i, j, newPixel)

	    redQuantErr := ( int16(oldPixel.R) - int16(newPixel.R) )/ 16
	    greenQuantErr := ( int16(oldPixel.G) - int16(newPixel.G) )/ 16
	    blueQuantErr := ( int16(oldPixel.B) - int16(newPixel.B) )/ 16

	    output.SetRGBA(i+1, j, color.RGBA{
		clamp(int16(output.RGBAAt(i+1, j).R) + redQuantErr*7),
		clamp(int16(output.RGBAAt(i+1, j).G) + greenQuantErr*7),
		clamp(int16(output.RGBAAt(i+1, j).B) + blueQuantErr*7),
		255,
	    })
	    output.SetRGBA(i-1, j+1, color.RGBA{
		clamp(int16(output.RGBAAt(i-1, j+1).R) + redQuantErr*3),
		clamp(int16(output.RGBAAt(i-1, j+1).G) + greenQuantErr*3),
		clamp(int16(output.RGBAAt(i-1, j+1).B) + blueQuantErr*3),
		255,
	    })
	    output.SetRGBA(i, j+1, color.RGBA{
		clamp(int16(output.RGBAAt(i, j+1).R) + redQuantErr*5),
		clamp(int16(output.RGBAAt(i, j+1).G) + greenQuantErr*5),
		clamp(int16(output.RGBAAt(i, j+1).B) + blueQuantErr*5),
		255,
	    })
	    output.SetRGBA(i+1, j+1, color.RGBA{
		clamp(int16(output.RGBAAt(i+1, j+1).R) + redQuantErr*1),
		clamp(int16(output.RGBAAt(i+1, j+1).G) + greenQuantErr*1),
		clamp(int16(output.RGBAAt(i+1, j+1).B) + blueQuantErr*1),
		255,
	    })
	}
    }
    return output
}

func thresholdDither(grey *image.Gray) *image.Gray {
    output := image.NewGray(grey.Bounds())
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    ditheredColor := blackOrWhite(grey.GrayAt(x, y), 127)
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}


func thresholdDitherColor(img *image.RGBA) *image.RGBA {
    output := image.NewRGBA(img.Bounds())
    for x := 0; x < img.Bounds().Dx(); x++ {
	for y := 0; y < img.Bounds().Dy(); y++ {
	    oldColor := img.RGBAAt(x, y)
	    newColor := newColourPixel(oldColor)
	    output.SetRGBA(x, y, newColor)
	}
    }
    return output
}

func randomNoiseDither(grey *image.Gray) *image.Gray {
    output := image.NewGray(grey.Bounds())
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    ditheredColor := blackOrWhite(grey.GrayAt(x, y), uint8(rand.Intn(255)))
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}

func bayerDither(grey *image.Gray) *image.Gray {
    bayerMatrix := [4][4]float32{
	{0, 8, 2, 10},
	{12, 4, 14, 6},
	{3, 11, 1, 9},
	{15, 7, 13, 5},
    }
    output := copyGrey(grey)
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    oldPixel := output.GrayAt(x, y)
	    var ditheredColor color.Gray
	    if oldPixel.Y > uint8((255/16) * bayerMatrix[x%4][y%4]) {
		ditheredColor = color.Gray{255}
	    } else {
		ditheredColor = color.Gray{0}
	    }
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}

func bayerDither0(grey *image.Gray) *image.Gray {
    bayerMatrix := [2][2]float32{
	{0, 2},
	{3, 1},
    }
    output := copyGrey(grey)
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    oldPixel := output.GrayAt(x, y)
	    var ditheredColor color.Gray
	    if oldPixel.Y > uint8((255/4) * bayerMatrix[x%2][y%2]) {
		ditheredColor = color.Gray{255}
	    } else {
		ditheredColor = color.Gray{0}
	    }
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}

func halftoneDither(grey *image.Gray) *image.Gray {
    bayerMatrix := [4][4]float32{
	{12, 5, 6, 13},
	{4, 0, 1, 7},
	{11, 3, 2, 8},
	{15, 10, 9, 14},
    }
    output := copyGrey(grey)
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    oldPixel := output.GrayAt(x, y)
	    var ditheredColor color.Gray
	    if oldPixel.Y > uint8((255/15) * bayerMatrix[x%4][y%4]) {
		ditheredColor = color.Gray{255}
	    } else {
		ditheredColor = color.Gray{0}
	    }
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}

func halftoneDither2(grey *image.Gray) *image.Gray {
    bayerMatrix := [8][8]float32{
	 {0.567, 0.635, 0.608, 0.514, 0.424, 0.365, 0.392, 0.486},
	 {0.847, 0.878, 0.910, 0.698, 0.153, 0.122, 0.090, 0.302},
	 {0.820, 0.969, 0.941, 0.667, 0.180, 0.031, 0.059, 0.333},
	 {0.725, 0.788, 0.757, 0.545, 0.275, 0.212, 0.243, 0.455},
	 {0.424, 0.365, 0.392, 0.486, 0.567, 0.635, 0.608, 0.514},
	 {0.153, 0.122, 0.090, 0.302, 0.847, 0.878, 0.910, 0.698},
	 {0.180, 0.031, 0.059, 0.333, 0.820, 0.969, 0.941, 0.667},
	 {0.275, 0.212, 0.243, 0.455, 0.725, 0.788, 0.757, 0.545},
    }
    output := copyGrey(grey)
    for x := 0; x < grey.Bounds().Dx(); x++ {
	for y := 0; y < grey.Bounds().Dy(); y++ {
	    oldPixel := output.GrayAt(x, y)
	    var ditheredColor color.Gray
	    if oldPixel.Y > uint8(255 * bayerMatrix[x%8][y%8]) {
		ditheredColor = color.Gray{255}
	    } else {
		ditheredColor = color.Gray{0}
	    }
	    output.SetGray(x, y, ditheredColor)
	}
    }
    return output
}

