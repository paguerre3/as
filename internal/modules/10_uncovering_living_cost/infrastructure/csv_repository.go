package infrastructure

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
)

type CSVRepository interface {
	LoadTrainData(fileName string) ([]domain.TrainData, error)
	LoadTestData(fileName string) ([]domain.TestData, error)
	SavePredictions(fileName string, predictions []domain.Prediction) error
}

type csvRepositoryImpl struct {
}

func NewCSVRepository() CSVRepository {
	return &csvRepositoryImpl{}
}

// Loads training data from `train.csv`
func (r *csvRepositoryImpl) LoadTrainData(filePath string) ([]domain.TrainData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var data []domain.TrainData
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		cost, _ := strconv.ParseFloat(record[1], 64)
		data = append(data, domain.TrainData{HexID: record[0], CostOfLiving: cost})
	}
	return data, nil
}

func (r *csvRepositoryImpl) LoadTestData(filePath string) ([]domain.TestData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var data []domain.TestData
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		data = append(data, domain.TestData{HexID: record[0]})
	}
	return data, nil
}

// Saves predictions to a CSV file
func (r *csvRepositoryImpl) SavePredictions(filePath string, predictions []domain.Prediction) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"hex_id", "cost_of_living"})

	for _, p := range predictions {
		writer.Write([]string{p.HexID, strconv.FormatFloat(p.CostOfLiving, 'f', 6, 64)})
	}
	return nil
}
