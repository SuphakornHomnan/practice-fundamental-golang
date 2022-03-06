package main

import "fmt"

type rectangle struct {
	w, l float64
}

// Function style
// func Area(rec rectangle) float64 {
// 	return rec.l * rec.w
// }

// Method style
func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

func (rec *rectangle) SetWidth(w float64) {
	rec.w = w
}

func main() {
	rec := rectangle{
		w: 13,
		l: 173,
	}

	fmt.Println(rec.Area())

	rec.SetWidth(52)
	fmt.Println(rec.Area())
}
