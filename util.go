package main

import "github.com/go-gl/gl"
import "fmt"

func createTexture(width int, height int, depth int, pixels []float32) (gl.Texture, error) {
    //gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.REPLACE)
    textureID := gl.GenTexture()
    textureID.Bind(gl.TEXTURE_3D)

    gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_S, gl.CLAMP)
    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_T, gl.CLAMP)
    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_R, gl.CLAMP)

    //pixels := make([]float32, width*height*depth)
    //for x := 0; x<width; x++ {
    //    for y := 0; y<height; y++ {
    //        for z := 0; z<depth; z++ {
    //            pixels[z*width*height + y*width + x] =
    //            //pixels[z*width*height + x*height + y] =
    //            //pixels[y*width*depth + z*width + x] =
    //            //pixels[y*width*depth + x*depth + z] =
    //            //pixels[x*height*depth + y*depth + z] =
    //            //pixels[x*height*depth + z*height + y] =
    //                data[z*width*height + y*width + x]
    //                //data[z*width*height + x*height + y]
    //                //data[y*width*depth + z*width + x]
    //                //data[y*width*depth + x*depth + z]
    //                //data[x*height*depth + y*depth + z]
    //                //data[x*height*depth + z*height + y]
    //        }
    //    }
    //}


    gl.TexImage3D(gl.TEXTURE_3D, 0, gl.R32F,
        width, height, depth, 0, gl.RED,
        gl.FLOAT, pixels)

    return textureID, nil
}

func readTexture3D(width int, height int, depth int, path string) (gl.Texture, error) {
    _, data := readVTKFile(path)
    return createTexture(width, height, depth, data)
}

func readTexture3DBinary(width int, height int, depth int, path string) (gl.Texture, error) {
    fmt.Printf("Test1\n")
    _, data := readVTKBinaryFile(path)
    fmt.Printf("Test2\n")
    return createTexture(width, height, depth, data)
}

func createProgram(vs string, fs string) (gl.Program, error) {
    vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
    vertex_shader.Source(vs)
    vertex_shader.Compile()
    if vertex_shader.Get(gl.COMPILE_STATUS) != gl.TRUE {
        panic("Vertex shader error: " + vertex_shader.GetInfoLog())
    }

    fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
    fragment_shader.Source(fs)
    fragment_shader.Compile()
    if fragment_shader.Get(gl.COMPILE_STATUS) != gl.TRUE {
        panic("Fragment shader error: " + fragment_shader.GetInfoLog())
    }

    program := gl.CreateProgram()
    program.AttachShader(vertex_shader)
    program.AttachShader(fragment_shader)
    program.Link()
    if program.Get(gl.LINK_STATUS) != gl.TRUE {
        panic("Program error: " + program.GetInfoLog())
    }

    program.Use()
    return program, nil
}
