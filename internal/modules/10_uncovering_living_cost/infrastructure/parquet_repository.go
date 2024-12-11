package infrastructure

import (
	"sync"

	"github.com/fraugster/parquet-go/floor"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
	"golang.org/x/sync/errgroup"
)

type ParquetRepository interface {
	LoadMobilityData(fileName string) ([]domain.MobilityData, error)
}

type parquetRepositoryImpl struct {
}

func NewParquetRepository() ParquetRepository {
	return &parquetRepositoryImpl{}
}

// Loads mobility data from a Parquet file with parallelism
func (r *parquetRepositoryImpl) LoadMobilityData(filePath string) ([]domain.MobilityData, error) {
	var data []domain.MobilityData
	var mu sync.Mutex // Mutex to safely append data to the result slice
	var g errgroup.Group

	const totalRows = 340411133
	batchSize := 1
	offset := 0

	for offset < totalRows {
		end := offset + batchSize
		if end > totalRows {
			end = totalRows
		}

		// Spawn a goroutine to process each batch
		g.Go(func() error {
			// Create a parquet reader using the floor API
			reader, err := floor.NewFileReader(filePath)
			if err != nil {
				return err
			}
			defer reader.Close()

			var chunk []domain.MobilityData

			// Read the chunk from the parquet file
			chunkOffset := offset
			for reader.Next() && chunkOffset < end {
				var record domain.MobilityData
				if err := reader.Scan(&record); err != nil {
					return err
				}
				chunk = append(chunk, record)
				chunkOffset++
			}

			// Lock the mutex and append the chunk to the data slice
			mu.Lock()
			defer mu.Unlock()
			data = append(data, chunk...)

			return nil
		})

		offset += batchSize
	}

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return data, nil
}
