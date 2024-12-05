package common

import (
	"fmt"
)

const (
	AS_DOMAIN        = "https://makers-challenge.altscore.ai"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
	AUTHORIZATION    = "API-KEY"
	// original mail
	//BEARER_API_KEY = "80ea79f02e684f9eb5d979d9f09ba087" // hotmail (p_aguerre)
	//BEARER_API_KEY = "a79f99a48ee04b529605b797fe43182c" // gmail1 (pablo.aguerre)
	//BEARER_API_KEY = "255292ff68394c6eb6136069a034bf28" // gmail2 (IMPORTANT: aguerrepablodario)
	BEARER_API_KEY        = "e14b6ffe60a64994b923f31d088ef983" // hotmail2 (aguerre_p_d)
	SWAPI                 = "https://swapi.dev/api/"
	POKE_API              = "https://pokeapi.co/api/v2"
	TEMPLATES_DIR         = "internal/infrastructure/templates"
	EXPOSED_BASE_ENDPOINT = "https://studious-journey-7xp99gjrvjrfxp77-8080.app.github.dev"
)

func BuildASApiUri(version int, path string) string {
	return fmt.Sprintf("%s/v%d/%s", AS_DOMAIN, version, path)
}

func BuilSWAPIPeopleUri(path string, index int) string {
	return fmt.Sprintf("%s/%s/%d", SWAPI, path, index)
}

func BuildPockeApi(path string) string {
	return fmt.Sprintf("%s/%s", POKE_API, path)
}
