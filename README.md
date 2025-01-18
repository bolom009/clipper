# Clipper

## A polygon Clipping library

Geometric offsetting refers to the process of creating parallel curves that are offset a specified distance from their primary curves.


## Demo



## How to use

```go
package main

import (
	"fmt"
	
	"github.com/bolom009/clipper"
	"github.com/bolom009/geom"
)

func main() {
	polygon := []geom.Vector2{
		{X: 0, Y: 0},
		{X: 2, Y: 1},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 2, Y: 5},
		{X: 0, Y: 4},
	}

	offsetDistance := float32(0.5) // Set inward offset distance
	offsetPolygon := clipper.OffsetPolygon(polygon, offsetDistance)

	// Output results
	fmt.Println("Original Polygon:")
	for _, p := range polygon {
		fmt.Printf("(%f, %f) ", p.X, p.Y)
	}
	fmt.Println("\nOffset Polygon:")
	for _, p := range offsetPolygon {
		fmt.Printf("(%f, %f) ", p.X, p.Y)
	}
}


```