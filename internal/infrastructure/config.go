package infrastructure

import "fmt"

const (
	AS_DOMAIN        = "https://makers-challenge.altscore.ai"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
	AUTHORIZATION    = "API-KEY"
	BEARER_API_KEY   = "a79f99a48ee04b529605b797fe43182c"
)

func buildApiUri(version int, path string) string {
	return fmt.Sprintf("%s/v%d/%s", AS_DOMAIN, version, path)
}
