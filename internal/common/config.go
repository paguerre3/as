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
	//BEARER_API_KEY = "80ea79f02e684f9eb5d979d9f09ba087" // hotmail
	//BEARER_API_KEY = "a79f99a48ee04b529605b797fe43182c" // gmail1
	BEARER_API_KEY = "255292ff68394c6eb6136069a034bf28" // gmail2
	SWAPI          = "https://swapi.dev/api/"
	POKE_API       = "https://pokeapi.co/api/v2"
	TEMPLATES_DIR  = "internal/infrastructure/templates"
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
