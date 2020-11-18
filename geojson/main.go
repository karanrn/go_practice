package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "reflect"
)

type geoProperty struct {
	ID0       int    `json:"ID_0"`
	ISO       string `json:"ISO"`
	NAME0     string `json:"NAME_0"`
	ID1       int    `json:"ID_1"`
	NAME1     string `json:"NAME_1"`
	NLNAME1   string `json:"NL_NAME_1"`
	VARNAME1  string `json:"VARNAME_1"`
	TYPE1     string `json:"TYPE_1"`
	ENGTYPE1  string `json:"ENGTYPE_1"`
	FileName  string `json:"filename"`
	FileName1 string `json:"filename_1"`
	FileName2 string `json:"filename_2"`
	FileName3 string `json:"filename_3"`
}

type geoLocation struct {
	Lat float64
	Lon float64
}

type stateCoordinates struct {
	State    string
	Location []geoLocation
}

type geoGeometry struct {
	GeometryType string            `json:"type"`
	Coordinates  [][][]interface{} `json:"coordinates"`
}

type geoJSON struct {
	GeoType    string      `json:"type"`
	Properties geoProperty `json:"properties"`
	Geometry   geoGeometry `json:"geometry"`
}

type featureCollection struct {
	FType    string    `json:"type"`
	Features []geoJSON `json:"features"`
}

// Contains preforms linear search on the data
//func Contains()
func main() {

	data, err := ioutil.ReadFile("./IndianStates.json")
	if err != nil {
		fmt.Printf("error reading file: %v", err.Error())
		return
	}

	var featureSet featureCollection

	err = json.Unmarshal(data, &featureSet)
	if err != nil {
		fmt.Printf("error unmarshalling/decoding json: %v", err.Error())
		return
	}

	/*
		// List all states
		for _, state := range featureSet.Features {
			fmt.Println(state.Properties.NAME1)
		}
	*/

	// [93.789047, 6.852571]
	// Unmarshall coordinates from the geoJson data
	var stateMap []stateCoordinates
	for _, s := range featureSet.Features {
		var stateLoc stateCoordinates
		if s.Geometry.GeometryType == "MultiPolygon" {
			stateLoc.State = s.Properties.NAME1
			for _, g1 := range s.Geometry.Coordinates[0][0] {
				points, ok := g1.([]interface{})
				if !ok {
					fmt.Errorf("not a valid position, got %v", g1)
				}
				gdata := make([]float64, 0, len(points))
				for _, coord := range points {
					if f, ok := coord.(float64); ok {
						gdata = append(gdata, f)
					} else {
						fmt.Errorf("not a valid coordinate, got %v", coord)
					}
				}
				geoLoc := geoLocation{
					Lat: gdata[0],
					Lon: gdata[1],
				}
				stateLoc.Location = append(stateLoc.Location, geoLoc)
			}
		}

		if s.Geometry.GeometryType == "Polygon" {
			stateLoc.State = s.Properties.NAME1
			for _, g2 := range s.Geometry.Coordinates[0] {

				geoLoc := geoLocation{
					Lat: g2[0].(float64),
					Lon: g2[1].(float64),
				}
				stateLoc.Location = append(stateLoc.Location, geoLoc)
			}
		}

		stateMap = append(stateMap, stateLoc)

	}

	for _, state := range stateMap {
		//fmt.Printf("%v:%v\n", state.State, len(state.Location))
		fmt.Printf("%v: \n Max - %v, %v \n Min - %v, %v")
	}

	// Check for the location and return the state
}

func minLoc(gLocs []geoLocation)
