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
    dist float32
}

func CreateVolumetricRenderer(sf *ScalarField, min float32, max float32, sampleCount int, dist float32) *VolumetricRenderer {
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
    `, `
        uniform sampler1D transferFunction;
        uniform float transferFunctionMin;
        uniform float transferFunctionMax;

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
            if ((pos.r > 0.0) && (pos.r < 1.0) &&
                (pos.g > 0.0) && (pos.g < 1.0) &&
                (pos.b > 0.0) && (pos.b < 1.0))
                {
                return texture3D(volumeData, pos).r;
            }
            return 0.0;
        }

        vec3 lerp3(vec3 p1, vec3 p2, float t) {
            return p1 + ((p2 - p1) * t);
        }

        vec4 lerp4(vec4 p1, vec4 p2, float t) {
            return p1 + ((p2 - p1) * t);
        }

        float delerp(float v1, float v2, float v) {
            v = max(v1, min(v2, v));
            return (v - v1) / (v2 - v1);
        }

        vec3 getRay(float i, float j, int nx, int ny) {
            float pi = 3.1415926535897932384626433832795;
            vec3 ru = normalize(cross(-position, up));
            vec3 rv = normalize(cross(-position, ru));

            float hangle = angle;
            float wangle = angle * float(width) / float(height);

            vec3 rx = ru*(2.0*tan(wangle*pi/360.0)/float(width));
            vec3 ry = rv*(2.0*tan(hangle*pi/360.0)/float(height));

            return normalize(-position) + 
                (rx * ((2.0*i + 1.0 - float(width)) / 2.0)) +
                (ry * ((2.0*j + 1.0 - float(height)) / 2.0));
        }

        vec4 transferFunction(float value) {
            return texture1D(transferFunction, delerp(transferFunctionMin, transferFunctionMax, value));
        }

        vec4 mix(vec4 color, vec4 sample, float dist) {
            sample.a = sample.a * dist;
            vec4 newColor = color * (1.0 - sample.a) + sample * sample.a;
            newColor.a = color.a + (1.0 - color.a) * sample.a;
            return newColor;
        }

        float alphaFunction(float value) {
            return texture1D(transferFunction, delerp(transferFunctionMin, transferFunctionMax, value)).a;
        }

        float mix(float alpha, float sample, float dist) {
            return alpha + (1.0 - alpha) * sample;
        }

        vec4 sampleRay(vec3 ray) {
            float len = length(ray) * (far - near);
            float thickness = 0.000001*len/float(samples);

            int start = 0;
            int end = samples;

            // Cull unnecessary samples
            //float alpha = 0.0;
            //for (int i=0; i<samples; i++) {
            //    vec3 point = position + lerp3(ray*near, ray*far, float(i)/float(samples));
            //    float value = sample(point);
            //    alpha = mix(alpha, alphaFunction(value), thickness);

            //    if (start == 0 && alpha > 0.0) {
            //        start = i;
            //    }
            //    if (alpha > 0.99999) {
            //        end = i;
            //        break;
            //    }
            //}

            vec4 color = vec4(0);
            for (int i=end; i>start; i--) {
                vec3 point = position + lerp3(ray*near, ray*far, float(i)/float(samples));
                float value = sample(point);
                color = mix(color, transferFunction(value), thickness);
            }

            //vec4 color = vec4(0);
            //for (int i=samples-1; i>=0; i--) {
            //    vec3 point = position + lerp3(ray*far, ray*near, float(i)/float(samples));
            //    float value = sample(point);
            //    color = mix(color, transferFunction(value), thickness);
            //}
            return color;
        }

        void main(void) {
            vec3 ray = getRay(gl_TexCoord[0].s*float(width),
                              gl_TexCoord[0].t*float(height),
                              width, height);
            gl_FragColor = sampleRay(ray);
        }
    `)

    vr.dist = dist


    //dim := 256
    //vr.volumeDataTexture, _ = readTexture3DBinary(dim, dim, dim, "astro512.txt")

    //dim := 64
    //vr.volumeDataTexture, _ = readTexture3D(dim, dim, dim, "astro64.txt")
    vr.volumeDataTexture, _ = sf.CreateTexture()

    volumeData := vr.program.GetUniformLocation("volumeData")
    volumeData.Uniform1i(0)

    vr.transferFunctionTexture = SetupTransferFunction()

    transferFunctionMin := vr.program.GetUniformLocation("transferFunctionMin")
    transferFunctionMin.Uniform1f(min)

    transferFunctionMax := vr.program.GetUniformLocation("transferFunctionMax")
    transferFunctionMax.Uniform1f(max)

    position := vr.program.GetUniformLocation("position")
    position.Uniform3f(-8.25e+7, -3.45e+7, 3.35e+7)

    up := vr.program.GetUniformLocation("up")
    up.Uniform3f(0.0, -1.0, 0.0)

    focus := vr.program.GetUniformLocation("focus")
    focus.Uniform3f(0.0, 0.0, 0.0)

    angle := vr.program.GetUniformLocation("angle")
    angle.Uniform1f(30.0)

    near := vr.program.GetUniformLocation("near")
    near.Uniform1f(7.5e+7)

    far := vr.program.GetUniformLocation("far")
    far.Uniform1f(1.8e+8)

    samples := vr.program.GetUniformLocation("samples")
    samples.Uniform1i(sampleCount)

    return vr
}

func (vr *VolumetricRenderer) SetSize(w int, h int) {
    width := vr.program.GetUniformLocation("width");
    width.Uniform1i(w)
    height := vr.program.GetUniformLocation("height");
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

    t += 0.03
    test := vr.program.GetUniformLocation("test")
    test.Uniform1f(float32(math.Sin(t) + 1.0) * 0.5)
    position := vr.program.GetUniformLocation("position")
    //dist := math.Sqrt(8.25e+7*8.25e+7 + 3.45e+7*3.45e+7 + 3.35e+7*3.35e+7)
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
