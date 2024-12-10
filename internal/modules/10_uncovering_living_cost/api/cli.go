package api

import (
	"fmt"
	"time"

	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/application"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/infrastructure"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
	"golang.org/x/exp/rand"
)

const (
	internalTrainCsvPath   = "internal/modules/10_uncovering_living_cost/assets/train.csv"
	internalParquetPath    = "internal/modules/10_uncovering_living_cost/assets/mobility_data.parquet"
	features_resolution    = 9 // between 0 (biggest cell) and 15 (smallest cell) -number suggested by the library
	internalTestCsvPath    = "internal/modules/10_uncovering_living_cost/assets/test.csv"
	internalPredictCsvPath = "internal/modules/10_uncovering_living_cost/assets/out/predict_%d.csv"
)

var (
	trainCsvPathResolver   = common_infra.NewPathResolver()
	parquetPathResolver    = common_infra.NewPathResolver()
	testCsvPathResolver    = common_infra.NewPathResolver()
	predictCsvPathResolver = common_infra.NewPathResolver()
)

func RunCommandLine() error {
	csvRepo := infrastructure.NewCSVRepository()
	parquetRepo := infrastructure.NewParquetRepository()
	h3UseCase := application.NewH3UseCase()
	modelUseCases := application.NewModelUseCases()

	// 0. load train data
	trainCsvPath := trainCsvPathResolver(internalTrainCsvPath)
	trainData, err := csvRepo.LoadTrainData(trainCsvPath)
	if err != nil {
		fmt.Println("Error loading train data:", err)
		return err
	}

	// 1. load mobility data from parquet
	mobilityParquetPath := parquetPathResolver(internalParquetPath)
	// all mobility data is loaded in memory:
	mobilityData, err := parquetRepo.LoadMobilityData(mobilityParquetPath)
	if err != nil {
		fmt.Println("Error loading mobility data:", err)
		return err
	}

	// 2. generate mobility data features from parquet in junks for higher performance:
	features := h3UseCase.GenerateFeatures(mobilityData, features_resolution)

	// 3. traing model using train and parquet data
	err = modelUseCases.TrainModel(trainData, features)
	if err != nil {
		fmt.Println("Error training model:", err)
		return err
	}

	// 4. load test csv as the input for predictions
	testCsvPath := testCsvPathResolver(internalTestCsvPath)
	testData, err := csvRepo.LoadTestData(testCsvPath)
	if err != nil {
		fmt.Println("Error loading test data:", err)
		return err
	}

	// 5. "make predictions" using testData and features
	predictions, err := modelUseCases.Predict(testData, features)
	if err != nil {
		fmt.Println("Error making predictions:", err)
		return err
	}

	// 6. save predictions to csv
	predictCsvPath := calculatePredictPath()
	err = csvRepo.SavePredictions(predictCsvPath, predictions)
	if err != nil {
		fmt.Println("Error saving predictions:", err)
		return err
	}

	fmt.Println("Predictions saved successfully to: ", predictCsvPath)
	return nil
}

func calculatePredictPath() string {
	// Seed the random number generator with the current time
	rand.Seed(uint64(time.Now().UnixNano()))
	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1
	path := fmt.Sprintf(internalPredictCsvPath, randomNumber)
	return predictCsvPathResolver(path)
}
