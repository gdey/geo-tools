package maps

import (
	"math"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
)

func BoundsCenterZoom(proj crs.ProjectTransformer, bounds geo.LatLngBounds, size geometry.Point, maxZoom float64) (geo.LatLng, float64) {

	//	calculate our zoom
	zoom := BoundsZoom(proj, bounds, size, maxZoom)

	//	convert to points
	swPoint := crs.LatLngToPoint(proj, bounds.SouthWest(), zoom)
	nePoint := crs.LatLngToPoint(proj, bounds.NorthEast(), zoom)

	//	find center
	center := crs.PointToLatLng(proj, swPoint.Add(nePoint).DivideBy(2), zoom)

	return center, zoom
}

//	returns the zoom level for supplied bounds
//	useful when rendering static map images
//	pass in 0.0 for maxZoom for no max zoom
//
//	TODO: add padding support
func BoundsZoom(proj crs.ProjectTransformer, bounds geo.LatLngBounds, size geometry.Point, maxZoom float64) float64 {
	var zoom float64

	nw := bounds.NorthWest()
	se := bounds.SouthEast()

	b := geometry.NewBounds(crs.LatLngToPoint(proj, se, zoom), crs.LatLngToPoint(proj, nw, zoom))
	boundsSize := b.Size()

	scale := math.Min(size.X/boundsSize.X, size.Y/boundsSize.Y)

	zoom = math.Floor(ScaleZoom(proj, scale, zoom))
	if zoom > maxZoom {
		return maxZoom
	}

	return zoom
}

func ScaleZoom(proj crs.ProjectTransformer, scale float64, fromZoom float64) float64 {
	return crs.Zoom(scale * crs.Scale(fromZoom))
}
