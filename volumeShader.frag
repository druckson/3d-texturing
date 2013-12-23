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
uniform int screenWidth;
uniform int screenHeight;
uniform int samples;

uniform vec3 size;

uniform float test;

float sample(vec3 pos) {
    pos = pos / size / 2.0 + vec3(0.5);
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
    float wangle = angle * float(screenWidth) / float(screenHeight);

    vec3 rx = ru*(2.0*tan(wangle*pi/360.0)/float(screenWidth));
    vec3 ry = rv*(2.0*tan(hangle*pi/360.0)/float(screenHeight));

    return normalize(-position) + 
        (rx * ((2.0*i + 1.0 - float(screenWidth)) / 2.0)) +
        (ry * ((2.0*j + 1.0 - float(screenHeight)) / 2.0));
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
    float thickness = len/float(samples);

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
    return color;
}

void main(void) {
    vec3 ray = getRay(gl_TexCoord[0].s*float(screenWidth),
                      gl_TexCoord[0].t*float(screenHeight),
                      screenWidth, screenHeight);
    gl_FragColor = sampleRay(ray);
}
