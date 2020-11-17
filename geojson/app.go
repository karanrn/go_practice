package main
/*
https://golangcode.com/is-point-within-polygon-from-geojson/

*/
import (
	"fmt"
	"io/ioutil"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

const (
	GEOFILE = "IndianStates.json"
)

func main() {
	// Open/load the file
	f, err := ioutil.ReadFile(GEOFILE)
	if err != nil {
		fmt.Errorf("error while reading json file, got %v", err.Error())
		return
	}

	featureCollections, err := geojson.UnmarshalFeatureCollection(f)
	/*
	// List all states
	for _, feature := range featureCollections.Features {
		fmt.Println(feature.Properties["NAME_1"])
	}
	*/

	// Find the state in which geolocation lies - [93.789047, 6.852571]
	result := isPointInsidePolygon(featureCollections, orb.Point{93.789047, 6.852571})
	fmt.Println(result)

	for _, feature := range featureCollections.Features {
		fmt.Println(feature.Geometry)
		break
	}

}

func isPointInsidePolygon(fc *geojson.FeatureCollection, point orb.Point) string {
    for _, feature := range fc.Features {
        // Try on a MultiPolygon to begin
        multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
        if isMulti {
            if planar.MultiPolygonContains(multiPoly, point) {
                return feature.Properties["NAME_1"].(string)
            }
        } else {
            // Fallback to Polygon
            polygon, isPoly := feature.Geometry.(orb.Polygon)
            if isPoly {
                if planar.PolygonContains(polygon, point) {
                    return feature.Properties["NAME_1"].(string)
                }
            }
        }
    }
    return ""
}