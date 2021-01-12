package s2covergeojson

import (
	"errors"

	"github.com/golang/geo/s2"
	"github.com/paulmach/go.geojson"
)

// geojsonLoopsToS2Loops converts [][][2] to array of s2 Loops
func geojsonLoopsToS2Loops(geomLoops [][][]float64) []*s2.Loop {
	var loops []*s2.Loop
	for _, geomLoop := range geomLoops {
		var points []s2.Point
		for _, coordinate := range geomLoop {
			points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(coordinate[1], coordinate[0])))
		}
		loops = append(loops, s2.LoopFromPoints(points))
	}
	return loops
}

// geometryToS2Polygon converts geojson geometry object of geojson to s2 Polygon
func geometryToS2Polygon(geom *geojson.Geometry) (*s2.Polygon, error) {
	if geom.IsPolygon() {
		loops := geojsonLoopsToS2Loops(geom.Polygon)
		return s2.PolygonFromLoops(loops), nil
	} else if geom.IsMultiPolygon() {
		var combinedLoops []*s2.Loop
		for i := range geom.MultiPolygon {
			loops := geojsonLoopsToS2Loops(geom.MultiPolygon[i])
			for _, loop := range loops {
				combinedLoops = append(combinedLoops, loop)
			}
		}
		return s2.PolygonFromLoops(combinedLoops), nil
	}
	return nil, errors.New("Geometry is not polygon")
}

// Cover function returns s2 cell union from specified geojson feature and RegionCoverer
func Cover(feature *geojson.Feature, regionCoverer s2.RegionCoverer) (s2.CellUnion, error) {
	polygon, err := geometryToS2Polygon(feature.Geometry);
	if err != nil {
		return nil, err
	}

	return regionCoverer.Covering(s2.Region(polygon)), nil
}
