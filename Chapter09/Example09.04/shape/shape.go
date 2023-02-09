package shape

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

func GetName(shape Shape) string {
	return shape.name()
}

func GetArea(shape Shape) float64 {
	return shape.area()
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height) / 2
	// return (t.Base * t.Height)
}

func (t Triangle) name() string {
	return "三角形"
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "長方形"
	// return "矩形"
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) name() string {
	return "正方形"
}
