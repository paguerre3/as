package infrastructure

import (
	"fmt"
	"testing"

	"github.com/paguerre3/as/internal/modules/5_valiant_last_stand/application"
	"github.com/stretchr/testify/assert"
)

// E5: La Última Defensa de la "Valiant" - ¡Cuenta Regresiva!
// Note: Basic prediction (simplified), final result is done via debug seen in the console.
/*
La Última Defensa de la "Valiant" - ¡Cuenta Regresiva!
Situación:

La nave insignia "Valiant", orgullo de la flota estelar, está bajo asedio. Una nave enemiga, veloz y letal, se aproxima con una intención clara: interceptar y destruir a la nave de suministros "Hope", que se encuentra en una posición vulnerable. El puesto de defensa de la "Valiant", crucial para repeler la amenaza, ha sufrido daños críticos en el ataque. La pantalla del radar, que normalmente mostraría la posición de la nave enemiga, es una maraña de estática y distorsiones. Sin embargo, el puerto de comunicaciones de la consola aún funciona, transmitiendo datos en texto sin formato. Afortunadamente, se ha encontrado un manual de instrucciones parcialmente legible que describe el formato de estos datos. Además, tienes acceso a la bitácora del operador anterior, donde encuentras la última lectura de la batalla anterior junto a una captura de pantalla del radar, proporcionando información valiosa sobre la situación inicial.

Misión:

Tu misión es crucial: detener a la nave enemiga antes de que alcance su objetivo. Dispones de un solo disparo, una descarga de energía concentrada capaz de aniquilar cualquier nave en un instante. El disparo debe ser preciso, dirigido a las coordenadas exactas (x, y) donde se encontrará la nave enemiga en el momento del impacto.

El Desafío:

La nave enemiga se mueve con una agilidad sorprendente, cambiando de dirección y velocidad de forma inesperada. Su objetivo es interceptar a la "Hope" lo antes posible, por lo que elegirá la ruta más directa y eficiente, evitando los obstáculos que se encuentran en el espacio de batalla. El espacio de acción de la torre de defensa es limitado y plagado de estos obstáculos. Tienes un total de 4 turnos para completar tu misión. En cada turno, puedes elegir entre dos acciones:

Leer el radar: Acceder al puerto de comunicaciones y obtener una lectura textual de la posición actual de la nave enemiga, en el formato descrito en el manual.
Atacar: Disparar tu arma a las coordenadas especificadas (x, y), teniendo en cuenta que la nave enemiga se habrá movido desde la última lectura del radar. Recuerda que solo tienes un disparo.
Cuenta Regresiva:

¡El tiempo es esencial! Desde el momento en que inicies la batalla llamando a la API de "start", tendrás solo 10 minutos para analizar los datos, predecir el movimiento de la nave enemiga considerando los obstáculos, y ejecutar el ataque.

Objetivo:

Escribe un programa que, bajo la presión del tiempo, analice los datos del puerto de comunicaciones en el formato especificado, utilice la información de la bitácora para comprender la situación inicial, prediga el movimiento de la nave enemiga considerando los obstáculos, y determine el momento y las coordenadas precisas para lanzar el ataque. ¡El destino de la "Hope" y de la "Valiant" está en tus manos!

Pistas:

El espacio de acción es una cuadrícula de 8x8, identificada desde "a1" hasta "h8".
La nave enemiga se representa con el carácter "^".
La nave amiga "Hope" se representa con el carácter "#".
Los obstáculos se representan con el carácter "$".
El espacio no ocupado se representa con el carácter "0".
Un salto de línea en los datos del radar se indica con el carácter "|".
La nave enemiga se mueve con un patrón peculiar, pero predecible si lo observas con atención.
Tienes 4 turnos.
Puedes "leer" o "atacar" en cada turno.
Solo tienes un disparo.
La nave enemiga busca la intercepción más rápida, evitando obstáculos.
Tienes 10 minutos para hacer tu disparo, luego de esto no podrás volver a intentar.
Tienes acceso a la última lectura del radar y una captura de pantalla de la batalla anterior en la bitácora.
*/
func TestLastDefenseSolution(t *testing.T) {
	battleActionsClient := NewBattleActionsClient()
	lastStandCounUseCase := application.NewLastStandCountdownUseCase(battleActionsClient)
	result, error := lastStandCounUseCase.Execute()
	assert.NoError(t, error)
	assert.NotEmpty(t, result)
	fmt.Printf("result: %+v\n", result)
}
