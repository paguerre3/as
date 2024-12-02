package domain

import "sync"

type ClientHandler interface {
	Register(alias, country, email, applyRole string) (response map[string]interface{}, statusCode int, err error)

	Measurement() (response map[string]interface{}, statusCode int, err error)
	MeassurmentSolution(speed int) (response map[string]interface{}, statusCode int, err error)

	FetchStars(page int) (response []map[string]interface{}, statusCode int, err error)
	ResonanceSolution(averaegeResonance int) (response map[string]interface{}, statusCode int, err error)

	Fetch(uri string) (response map[string]interface{}, statusCode int, err error)
	FetchSWAPIPlanets(index int) (response map[string]interface{}, statusCode int, err error)
	QueryOracle(name string) (response map[string]interface{}, statusCode int, err error)
	OracleSolution(balancedBlanet string) (response map[string]interface{}, statusCode int, err error)

	UserAndPasswordSolution(username, password string) (response map[string]interface{}, statusCode int, err error)

	StartBattle() (response string, statusCode int, err error)
	PerformTurn(action string, x string, y int) (response map[string]interface{}, statusCode int, err error)

	GetPockemonTypes() (response map[string]interface{}, statusCode int, err error)
	GetTypeData(typeUrl, typeName string) (response map[string]interface{}, statusCode int, err error)
	GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error)
	PokemonSolution(pokeSolution map[string]interface{}) (response map[string]interface{}, statusCode int, err error)
}
