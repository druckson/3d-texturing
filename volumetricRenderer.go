package main

import ("github.com/go-gl/gl"
        glfw "github.com/go-gl/glfw3"
        "fmt"
        "math")

var _vr *VolumetricRenderer

func ErrorCallback(err glfw.ErrorCode, desc string) {
    fmt.Printf("%v: %v\n", err, desc)
}

func SetWindowSize(window *glfw.Window, width int, height int) {
    _vr.SetSize(width, height)
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

        colors[4*i+0] = float32(controlPointColors[3*start+0])/255.0 + (t * (float32(controlPointColors[3*end+0])/255.0 - float32(controlPointColors[3*start+0])/255.0))
        colors[4*i+1] = float32(controlPointColors[3*start+1])/255.0 + (t * (float32(controlPointColors[3*end+1])/255.0 - float32(controlPointColors[3*start+1])/255.0))
        colors[4*i+2] = float32(controlPointColors[3*start+2])/255.0 + (t * (float32(controlPointColors[3*end+2])/255.0 - float32(controlPointColors[3*start+2])/255.0))
        colors[4*i+3] = float32(opacities[i])/255.0
    }

    gl.TexImage1D(gl.TEXTURE_1D, 0, gl.RGBA, numBins, 0, gl.RGBA, gl.FLOAT, colors)

    return texture
}

type VolumetricRenderer struct {
    window *glfw.Window
    program gl.Program

    volumeDataTexture gl.Texture
    transferFunctionTexture gl.Texture
    scalarField *ScalarField

    samples int
    width float32
    height float32
    depth float32
    dist float32
    near float32
    far float32
    min float32
    max float32
}

func CreateVolumetricRenderer(sf *ScalarField, min float32, max float32, samples int,
                              width float32, height float32, depth float32,
                              dist float32, near float32, far float32) *VolumetricRenderer {
    var err error;
    vr := new(VolumetricRenderer)
    vr.scalarField = sf
    _vr = vr
    glfw.SetErrorCallback(ErrorCallback)

    if !glfw.Init() {
        panic("Can't init glfw!")
    }

    vr.window, err = glfw.CreateWindow(300, 300, "Test", nil, nil)
    if err != nil {
        panic(err)
    }

    vr.window.SetSizeCallback(SetWindowSize)
    vr.window.MakeContextCurrent()
    gl.Init()

    vr.program, _ = createProgram(`
        void main() {
            gl_TexCoord[0] = gl_MultiTexCoord0;
            gl_Position = gl_ModelViewProjectionMatrix * gl_Vertex;
        }
    `, string(readFile("volumeShader.frag")))

    vr.samples = samples
    vr.width = width
    vr.height = height
    vr.depth = depth
    vr.dist = dist
    vr.near = near
    vr.far = far
    vr.min = min
    vr.max = max

    vr.volumeDataTexture, _ = sf.CreateTexture()

    vr.Init()

    return vr
}

func (vr *VolumetricRenderer) Init() {
    volumeData := vr.program.GetUniformLocation("volumeData")
    volumeData.Uniform1i(0)

    vr.transferFunctionTexture = SetupTransferFunction()

    transferFunctionMin := vr.program.GetUniformLocation("transferFunctionMin")
    transferFunctionMin.Uniform1f(vr.min)

    transferFunctionMax := vr.program.GetUniformLocation("transferFunctionMax")
    transferFunctionMax.Uniform1f(vr.max)

    up := vr.program.GetUniformLocation("up")
    up.Uniform3f(0.0, 1.0, 0.0)

    focus := vr.program.GetUniformLocation("focus")
    focus.Uniform3f(0.0, 0.0, 0.0)

    size := vr.program.GetUniformLocation("size")
    size.Uniform3f(vr.width, vr.height, vr.depth)

    angle := vr.program.GetUniformLocation("angle")
    angle.Uniform1f(30.0)

    near := vr.program.GetUniformLocation("near")
    near.Uniform1f(vr.near)

    far := vr.program.GetUniformLocation("far")
    far.Uniform1f(vr.far)

    samples := vr.program.GetUniformLocation("samples")
    samples.Uniform1i(vr.samples)
}

func (vr *VolumetricRenderer) SetSize(w int, h int) {
    width := vr.program.GetUniformLocation("screenWidth");
    width.Uniform1i(w)
    height := vr.program.GetUniformLocation("screenHeight");
    height.Uniform1i(h)
    gl.Viewport(0, 0, w, h)
}

var t float64 = 0.0;

func (vr *VolumetricRenderer) Draw() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    t += 0.01
    test := vr.program.GetUniformLocation("test")
    test.Uniform1f(float32(math.Sin(t) + 1.0) * 0.5)
    position := vr.program.GetUniformLocation("position")
    position.Uniform3f(vr.dist*float32(math.Sin(t)), 0, vr.dist*float32(math.Cos(t)))

    gl.ActiveTexture(gl.TEXTURE0)

    volumeData := vr.program.GetUniformLocation("volumeData")
    volumeData.Uniform1i(0)
    vr.volumeDataTexture.Bind(gl.TEXTURE_3D)
    gl.Enable(gl.TEXTURE_3D)

    gl.ActiveTexture(gl.TEXTURE1)
    transferFunction := vr.program.GetUniformLocation("transferFunction")
    transferFunction.Uniform1i(1)
    vr.transferFunctionTexture.Bind(gl.TEXTURE_1D)
    gl.Enable(gl.TEXTURE_1D)

    vr.program.Use()

    gl.Begin(gl.QUADS)
    gl.TexCoord2f(0, 0)
    gl.Vertex2f( -1,-1)
    gl.TexCoord2f(0, 1)
    gl.Vertex2f( -1, 1)
    gl.TexCoord2f(1, 1)
    gl.Vertex2f(  1, 1)
    gl.TexCoord2f(1, 0)
    gl.Vertex2f(  1,-1)
    gl.End()

    vr.window.SwapBuffers()
    glfw.PollEvents()
}

func (vr *VolumetricRenderer) Destroy() {
    vr.volumeDataTexture.Delete()
    vr.window.Destroy()
    glfw.Terminate()
}
