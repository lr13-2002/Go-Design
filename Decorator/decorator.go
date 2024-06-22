package decorator

type IDraw interface {
	Draw() string
}

type Sqare struct{}

func (s Sqare) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	IDraw
	color  string
}

func (c ColorSquare) Draw() string {
	return c.IDraw.Draw() + ", color is " + c.color
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{IDraw: square, color: color}
}
