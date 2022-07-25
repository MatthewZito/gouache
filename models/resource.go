package models

type Resource struct {
	X string
	Y int
	Z int
}

func NewResource(x string, y int, z int) *Resource {
	r := &Resource{
		X: x,
		Y: y,
		Z: z,
	}

	return r
}
