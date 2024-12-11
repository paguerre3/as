package infrastructure

import (
	"github.com/fraugster/parquet-go/floor"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
)

type ParquetRepository interface {
	LoadMobilityData(fileName string) ([]domain.MobilityData, error)
}

type parquetRepositoryImpl struct {
}

func NewParquetRepository() ParquetRepository {
	return &parquetRepositoryImpl{}
}

// Loads mobility data from a Parquet file
func (r *parquetRepositoryImpl) LoadMobilityData(filePath string) ([]domain.MobilityData, error) {
	// Create a parquet reader using the floor API
	reader, err := floor.NewFileReader(filePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var data []domain.MobilityData

	// Iterate through rows in the parquet file
	for reader.Next() {
		var record domain.MobilityData
		if err := reader.Scan(&record); err != nil {
			return nil, err
		}
		data = append(data, record)
	}

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
