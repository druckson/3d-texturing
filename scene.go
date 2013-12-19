package main

import ("github.com/go-gl/gl"
        glfw "github.com/go-gl/glfw3"
        "fmt")

var _vr *VolumetricRenderer

func ErrorCallback(err glfw.ErrorCode, desc string) {
    fmt.Printf("%v: %v\n", err, desc)
}

func SetWindowSize(window *glfw.Window, width int, height int) {
    _vr.SetSize(width, height)
}

type VolumetricRenderer struct {
    window *glfw.Window
    volumeData gl.Texture
    program gl.Program
    width gl.UniformLocation
    height gl.UniformLocation
    position gl.UniformLocation
    up gl.UniformLocation
    focus gl.UniformLocation
}

func CreateVolumetricRenderer() *VolumetricRenderer {
    var err error;
    vr := new(VolumetricRenderer)
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
    glfw.SwapInterval(1)
    gl.Init()


    gl.ClearColor(0.5, 0.5, 0.5, 0)

    gl.Viewport(0, 0, 100, 100)
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    dim := 64

    vr.volumeData, _ = readTexture3D(dim, dim, dim, "astro64.txt")
    vr.program, _ = createProgram(`
        void main() {
            gl_TexCoord[0] = gl_MultiTexCoord0;
            gl_Position = gl_ModelViewProjectionMatrix * gl_Vertex;
        }
    `, `
        uniform sampler1D transferFunction;
        uniform sampler3D volumeData;
        uniform vec3 position;
        uniform vec3 up;
        uniform vec3 focus;
        uniform float near;
        uniform float far;
        uniform float angle;
        uniform int width;
        uniform int height;

        vec3 lerp(vec3 p1, vec3 p2, float t) {
            return p1 + (p2 - p1) * t;
        }

        //vec3 getRay(int i, int j, int nx, int ny) {
        //    vec3 ru = normalize(-position + up)
        //    vec3 rv = normalize(-position + ru)
        //    vec3 rx = ru*(2*tan(angle*PI/360)/nx);
        //    vec3 ry = rv*(2*tan(angle*PI/360)/ny);

        //    return normalize(-position) + 
        //        (rx * (2*i + 1 - nx) / 2) +
        //        (ry * (2*j + 1 - ny) / 2);
        //}

        void main(void) {
            //vec3 start = texture2D(frontPlane, gl_TexCoord[0].st).xyz;
            //vec3 end   = start * back;

            //float t;
            //int samples = 100;
            //gl_FragColor = vec4(0, 0, 0, 255);
            //for (int i=0; i<samples; i++) {
            //    t = float(i) / float(samples);
            //    vec3 pos = lerp(start, end, t);
            //    float value = texture3D(volumeData, pos).r * float(1)/float(samples);
            //    value *= 255f;
            //    gl_FragColor += vec4(value, value, value, 255);
            //}

            //gl_FragColor += vec4(0.1, 0, 0, 255);
            gl_FragColor = vec4(gl_TexCoord[0].s, gl_TexCoord[0].t, 0, 255);
            //gl_FragColor += vec4(float(width)/100f, 0, 0, 255);
            //gl_FragColor = vec4(1f, 0, 0, 255);
        }
    `)

    //vr.SetSize(400, 400)

    vr.position = vr.program.GetUniformLocation("position")
    vr.position.Uniform3f(0.0, 0.0, 0.0)

    vr.up = vr.program.GetUniformLocation("up")
    vr.up.Uniform3f(0.0, 0.0, -1.0)

    vr.focus = vr.program.GetUniformLocation("focus")
    vr.focus.Uniform3f(-8.25e+7, -3.45e+7, 3.35e+7)

    return vr
}

func (vr *VolumetricRenderer) SetSize(width int, height int) {
    vr.width = vr.program.GetUniformLocation("width");
    vr.width.Uniform1i(width)
    vr.height = vr.program.GetUniformLocation("height");
    vr.height.Uniform1i(height)
    gl.Viewport(0, 0, width, height)
}

func (vr *VolumetricRenderer) Draw() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    vr.program.Use()

    gl.Begin(gl.QUADS)
    gl.TexCoord3f(0, 0, 0.5)
    gl.Vertex3f( -1,-1, 0)
    gl.TexCoord3f(0, 1, 0.5)
    gl.Vertex3f( -1, 1, 0)
    gl.TexCoord3f(1, 1, 0.5)
    gl.Vertex3f(  1, 1, 0)
    gl.TexCoord3f(1, 0, 0.5)
    gl.Vertex3f(  1,-1, 0)
    gl.End()

    vr.window.SwapBuffers()
    glfw.PollEvents()
}

func (vr *VolumetricRenderer) Destroy() {
    vr.volumeData.Delete()
    vr.window.Destroy()
    glfw.Terminate()
}
