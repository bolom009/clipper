package clipper

import (
	"github.com/bolom009/geom"
	"math"
)

// OffsetPolygon generate new polygon with clipper offset
func OffsetPolygon(polygon []geom.Vector2, offset float32) []geom.Vector2 {
	offsetPoints := make([]geom.Vector2, len(polygon))
	for i := 0; i < len(polygon); i++ {
		prevIndex := (i - 1 + len(polygon)) % len(polygon)
		nextIndex := (i + 1) % len(polygon)

		// Edge vectors
		edge1 := geom.Vector2{X: polygon[i].X - polygon[prevIndex].X, Y: polygon[i].Y - polygon[prevIndex].Y}
		edge2 := geom.Vector2{X: polygon[nextIndex].X - polygon[i].X, Y: polygon[nextIndex].Y - polygon[i].Y}

		// Compute edge normals
		normal1 := geom.Vector2{X: -edge1.Y, Y: edge1.X}
		normal2 := geom.Vector2{X: -edge2.Y, Y: edge2.X}

		// Normalize the normals
		norm1Length := float32(math.Sqrt(float64(normal1.X*normal1.X + normal1.Y*normal1.Y)))
		norm2Length := float32(math.Sqrt(float64(normal2.X*normal2.X + normal2.Y*normal2.Y)))
		if norm1Length > 0 {
			normal1.X /= norm1Length
			normal1.Y /= norm1Length
		}
		if norm2Length > 0 {
			normal2.X /= norm2Length
			normal2.Y /= norm2Length
		}

		// Calculate average of normals for the offset
		averageNormal := geom.Vector2{
			X: (normal1.X + normal2.X) / 2,
			Y: (normal1.Y + normal2.Y) / 2,
		}

		// Normalize the average normal
		averageNormalLength := float32(math.Sqrt(float64(averageNormal.X*averageNormal.X + averageNormal.Y*averageNormal.Y)))
		if averageNormalLength > 0 {
			averageNormal.X /= averageNormalLength
			averageNormal.Y /= averageNormalLength
		}

		// Create the offset point
		offsetPoints[i] = geom.Vector2{
			X: polygon[i].X + averageNormal.X*offset,
			Y: polygon[i].Y + averageNormal.Y*offset,
		}
	}

	return offsetPoints
}
