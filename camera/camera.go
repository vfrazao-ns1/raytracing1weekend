package camera

import (
	"math"

	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/utils"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	focalLength = 1.0
)

// Camera struct to represent our camera
type Camera struct {
	LookFrom       vec3.Point // LookFrom point where the camera is positioned - looking from this point
	LookAt         vec3.Point // LookAt point where the camera is looking at
	ViewUp         vec3.Vec3  // ViewUp vector dictating which direction is up - where's the sky?
	VFOV           float64    // VFOV Vertical field of view
	Theta          float64    // Theta field of view in radian
	H              float64    // H the ratio of half height of the FOV at z distance
	AspectRatio    float64    // AspectRatio usually 16:9
	ViewPortHeight float64    // ViewPortHeight usually 2.0
	ViewPortWidth  float64    // ViewPortWidth constrained by the aspect ratio and height
	FocalLength    float64    // FocalLength usually 1.0

	U          vec3.Vec3
	V          vec3.Vec3
	W          vec3.Vec3
	LensRadius float64

	Origin     vec3.Point // Origin usually (0, 0, 0)
	Horizontal vec3.Vec3  // Horizontal line: <width, 0, 0>
	Vertical   vec3.Vec3  // Vertical line <0, height, 0>
	LowerLeft  vec3.Point // LowerLeft corner of view
}

// InitCamera Creates and initializes a Camera struct
func InitCamera(lookfrom vec3.Point, lookat vec3.Point, vup vec3.Vec3, vfov, aspect, aperture, focusDist float64) *Camera {
	c := new(Camera)

	c.LookFrom = lookfrom
	c.LookAt = lookat
	c.ViewUp = vup

	c.VFOV = vfov
	c.Theta = utils.Degrees2radians(vfov)
	c.H = math.Tan(c.Theta / 2.0)
	c.AspectRatio = aspect
	c.ViewPortHeight = 2 * c.H
	c.ViewPortWidth = c.AspectRatio * c.ViewPortHeight
	c.FocalLength = focalLength

	c.W = lookfrom.Sub(lookat).Unit()
	c.U = vup.Cross(c.W).Unit()
	c.V = c.W.Cross(c.U)

	c.Origin = lookfrom
	c.Horizontal = c.U.ScalarMul(c.ViewPortWidth).ScalarMul(focusDist)
	c.Vertical = c.V.ScalarMul(c.ViewPortHeight).ScalarMul(focusDist)
	c.LowerLeft = c.Origin.Sub(c.Horizontal.ScalarDiv(2)).Sub(c.Vertical.ScalarDiv(2)).Sub(c.W.ScalarMul(focusDist))
	c.LensRadius = aperture / 2.0

	return c
}

// GetRay returns the ray from our camera
func (c Camera) GetRay(s, t float64) ray.Ray {
	rd := utils.RandomInUnitDisk().ScalarMul(c.LensRadius)
	offset := c.U.ScalarMul(rd.X).Add(c.V.ScalarMul(rd.Y))
	return ray.Ray{
		Origin:    c.Origin.Add(offset),
		Direction: c.LowerLeft.Add((c.Horizontal.ScalarMul(s))).Add(c.Vertical.ScalarMul(t)).Sub(c.Origin).Sub(offset),
	}
}
