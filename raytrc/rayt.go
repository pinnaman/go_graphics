package raytrc

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// Vector - struct holding X Y Z values of a 3D vector
type Vector struct {
	X, Y, Z float64
}

// Add - adds two vectors together
func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Sub - subtracts b Vector from a Vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// MultiplyByScalar - multiplies a Vector by s float64
func (a Vector) MultiplyByScalar(s float64) Vector {
	return Vector{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

// Length - calculates the length(magnitude) of the Vector
func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Dot - calculates the dot product of two Vectors
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross - calculates the cross product of two Vectors
func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Normalize - returns a versor created from the given vector
func (a Vector) Normalize() Vector {
	return a.MultiplyByScalar(1. / a.Length())
}

// Scene represents the s on which the scene is projected as a 2D picture
type Scene struct {
	Width, Height int
	Img           *image.RGBA
}

// NewScene returns a new Scene
func NewScene(width int, height int) *Scene {
	return &Scene{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// EachPixel traverses the image s and calls the provided function for each pixel
func (s *Scene) EachPixel(colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.setPixel(x, y, colorFunction(x, y))
		}
	}
}

// Save exports the image to hdd
func (s *Scene) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, s.Img)
}

func (s *Scene) setPixel(x int, y int, color color.RGBA) {
	s.Img.Set(x, y, color)
}

/*
func main() {

	var width = 200
	var height = 400
	scene := NewScene(width, height)
	scene.EachPixel(func(x, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / width),
			uint8(y * 255 / height),
			150,
			255,
		}
	})
	scene.Save(fmt.Sprintf("./renders/%d.png", time.Now().Unix()))
}
*/
