package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (p *Square) End() Point {
	end := Point{
		x: p.start.x + int(p.a),
		y: p.start.y + int(p.a),
	}
	return end
}

func (p *Square) Area() uint {
	return p.a * p.a
}

func (p *Square) Perimeter() uint {
	return 4 * p.a
}
