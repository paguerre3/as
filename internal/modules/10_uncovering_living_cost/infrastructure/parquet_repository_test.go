package infrastructure

import (
	"testing"

	"github.com/hashicorp/go-set/v3"
	app "github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/application"
	domain "github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"
	"github.com/stretchr/testify/assert"
)

func TestIsValidKey(t *testing.T) {
	h3TrainKeys := set.New[string](100)
	h3TrainKeys.Insert("8866d3285dfffff")
	h3TestKeys := set.New[string](100)
	h3TestKeys.Insert("8866d3285dfffff")
	h3TestKeys.Insert("9966d3285c3ffff")
	h3Keys := h3TrainKeys.Union(h3TestKeys)
	uc := app.NewH3UseCase()
	pr := &parquetRepositoryImpl{
		offset:    0,
		h3Keys:    h3Keys,
		h3UseCase: uc,
	}
	assert.True(t, true)
	md := new(domain.MobilityData)
	md.Lat = -0.313305
	md.Lon = -78.536396
	assert.True(t, pr.validH3Key(*md))
	md = new(domain.MobilityData)
	md.Lat = -0.31339
	md.Lon = -78.536236
	assert.True(t, pr.validH3Key(*md))
	md = new(domain.MobilityData)
	md.Lat = -0.21349
	md.Lon = -58.536236
	assert.False(t, pr.validH3Key(*md))
}
