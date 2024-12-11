package domain

// Represents training data from `train.csv`
type TrainData struct {
	HexID        string
	CostOfLiving float64
}

// MobilityData represents a record in the mobility_data.parquet file type
type MobilityData struct {
	DeviceID  int32   `parquet:"device_id"`
	Lat       float64 `parquet:"lat"`
	Lon       float64 `parquet:"lon"`
	Timestamp int64   `parquet:"timestamp"`
}

// Represents test data from `test.csv`
type TestData struct {
	HexID string
}

// Represents a prediction for `test.csv`
type Prediction struct {
	HexID        string
	CostOfLiving float64
}
