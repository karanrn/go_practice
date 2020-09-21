package distconv

// Kilometer .
type Kilometer float64

// Mile .
type Mile float64

/*
	1 Mile = 1.6 Kms
	1 Km = 0.6 Miles
*/

// KmsToMiles converts Kilometers to Miles, returns Mile
func KmsToMiles(kms Kilometer) Mile {
	return Mile(kms * 0.6)
}

// MilesToKms converts Miles to Kms, returns Kilometer
func MilesToKms(miles Mile) Kilometer {
	return Kilometer(miles * 1.6)
}
