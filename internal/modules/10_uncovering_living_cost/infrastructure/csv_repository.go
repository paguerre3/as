package infrastructure

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/hashicorp/go-set/v3"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
)

type CSVRepository interface {
	LoadTrainData(fileName string) ([]domain.TrainData, *set.Set[string], error)
	LoadTestData(fileName string) ([]domain.TestData, *set.Set[string], error)
	SavePredictions(fileName string, predictions []domain.Prediction) error
}

type csvRepositoryImpl struct {
}

func NewCSVRepository() CSVRepository {
	return &csvRepositoryImpl{}
}

// Loads training data from `train.csv`
func (r *csvRepositoryImpl) LoadTrainData(filePath string) ([]domain.TrainData, *set.Set[string], error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var data []domain.TrainData
	hs := set.New[string](100)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		cost, _ := strconv.ParseFloat(record[1], 64)
		hexId := record[0]
		data = append(data, domain.TrainData{HexID: hexId, CostOfLiving: cost})
		hs.Insert(hexId)
	}
	return data, hs, nil
}

func (r *csvRepositoryImpl) LoadTestData(filePath string) ([]domain.TestData, *set.Set[string], error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var data []domain.TestData
	hs := set.New[string](100)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		hexId := record[0]
		data = append(data, domain.TestData{HexID: hexId})
		hs.Insert(hexId)
	}
	return data, hs, nil
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
