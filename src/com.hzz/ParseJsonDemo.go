package temporary

//DriverTrackers parse json
type DriverTrackers struct {
	Status  int
	Code    int
	Message string
	Data    DataStruct
}

// DataStruct json data
type DataStruct struct {
	DriverID   int64 `json:"driverId"`
	DriverName string
	Locations  []Location
}

// Location location data
type Location struct {
	Latlng struct {
		Latitude  string
		Longitude string
	}
	Timestamp int64 `json:"#timestamp"`
}

// /*
//   unmarshal for json
// */
// func (o *DriverTrackers) UnmarshalJSON(data []byte]) {
// 	fmt.Println(data)
// 	return nil
// }
