package domain

// Represents training data from `train.csv`
type TrainData struct {
	HexID        string
	CostOfLiving float64
}

// Represents a row in `mobility.parquet`
type MobilityData struct {
	DeviceID  string
	Lat       float64
	Lon       float64
	Timestamp int64
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
