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

    volumeData gl.UniformLocation
    width gl.UniformLocation
    height gl.UniformLocation
    samples gl.UniformLocation

    test gl.UniformLocation

    angle gl.UniformLocation
    near gl.UniformLocation
    far gl.UniformLocation

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
    gl.Init()

    gl.ClearColor(0.5, 0.5, 0.5, 0)

    //gl.Viewport(0, 0, 100, 100)
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

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
        uniform int samples;

        uniform float test;

        float sample(vec3 pos) {
            float offs = 3.02638 * pow(10.0, 7.0);
            float mult = offs * 2.0;
            pos = (pos + vec3(offs)) / mult;
            if ((pos.x > 0.0) && (pos.x < 1.0) &&
                (pos.y > 0.0) && (pos.y < 1.0) &&
                (pos.z > 0.0) && (pos.z < 1.0))
                {
                return texture3D(volumeData, pos).r;
            }
            return 0.0;
        }

        vec3 lerp(vec3 p1, vec3 p2, float t) {
            return p1 + ((p2 - p1) * t);
        }

        vec4 lerp4(vec4 p1, vec4 p2, float t) {
            return p1 + ((p2 - p1) * t);
        }

        float delerp(float v1, float v2, float v) {
            return (v - v1) / (v2 - v1);
        }

        vec4 transferFunction(float value, float dist) {
            float t = delerp(0, 15, value);
            return dist*lerp4(vec4(1, 0, 0, 0), vec4(0, 1, 0, 1), t);
        }

        vec4 mix(vec4 previousColor, vec4 newColor) {
            
            return previousColor + newColor;
        }

        vec3 getRay(float i, float j, int nx, int ny) {
            float pi = 3.1415926535897932384626433832795;
            vec3 ru = normalize(-position * up);
            vec3 rv = normalize(-position * ru);
            vec3 rx = ru*(2.0*tan(angle*pi/360.0)/float(width));
            vec3 ry = rv*(2.0*tan(angle*pi/360.0)/float(height));

            return normalize(-position) + 
                (rx * ((2.0*i + 1.0 - float(width)) / 2.0)) +
                (ry * ((2.0*j + 1.0 - float(height)) / 2.0));
        }

        vec4 sampleRay(vec3 ray) {
            vec4 color = vec4(0);
            float len = length(ray) * (far - near);
            for (int i=0; i<samples; i++) {
                vec3 point = lerp(ray*near, ray*far, float(i)/float(samples));
                float value = sample(point);
                color += transferFunction(value, 0.00000001*len/float(samples));
            }
            return color;
        }

        void main(void) {
            vec3 ray = getRay(gl_TexCoord[0].s*float(width),
                              gl_TexCoord[0].t*float(height),
                              width, height);

            gl_FragColor = sampleRay(ray);
            //gl_FragColor.rgb = ray;

            //gl_FragColor = vec4(0.0);
            //for (float t = 0.0; t<1.0; t+=0.01) {
            //    float value = texture3D(volumeData, vec3(t, gl_TexCoord[0].t, gl_TexCoord[0].s)).r;

            //    
            //    vec4 color = transferFunction(value, 1.0/float(width));
            //    gl_FragColor = mix(gl_FragColor, color);
            //}

            //vec4 color = texture3D(volumeData, vec3(test, gl_TexCoord[0].t, gl_TexCoord[0].s));
            //gl_FragColor = color * 0.1;
        }
    `)


    //dim := 512
    //vr.volumeDataTexture, _ = readTexture3DBinary(dim, dim, dim, "astro512.txt")
    dim := 64
    vr.volumeDataTexture, _ = readTexture3D(dim, dim, dim, "astro64.txt")

    vr.volumeData = vr.program.GetUniformLocation("volumeData")
    vr.volumeData.Uniform1i(0)

    vr.position = vr.program.GetUniformLocation("position")
    vr.position.Uniform3f(-8.25e+7, -3.45e+7, 3.35e+7)

    vr.up = vr.program.GetUniformLocation("up")
    vr.up.Uniform3f(0.0, 0.0, -1.0)

    vr.focus = vr.program.GetUniformLocation("focus")
    vr.focus.Uniform3f(0.0, 0.0, 0.0)

    vr.angle = vr.program.GetUniformLocation("angle")
    vr.angle.Uniform1f(30.0)

    vr.near = vr.program.GetUniformLocation("near")
    vr.near.Uniform1f(7.5e+7)

    vr.far = vr.program.GetUniformLocation("far")
    vr.far.Uniform1f(1.4e+8)

    vr.samples = vr.program.GetUniformLocation("samples")
    vr.samples.Uniform1i(100)

    return vr
}

func (vr *VolumetricRenderer) SetSize(width int, height int) {
    vr.width = vr.program.GetUniformLocation("width");
    vr.width.Uniform1i(width)
    vr.height = vr.program.GetUniformLocation("height");
    vr.height.Uniform1i(height)
    gl.Viewport(0, 0, width, height)
}

var t float64 = 0.0;

func (vr *VolumetricRenderer) Draw() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()

    t += 0.01;
    fmt.Printf("Test %v\n", math.Sin(t))
    vr.test = vr.program.GetUniformLocation("test")
    vr.test.Uniform1f(float32(math.Sin(t) / 2.0 + 0.5))

    gl.ActiveTexture(gl.TEXTURE0)

    vr.volumeData = vr.program.GetUniformLocation("volumeData")
    vr.volumeData.Uniform1i(0)

    vr.volumeDataTexture.Bind(gl.TEXTURE_3D)
    gl.Enable(gl.TEXTURE_3D)

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
