package application

type MagicDoorUseCase interface {
	Execute() (map[string]interface{}, int, error)
}

type MagicDoorClient interface {
	FirstClues() (response map[string]interface{}, statusCode int, err error)
	HiddenMessageSolution(hiddenMessagePayload map[string]interface{}) (map[string]interface{}, int, error)
} // Exposing duplicated interface to avoid DDD violations

type magicDoorUseCaseImpl struct {
	magicDoorClient MagicDoorClient
}

func NewMagicDoorUseCase(magicDoorClient MagicDoorClient) MagicDoorUseCase {
	return &magicDoorUseCaseImpl{
		magicDoorClient: magicDoorClient,
	}
}

func (c *magicDoorUseCaseImpl) Execute() (map[string]interface{}, int, error) {
	response, statusCode, err := c.magicDoorClient.FirstClues()
	if err != nil {
		return nil, statusCode, err
	}
	return c.magicDoorClient.HiddenMessageSolution(response)
}
