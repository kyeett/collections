package grid

type Point struct {
	X, Y int
}
type Grid struct {
	values     []interface{}
	cols, rows int
}

func New(cols, rows int) *Grid {
	return &Grid{
		values: make([]interface{}, cols*rows),
		cols:   cols,
		rows:   rows,
	}
}

func (g *Grid) Do(f func(p Point, value interface{})) {
	for x := 0; x < g.cols; x++ {
		for y := 0; y < g.rows; y++ {
			f(Point{x, y}, g.values[y*g.cols+x])
		}
	}
}

func (g *Grid) Get(p Point) interface{} {
	if p.X < 0 || p.Y < 0 || p.X >= g.cols || p.Y >= g.rows {
		return nil
	}
	return g.values[p.Y*g.cols+p.X]
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) Cols() int {
	return g.cols
}

func (g *Grid) Len() int {
	return g.rows * g.cols
}

func (g *Grid) Set(p Point, v interface{}) {
	g.values[p.Y*g.cols+p.X] = v
}
