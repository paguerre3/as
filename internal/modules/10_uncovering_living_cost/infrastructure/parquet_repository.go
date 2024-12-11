package infrastructure

import (
	"github.com/fraugster/parquet-go/floor"
	"github.com/hashicorp/go-set/v3"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
)

type ParquetRepository interface {
	LoadNextMobilityDataBatch(batchSizeInRows int) ([]domain.MobilityData, bool, error)
	CloseLoading()
}

type H3UseCase interface {
	GenerateFeatures(mobilityData []domain.MobilityData) map[string]int
	CalculateH3Key(lat float64, lon float64) string
} // clonning interface in order to avoid DDD violation (maybe redundant)

type parquetRepositoryImpl struct {
	reader    *floor.Reader
	offset    int
	h3Keys    set.Collection[string]
	h3UseCase H3UseCase
}

func NewParquetRepository(filePath string, h3Keys set.Collection[string], h3UseCase H3UseCase) (ParquetRepository, error) {
	// Create a parquet reader using the floor API
	reader, err := floor.NewFileReader(filePath)
	if err != nil {
		return nil, err
	}
	return &parquetRepositoryImpl{
		reader:    reader,
		offset:    0,
		h3Keys:    h3Keys,
		h3UseCase: h3UseCase,
	}, nil
}

// Loads mobility data from a Parquet file
func (r *parquetRepositoryImpl) LoadNextMobilityDataBatch(batchSizeInRows int) (data []domain.MobilityData, endOfFile bool, err error) {
	end := r.offset + batchSizeInRows
	// Iterate through rows in the parquet file until batchsieInRows is reached:
	for r.offset < end {
		next := r.reader.Next()
		if !next {
			endOfFile = true
			return data, endOfFile, nil
		}
		var record domain.MobilityData
		if err = r.reader.Scan(&record); err != nil {
			return nil, endOfFile, err
		}
		if r.validH3Key(record) {
			data = append(data, record)
		}
		// don't all inside valid key region as GC will be called later on
		r.offset++
	}

	/*if err := r.reader.Err(); err != nil {
		return nil, err
	}*/
	return data, endOfFile, nil
}

func (r *parquetRepositoryImpl) CloseLoading() {
	r.reader.Close()
}

func (r *parquetRepositoryImpl) validH3Key(record domain.MobilityData) bool {
	h3Key := r.h3UseCase.CalculateH3Key(record.Lat, record.Lon)
	return h3Key != "" && r.h3Keys.Contains(h3Key)
}
