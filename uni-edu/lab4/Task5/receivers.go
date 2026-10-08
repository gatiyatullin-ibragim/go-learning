package Task5


type Report struct{
	Title   string
	Values  [1024]float64
}

func (r Report) FirstValue() float64 {
	return  r.Values[0]
}

func (r *Report) FirstPointer() float64 {
	return r.Values[0]
}

type Point struct{
	X, Y  float64
}

func (p Point) SumValue() float64 {
	return p.X + p.Y
}

func (p *Point) SumPointer() float64 {
	return p.X + p.Y
}