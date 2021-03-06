package vec3

import (
	"fmt"
	"math"
)

// Vec3 vector class
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// These are all alternate implementations of LengthSquared in ASM
// These are all slower since they can't be inlined so they spec a ton
// of time putting and taking stuff on the stack
func lensq(x, y, z float64) float64
func lensq2(x, y, z float64) float64
func lensq3(x, y, z float64) float64
func lensq4(x, y, z float64) float64
func lensq5(x, y, z float64) float64

// LengthSquared length squared of the vector (duh)
func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

// Length length of the vector (duh)
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// String implements the stringer interface
func (v Vec3) String() string {
	return fmt.Sprintf("[%.1f %.1f %.1f]", v.X, v.Y, v.Z)
}

// Add adds two Vec3s together
func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Sub subtracts two Vec3s together
func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// Negate Negates a Vec3
func (v Vec3) Negate() Vec3 {
	return Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Mul multiplies two Vec3s together
func (v Vec3) Mul(other Vec3) Vec3 {
	return Vec3{X: v.X * other.X, Y: v.Y * other.Y, Z: v.Z * other.Z}
}

// ScalarAdd adds a Vec3 by a constant
func (v Vec3) ScalarAdd(scalar float64) Vec3 {
	return Vec3{X: v.X + scalar, Y: v.Y + scalar, Z: v.Z + scalar}
}

// ScalarMul multiplies a Vec3 by a constant
func (v Vec3) ScalarMul(scalar float64) Vec3 {
	return Vec3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// ScalarDiv divides a Vec3 by a constant
func (v Vec3) ScalarDiv(scalar float64) Vec3 {
	return Vec3{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}

// Dot dot product of two Vec3s
func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross cross product of two Vec3
func (v Vec3) Cross(other Vec3) Vec3 {
	x := v.Y*other.Z - v.Z*other.Y
	y := v.Z*other.X - v.X*other.Z
	z := v.X*other.Y - v.Y*other.X
	return Vec3{X: x, Y: y, Z: z}
}

// Unit returns the unit vector
func (v Vec3) Unit() Vec3 {
	return v.ScalarDiv(v.Length())
}

// Reflect returns the vector resulting from reflecting off a material
func Reflect(v Vec3, normal Vec3) Vec3 {
	return v.Sub(normal.ScalarMul(v.Dot(normal) * 2))
}

// Refract returns the vector resulting from refracting off a material
func Refract(uv Vec3, normal Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := uv.Negate().Dot(normal)
	// cosTheta := utils.Fmin(uv.Negate().Dot(normal))
	rOutParallel := normal.ScalarMul(cosTheta).Add(uv).ScalarMul(etaiOverEtat)
	rOutPerp := normal.ScalarMul(-math.Sqrt(1 - rOutParallel.LengthSquared()))
	return rOutParallel.Add(rOutPerp)
}
