package application

import (
	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
	"github.com/sajari/regression"
)

type ModelUseCases interface {
	TrainModel(trainData []domain.TrainData, features map[string]int) error
	Predict(testData []domain.TestData, features map[string]int) ([]domain.Prediction, error)
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
func (m *modelUseCasesImpl) TrainModel(trainData []domain.TrainData, features map[string]int) error {
	for _, data := range trainData {
		count := float64(features[data.HexID])
		m.model.Train(regression.DataPoint(data.CostOfLiving, []float64{count}))
	}
	m.model.Run()
	return nil
}

// Predict makes predictions using the trained model
func (m *modelUseCasesImpl) Predict(testData []domain.TestData, features map[string]int) ([]domain.Prediction, error) {
	var predictions []domain.Prediction
	for _, test := range testData {
		count := float64(features[test.HexID])
		prediction, err := m.model.Predict([]float64{count})
		if err != nil {
			return nil, err
		}
		predictions = append(predictions, domain.Prediction{
			HexID:        test.HexID,
			CostOfLiving: prediction,
		})
	}
	return predictions, nil
}
