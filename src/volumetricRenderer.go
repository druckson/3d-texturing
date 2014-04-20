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

type VolumetricRenderer struct {
    window *glfw.Window
    program gl.Program

    volumeDataTexture gl.Texture
    transferFunctionTexture gl.Texture
    scalarField *ScalarField

    conf *Configuration

    perspectiveX float64
    perspectiveY float64
}

func CreateVolumetricRenderer(sf *ScalarField, conf *Configuration) *VolumetricRenderer {
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

    vr.conf = conf
    vr.volumeDataTexture, _ = sf.CreateTexture()

    vr.Init()

    return vr
}

func (vr *VolumetricRenderer) Init() {
    volumeData := vr.program.GetUniformLocation("volumeData")
    volumeData.Uniform1i(0)

    vr.transferFunctionTexture = SetupTransferFunction()

    transferFunctionMin := vr.program.GetUniformLocation("transferFunctionMin")
    transferFunctionMin.Uniform1f(vr.conf.Min)

    transferFunctionMax := vr.program.GetUniformLocation("transferFunctionMax")
    transferFunctionMax.Uniform1f(vr.conf.Max)

    up := vr.program.GetUniformLocation("up")
    up.Uniform3f(0.0, 1.0, 0.0)

    focus := vr.program.GetUniformLocation("focus")
    focus.Uniform3f(0.0, 0.0, 0.0)

    size := vr.program.GetUniformLocation("size")
    size.Uniform3f(vr.conf.Width, vr.conf.Height, vr.conf.Depth)

    angle := vr.program.GetUniformLocation("angle")
    angle.Uniform1f(30.0)

    near := vr.program.GetUniformLocation("near")
    near.Uniform1f(vr.conf.Near)

    far := vr.program.GetUniformLocation("far")
    far.Uniform1f(vr.conf.Far)

    samples := vr.program.GetUniformLocation("samples")
    samples.Uniform1i(vr.conf.Samples)
}

func (vr *VolumetricRenderer) SetSize(w int, h int) {
    width := vr.program.GetUniformLocation("screenWidth");
    width.Uniform1i(w)
    height := vr.program.GetUniformLocation("screenHeight");
    height.Uniform1i(h)
    gl.Viewport(0, 0, w, h)
}

func (vr *VolumetricRenderer) Rotate(amount float64) {
    vr.perspectiveX += amount
}

func (vr *VolumetricRenderer) Zoom(amount float64) {
    vr.perspectiveY += amount
}

func (vr *VolumetricRenderer) Update(dt float64) {
    rotateSpeed := dt * 5.0
    zoomSpeed := dt * 100.0
    if  vr.window.GetKey(glfw.KeyLeft) == glfw.Press ||
        vr.window.GetKey(glfw.KeyLeft) == glfw.Repeat {
        vr.Rotate(-rotateSpeed)
    }
    if  vr.window.GetKey(glfw.KeyRight) == glfw.Press ||
        vr.window.GetKey(glfw.KeyRight) == glfw.Repeat {
        vr.Rotate(rotateSpeed)
    }
    if  vr.window.GetKey(glfw.KeyUp) == glfw.Press ||
        vr.window.GetKey(glfw.KeyUp) == glfw.Repeat {
        vr.Zoom(-zoomSpeed)
    }
    if  vr.window.GetKey(glfw.KeyDown) == glfw.Press ||
        vr.window.GetKey(glfw.KeyDown) == glfw.Repeat {
        vr.Zoom(zoomSpeed)
    }
}

func (vr *VolumetricRenderer) Draw() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    test := vr.program.GetUniformLocation("test")
    test.Uniform1f(float32(math.Sin(vr.perspectiveX) + 1.0) * 0.5)
    position := vr.program.GetUniformLocation("position")
    position.Uniform3f((vr.conf.Dist+float32(vr.perspectiveY))*float32(math.Sin(vr.perspectiveX)), 0,
                       (vr.conf.Dist+float32(vr.perspectiveY))*float32(math.Cos(vr.perspectiveX)))

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
