package application

import (
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
	"github.com/sajari/regression"
)

type ModelUseCases interface {
	TrainModel(trainData []domain.TrainData, features map[string]int) bool
	RunModel() error
	Predict(testData []domain.TestData, features map[string]int) (bool, []domain.Prediction, error)
}

// Handles machine learning model operations
type modelUseCasesImpl struct {
	model *regression.Regression
}

// Initializes a new ModelService
func NewModelUseCases() ModelUseCases {
	r := new(regression.Regression)
	r.SetObserved("CostOfLiving")
	r.SetVar(0, "Count")
	return &modelUseCasesImpl{model: r}
}

// TrainModel trains the regression model using train data and features
func (m *modelUseCasesImpl) TrainModel(trainData []domain.TrainData, features map[string]int) (trainedAtLeastOnce bool) {
	for _, data := range trainData {
		count, exists := features[data.HexID]
		if !exists {
			continue
		}
		m.model.Train(regression.DataPoint(data.CostOfLiving, []float64{float64(count)}))
		trainedAtLeastOnce = true
	}
	return trainedAtLeastOnce
}

func (m *modelUseCasesImpl) RunModel() error {
	return m.model.Run()
}

// Predict makes predictions using the trained model
func (m *modelUseCasesImpl) Predict(testData []domain.TestData, features map[string]int) (bool, []domain.Prediction, error) {
	var predictedatLeastOne bool
	var predictions []domain.Prediction
	for _, test := range testData {
		count, exists := features[test.HexID]
		if !exists {
			continue
		}
		prediction, err := m.model.Predict([]float64{float64(count)})
		if err != nil {
			return predictedatLeastOne, nil, err
		}
		predictedatLeastOne = true
		predictions = append(predictions, domain.Prediction{
			HexID:        test.HexID,
			CostOfLiving: prediction,
		})
	}
	return predictedatLeastOne, predictions, nil
}
