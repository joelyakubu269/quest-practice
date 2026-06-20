type interface struct {
	Area() float64
}
type Triangle {
	Base float64
	Height float64
}
func (p Triangle) Area() float64 {
	return p.Base * p.Height
}