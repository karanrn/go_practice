package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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
	Lat float32
	Lon float32
}

type stateCoordinates struct {
	State    string
	Location []geoLocation
}

type geoGeometry struct {
	GeometryType string            `json:"type"`
	Coordinates  [][][]interface{} `json:"coordinates"`
}

type geojson struct {
	GeoType    string      `json:"type"`
	Properties geoProperty `json:"properties"`
	Geometry   geoGeometry `json:"geometry"`
}

type featureCollection struct {
	FType    string    `json:"type"`
	Features []geojson `json:"features"`
}

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

	// List all states
	for _, state := range featureSet.Features {
		fmt.Println(state.Properties.NAME1)
	}

	// [93.789047, 6.852571]
	for _, s := range featureSet.Features {
		//var stateLoc stateCoordinates

		if s.Geometry.GeometryType == "MultiPolygon" {
			//fmt.Println(s.Properties.NAME1, s.Geometry.Coordinates[0][0])
			for _, g1 := range s.Geometry.Coordinates[0][0] {
				fmt.Println(reflect.TypeOf(g1), g1)
				/*
					stateLoc.State = s.Properties.NAME1
					stateLoc.Location[i].Lat = g1
				*/
			}
		}
		/*
			if s.Geometry.GeometryType == "Polygon" {
				fmt.Println("Polygon:", s.Properties.NAME1, s.Geometry.Coordinates[0][0])
			}
		*/
	}
}
