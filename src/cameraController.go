package main

import (glfw "github.com/go-gl/glfw3")

type CameraController struct {
    manager *GlfwManager
    camera *VolumetricRenderer
}

func CreateCameraController(manager *GlfwManager, camera *VolumetricRenderer) *CameraController {
    controller := new(CameraController)
    controller.manager = manager
    controller.camera = camera
    return controller
}

func (controller *CameraController) Update(dt float64) {
    rotateSpeed := dt * 5.0
    zoomSpeed := dt * 100.0
    if  controller.manager.Window.GetKey(glfw.KeyLeft) == glfw.Press ||
        controller.manager.Window.GetKey(glfw.KeyLeft) == glfw.Repeat {
        controller.camera.Rotate(-rotateSpeed)
    }
    if  controller.manager.Window.GetKey(glfw.KeyRight) == glfw.Press ||
        controller.manager.Window.GetKey(glfw.KeyRight) == glfw.Repeat {
        controller.camera.Rotate(rotateSpeed)
    }
    if  controller.manager.Window.GetKey(glfw.KeyUp) == glfw.Press ||
        controller.manager.Window.GetKey(glfw.KeyUp) == glfw.Repeat {
        controller.camera.Zoom(-zoomSpeed)
    }
    if  controller.manager.Window.GetKey(glfw.KeyDown) == glfw.Press ||
        controller.manager.Window.GetKey(glfw.KeyDown) == glfw.Repeat {
        controller.camera.Zoom(zoomSpeed)
    }
}
