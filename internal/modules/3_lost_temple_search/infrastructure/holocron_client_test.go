package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/modules/3_lost_temple_search/application"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// E3: La Búsqueda del Templo Sith Perdido
/*
Año Galáctico 34 DBY

La Resistencia ha interceptado un fragmento de un antiguo holocrón Sith que contiene la clave para localizar un templo Sith perdido, el cual se rumorea que guarda secretos poderosos.

Sin embargo, el fragmento está codificado con un complejo acertijo que solo un verdadero maestro de la Fuerza y los datos puede descifrar.

El futuro de la galaxia depende de tu habilidad para descifrar el mensaje y encontrar el templo antes que la Primera Orden.

El Desafío
El fragmento del holocrón, el cual se encuentra fuera de nuestra galaxia, contiene datos sobre varios lugares y personajes clave. Tu misión es analizar la forma de realizar la conexión con el holocrón y encontrar el único planeta con equilibrio en la Fuerza.

Investigando la librería del templo Jedi en Coruscant, un antiguo pasaje menciona:

...el "Índice de Balance de la Fuerza" (IBF) para un planeta específico, es una medida de la
influencia del Lado Luminoso y del Lado Oscuro de la Fuerza en ese planeta se calcula como:

IBF = ((Número de Personajes del Lado Luminoso) - (Número de Personajes del Lado Oscuro)) /
       (Total de Personajes en el Planeta)

El IBF te dará un valor entre -1 y 1, donde -1 significa dominio total del Lado Oscuro, 0 significa equilibrio, y
1 significa dominio total del Lado Luminoso.
*/
func TestBalancedPlanetSolution(t *testing.T) {
	holocronClient := NewHolocronClient()
	searchLostTempleUseCase := application.NewSearchLostTempleUseCase(holocronClient)
	response, statusCode, err := searchLostTempleUseCase.Execute()
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
