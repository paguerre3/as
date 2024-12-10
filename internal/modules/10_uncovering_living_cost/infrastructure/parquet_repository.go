package infrastructure

import (
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
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
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(domain.MobilityData), 4)
	if err != nil {
		return nil, err
	}
	defer pr.ReadStop()

	var data []domain.MobilityData
	if err := pr.Read(&data); err != nil {
		return nil, err
	}
	return data, nil
}
