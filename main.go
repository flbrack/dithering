package main

func main() {

    bikeImage, _ := loadImage("pics/bike.png")

    greyImage := rgbaToGreyScale(bikeImage)

    thresholdImage := thresholdDither(greyImage)
    randomNoiseImage := randomNoiseDither(greyImage)
    fsDitherImage := floydSteinberg(greyImage)
    bayerImage := bayerDither(greyImage)
    bayerImage0 := bayerDither0(greyImage)
    halftoneImage := halftoneDither(greyImage)
    halftoneImage2 := halftoneDither2(greyImage)

    createAndSaveImage(greyImage, "pics/grey_bike.png")
    createAndSaveImage(thresholdImage, "pics/threshold_bike.png")
    createAndSaveImage(randomNoiseImage, "pics/randomnoise_bike.png")
    createAndSaveImage(fsDitherImage,  "pics/floydsteinberg_bike.png")
    createAndSaveImage(bayerImage,  "pics/bayer_bike.png")
    createAndSaveImage(bayerImage0,  "pics/bayer0_bike.png")
    createAndSaveImage(halftoneImage,  "pics/halftone_bike.png")
    createAndSaveImage(halftoneImage2,  "pics/halftone2_bike.png")

    rgbaImage := properFormatImage(bikeImage) 

    colourThreshold := thresholdDitherColor(rgbaImage)
    colourFS := colourFloydSteinberg(rgbaImage)

    createAndSaveImage(colourThreshold,  "pics/colourThreshold_bike.png")
    createAndSaveImage(colourFS,  "pics/colourFS_bike.png")

}

