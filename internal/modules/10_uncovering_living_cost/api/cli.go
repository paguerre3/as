package api

import (
	"fmt"
	"runtime"

	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/application"
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/infrastructure"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

const (
	internalTrainCsvPath   = "internal/modules/10_uncovering_living_cost/assets/train.csv"
	internalParquetPath    = "internal/modules/10_uncovering_living_cost/assets/mobility_data.parquet"
	internalTestCsvPath    = "internal/modules/10_uncovering_living_cost/assets/test.csv"
	internalPredictCsvPath = "internal/modules/10_uncovering_living_cost/assets/out/predict_%d.csv"
	batchSizeInRows        = 10000000 // num_rows: 340411133 / 5000000 = 68.0822266 predictions (being the latest the most accurate)
)

var (
	trainCsvPathResolver = common_infra.NewPathResolver()
	parquetPathResolver  = common_infra.NewPathResolver()
	testCsvPathResolver  = common_infra.NewPathResolver()
)

func RunCommandLine() error {
	h3UseCase := application.NewH3UseCase()
	modelUseCases := application.NewModelUseCases() // regression model holds the trained model state (stateful)

	// 0. load train data
	csvRepo := infrastructure.NewCSVRepository()
	trainCsvPath := trainCsvPathResolver(internalTrainCsvPath)
	trainData, h3TrainKeys, err := csvRepo.LoadTrainData(trainCsvPath)
	if err != nil {
		fmt.Println("Error loading train data:", err)
		return err
	}

	// 1. load test csv as the input for predictions
	testCsvPath := testCsvPathResolver(internalTestCsvPath)
	testData, h3TestKeys, err := csvRepo.LoadTestData(testCsvPath)
	if err != nil {
		fmt.Println("Error loading test data:", err)
		return err
	}

	// merge keys for both train and test data
	h3Keys := h3TrainKeys.Union(h3TestKeys)

	// 2. load mobility data from parquet in batches as it can't be read in paralell or using splits with stable lib:
	mobilityParquetPath := parquetPathResolver(internalParquetPath)
	parquetRepo, err := infrastructure.NewParquetRepository(mobilityParquetPath, h3Keys, h3UseCase) // parquet repo holds reader state (sateful)
	if err != nil {
		fmt.Println("Error creating parquet repository:", err)
		return err
	}
	defer parquetRepo.CloseLoading()

	features := make(map[string]int) // feaures filtered with h3keys mapping mobility data
	for {
		// the entire mobility data batch is loaded in memory:
		mobilityData, endOfFile, err := parquetRepo.LoadNextMobilityDataBatch(batchSizeInRows)
		if err != nil {
			fmt.Println("Error loading mobility data:", err)
			return err
		}
		if endOfFile {
			fmt.Println("End of file reached")
			break
		}

		if len(mobilityData) == 0 {
			fmt.Println("Mobility data is empty")

			// Trigger garbage collection
			runtime.GC()
			continue
		}

		// 3. generate mobility data features from parquet in junks for higher performance:
		mergeFeatures(features, h3UseCase.GenerateFeatures(mobilityData))

		// 4. traing model using train and parquet data
		trainedatLeastOnce := modelUseCases.TrainModel(trainData, features)
		if !trainedatLeastOnce {
			fmt.Println("No model trained")

			// Trigger garbage collection
			runtime.GC()
			continue
		}

		// Trigger garbage collection after processing the batch
		runtime.GC()
		// disable enable the following "break" for testing:
		fmt.Println("Model trained successfully")
		break
	}

	// 5. run model "once" after training finishes
	err = modelUseCases.RunModel()
	if err != nil {
		fmt.Println("Error running model:", err)
		return err
	}

	// 6. "make predictions" using testData and features in batches, last one will be the most accurate
	predictedAtLeastOnce, predictions, err := modelUseCases.Predict(testData, features)
	if err != nil {
		fmt.Println("Error making predictions:", err)
		return err
	}
	if !predictedAtLeastOnce {
		fmt.Println("No predictions made")
		return nil
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

func mergeFeatures(map1, map2 map[string]int) {
	for key, value := range map2 {
		map1[key] += value
	}
}

var calculatePredictPath = func() func() string {
	seed := 0
	return func() string {
		seed++
		path := fmt.Sprintf(internalPredictCsvPath, seed)

		predictCsvPathResolver := common_infra.NewPathResolver()
		return predictCsvPathResolver(path)
	}
}()
