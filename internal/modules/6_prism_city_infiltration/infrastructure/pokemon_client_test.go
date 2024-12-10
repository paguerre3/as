package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/modules/6_prism_city_infiltration/application"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// E6: La Infiltración en Ciudad Prisma: Un Desafío para los Maestros de Datos (Pokemons)
/*
La Infiltración en Ciudad Prisma: Un Desafío para los Maestros de Datos
Año 3036

En un rincón remoto del mundo Pokémon, Ciudad Prisma ha permanecido cerrada al acceso de entrenadores y Pokémon durante décadas. Se rumorea que en su interior se ocultan secretos ancestrales y un artefacto legendario que podría cambiar el destino de todas las regiones.

Sin embargo, la entrada a la ciudad está protegida por un complejo sistema de seguridad diseñado por los guardianes de Prisma, seres de inteligencia superior y maestros en el manejo de datos. Para demostrar que posees las habilidades necesarias para adentrarte en la ciudad, deberás superar su desafío definitivo.

El Desafío de los Guardianes
Los guardianes de Ciudad Prisma han recopilado una vasta cantidad de datos sobre todos los Pokémon existentes. Para desbloquear el acceso, deberás demostrar tu dominio en la manipulación de datos y tu conocimiento profundo del mundo Pokémon.

Tu Misión:

Calcular la altura promedio de todos los tipos de Pokémon, siguiendo el orden alfabético,
y enviar el valor con una precisión de 3 decimales.
*/
// integration test
func TestPokemonsHeightAvgSolution(t *testing.T) {
	pokemonClient := NewPokemonClient()
	calculatePokemonTypesAverageHeightsUseCase := application.NewCalculatePokemonTypesAverageHeightsUseCase(pokemonClient)
	response, statusCode, err := calculatePokemonTypesAverageHeightsUseCase.Execute()
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
