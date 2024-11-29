package domain

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
}
