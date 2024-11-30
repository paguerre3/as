package infrastructure

import (
	"fmt"
)

const (
	AS_DOMAIN        = "https://makers-challenge.altscore.ai"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
	AUTHORIZATION    = "API-KEY"
	// original mail API KEY BEARER_API_KEY = "a79f99a48ee04b529605b797fe43182c"
	BEARER_API_KEY = "255292ff68394c6eb6136069a034bf28"
	SWAPI          = "https://swapi.dev/api/"
	POKE_API       = "https://pokeapi.co/api/v2"
)

func buildASApiUri(version int, path string) string {
	return fmt.Sprintf("%s/v%d/%s", AS_DOMAIN, version, path)
}

func builSWAPIPeopleUri(path string, index int) string {
	return fmt.Sprintf("%s/%s/%d", SWAPI, path, index)
}

func buildPockeApi(path string) string {
	return fmt.Sprintf("%s/%s", POKE_API, path)
}
