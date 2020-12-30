package main

import "github.com/vfrazao-ns1/raytracing1weekend/vec3"

type config struct {
	FileName        string  // Name of file to save to render
	ImgWidth        int     // Resolution width
	Aspect          float64 // Aspect ratio float e.g., 16:9 equals 1.7777777
	SamplesPerPixel int     // How many rays to simulate hitting a given pixel (higher is better quality)
	MaxDepth        int
	Camera          cameraConfig // Camera config
}

type cameraConfig struct {
	LookFrom  vec3.Point // Initial camera position
	LookAt    vec3.Point // Point at which the camera is initially looking at
	Vup       vec3.Vec3  // ViewUp vector, where is up
	VFOV      float64    // Vertical field of view
	Aperture  float64    // Camera aperture
	FocusDist float64    // Camera focus distance
}
