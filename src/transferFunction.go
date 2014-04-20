package main

import ("github.com/go-gl/gl")

func lerp(value1 float32, value2 float32, t float32) float32 {
    return value1 + (t * (value2 - value1));
}

func SetupTransferFunction() (gl.Texture) {
    texture := gl.GenTexture()
    texture.Bind(gl.TEXTURE_1D)

    gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    gl.TexParameterf(gl.TEXTURE_1D, gl.TEXTURE_WRAP_S, gl.CLAMP)

    numBins := 256

    colors := make([]float32, 4*numBins)

    opacities := [256]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 13, 14, 14, 14, 14, 14, 14, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 5, 4, 3, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 17, 17, 17, 17, 17, 17, 16, 16, 15, 14, 13, 12, 11, 9, 8, 7, 6, 5, 5, 4, 3, 3, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 16, 18, 20, 22, 24, 27, 29, 32, 35, 38, 41, 44, 47, 50, 52, 55, 58, 60, 62, 64, 66, 67, 68, 69, 70, 70, 70, 69, 68, 67, 66, 64, 62, 60, 58, 55, 52, 50, 47, 44, 41, 38, 35, 32, 29, 27, 24, 22, 20, 20, 23, 28, 33, 38, 45, 51, 59, 67, 76, 85, 95, 105, 116, 127, 138, 149, 160, 170, 180, 189, 198, 205, 212, 217, 221, 223, 224, 224, 222, 219, 214, 208, 201, 193, 184, 174, 164, 153, 142, 131, 120, 109, 99, 89, 79, 70, 62, 54, 47, 40, 35, 30, 25, 21, 17, 14, 12, 10, 8, 6, 5, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    controlPointColors := []int{ 71, 71, 219, 0, 0, 91, 0, 255, 255, 0, 127, 0, 255, 255, 0, 255, 96, 0, 107, 0, 0, 224, 76, 76 }
    controlPointPositions := []float32{ 0.0, 0.143, 0.285, 0.429, 0.571, 0.714, 0.857, 1.0 }

    for i:=0; i<numBins; i++ {
        end := 0
        for ; int(256.0*controlPointPositions[end])<=i; end++ {}
        start := end-1
        t := (float32(i) - 255.0*float32(controlPointPositions[start])) /
                          (255.0*float32(controlPointPositions[end]) -
                           255.0*float32(controlPointPositions[start]))

        colors[4*i+0] = lerp(float32(controlPointColors[3*start+0])/255.0, float32(controlPointColors[3*end+0])/255.0, t)
        colors[4*i+1] = lerp(float32(controlPointColors[3*start+1])/255.0, float32(controlPointColors[3*end+1])/255.0, t)
        colors[4*i+2] = lerp(float32(controlPointColors[3*start+2])/255.0, float32(controlPointColors[3*end+2])/255.0, t)

        colors[4*i+3] = float32(opacities[i])/255.0
    }

    gl.TexImage1D(gl.TEXTURE_1D, 0, gl.RGBA, numBins, 0, gl.RGBA, gl.FLOAT, colors)

    return texture
}
