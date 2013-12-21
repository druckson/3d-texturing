package main

import "github.com/go-gl/gl"

func (sf *ScalarField) CreateTexture() (gl.Texture, error) {
    textureID := gl.GenTexture()
    textureID.Bind(gl.TEXTURE_3D)

    gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_S, gl.CLAMP)
    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_T, gl.CLAMP)
    gl.TexParameterf(gl.TEXTURE_3D, gl.TEXTURE_WRAP_R, gl.CLAMP)

    gl.TexImage3D(gl.TEXTURE_3D, 0, gl.R32F,
        sf.width, sf.height, sf.depth, 0, gl.RED,
        gl.FLOAT, sf.data)

    return textureID, nil
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
