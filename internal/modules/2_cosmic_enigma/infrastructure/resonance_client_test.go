package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/modules/2_cosmic_enigma/application"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// E2: El Enigma Cósmico de Kepler-452b
/*
Año 3042: Eres un intrépido navegante a bordo del CSS Hawking, embarcado en una misión de vital importancia:
contactar al legendario Oráculo de Kepler-452b.
Se rumorea que este ser enigmático posee el conocimiento para guiar a la humanidad hacia una era dorada.
Pero el Oráculo no se revela a cualquiera; solo aquellos que demuestren su ingenio resolviendo un acertijo cósmico serán dignos de su sabiduría.

La Prueba: El Oráculo te presenta una interfaz holográfica que muestra una nebulosa estelar ondulante llamada " Lyra".
La interfaz te permite acceder a los datos de las estrellas en la nebulosa.
Para cada estrella, obtienes su " resonancia" y sus coordenadas.
El Oráculo te desafía a calcular la "resonancia promedio" de las estrellas en la nebulosa.
*/
// integration test
func TestFetchStarsAndResonanceSolution(t *testing.T) {
	resonanceClient := NewResonanceClient()
	averageResonanceUseCase := application.NewCalculateAverageResonanceUseCase(resonanceClient)
	response, statusCode, err := averageResonanceUseCase.Execute()
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
