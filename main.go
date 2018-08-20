package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func interpolate(f []Point, p float64) float64 {
	// Since we will loop over all the points in the vector, capture n.
	// Also, convert n to usize because we will be using th iterator
	// associated types as indices.
	n := len(f)
	result := 0.0

	// Each point in the vector of "known" points will be interpolated
	// to calculate the point at f(0).
	for i := 0; i < n; i++ {
		term := f[i].y
		// A good old nested for loop :)
		for j := 0; j < n; j++ {
			if i != j {
				// X's should be unique
				//assert!(f[i].x - f[j].x != 0.0);
				denominator := f[i].x - f[j].x
				numerator := -f[j].x
				term = term * (numerator / denominator)
			}
		}

		result += term
		result = math.Mod(result, p)
	}

	return result
}

func main() {
	f := make([]Point, 3)
	f[0] = Point{x: 4.0, y: 1.0}
	f[1] = Point{x: 3.0, y: 2.0}
	f[2] = Point{x: 5.0, y: 2.0}

	p := 11.0

	result := interpolate(f, p)
	fmt.Println("f(0) =", result)
}

// Output:
// f(0) = 6
