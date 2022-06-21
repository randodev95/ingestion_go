package utils

type Address struct {
	CountryCode string     `json:"country_code"`
	Country     string     `json:"country"`
	Region      string     `json:"region"`
	City        string     `json:"city"`
	ZipCode     string     `json:"zipcode"`
	LatLong     [2]float32 `json:"lat_long"`
}
