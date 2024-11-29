package infrastructure

import "fmt"

const (
	AS_DOMAIN        = "https://makers-challenge.altscore.ai"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
	AUTHORIZATION    = "API-KEY"
	BEARER_API_KEY   = "a79f99a48ee04b529605b797fe43182c"
	SWAPI            = "https://swapi.dev/api/"
)

func buildASApiUri(version int, path string) string {
	return fmt.Sprintf("%s/v%d/%s", AS_DOMAIN, version, path)
}

func builSWAPIPeopleUri(path string, index int) string {
	return fmt.Sprintf("%s/%s/%d", SWAPI, path, index)
}
