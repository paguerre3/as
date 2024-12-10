package infrastructure

import (
	"fmt"
)

const (
	AS_DOMAIN        = "https://makers-challenge.altscore.ai"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
	AUTHORIZATION    = "API-KEY"
	//API_KEY             = "36b39441ba7040fea158db6a4103aaa6" // hotmail2 (IMP: aguerrepablodario)
	//API_KEY             = "20adafe2bf3a4a43b20792aceac3dad9" // outlook (IMP: p_d_aguerre)
	API_KEY               = "7ae204f5e2fd450fa0e65e4fe58404bc" // sejax32599@pokeline.com
	SWAPI                 = "https://swapi.dev/api/"
	POKE_API              = "https://pokeapi.co/api/v2"
	TEMPLATES_DIR         = "internal/modules/7_9_damaged_spaceship/infrastructure/templates"
	EXPOSED_BASE_ENDPOINT = "https://didactic-rotary-phone-w4g77wq97p9hv6r-8080.app.github.dev"
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
