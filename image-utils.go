package main

import (
    "os"
    "image"
    "image/png"
    "image/color"
)

func loadImage(path string) (image.Image, error) {
    imageFile, err := os.Open(path)
    if err != nil {
	return nil, err
    }
    defer imageFile.Close()
    img, _, err := image.Decode(imageFile)
    if err != nil {
	return nil, err
    }
    return img, nil
}

func createAndSaveImage(img image.Image, path string) error {
    file, err := os.Create(path)
    if err != nil {
	return err
    }
    defer file.Close()
    png.Encode(file, img)
    return nil
}

func copyGrey(grey *image.Gray) *image.Gray {
    output := image.NewGray(grey.Bounds())
    copy(output.Pix, grey.Pix)
    return output
}

func copyRGBA(img *image.RGBA) *image.RGBA {
    output := image.NewRGBA(img.Bounds())
    copy(output.Pix, img.Pix)
    return output
}

func rgbaToGreyScale(img image.Image) *image.Gray {
    output := image.NewGray(img.Bounds()) 
    for x := 0; x < img.Bounds().Max.X; x++ {
	for y := 0; y < img.Bounds().Max.Y; y++ {
	    rgba := img.At(x, y)
	    output.Set(x, y, rgba)
	}
    }
    return output
}

func properFormatImage(img image.Image) *image.RGBA {
    output := image.NewRGBA(img.Bounds()) 
    for x := 0; x < img.Bounds().Max.X; x++ {
	for y := 0; y < img.Bounds().Max.Y; y++ {
	    rgba := img.At(x, y)
	    output.Set(x, y, rgba)
	}
    }
    return output
}

func blackOrWhite(grey color.Gray, threshold uint8) color.Gray {
    if grey.Y > threshold {
	return color.Gray{255}
    } else {
	return color.Gray{0}
    }
}

func clamp(x int16) uint8 {
    if x < 1 {
	return uint8(0)
    } else if x > 254 {
	return uint8(255)
    } else {
	return uint8(x)
    }
}

func newColourPixel(oldColor color.RGBA) color.RGBA {
    var newColor color.RGBA
    newColor.A = 255
    if oldColor.R > 127 {
	newColor.R = 255
    } else {
	newColor.R = 0
    }
    if oldColor.G > 127 {
	newColor.G = 255
    } else {
	newColor.G = 0
    }
    if oldColor.B > 127 {
	newColor.B = 255
    } else {
	newColor.B = 0
    }
    return newColor
}
