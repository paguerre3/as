package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/modules/1_stealth_probe/application"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// E1: La Sonda Silenciosa
/*
Misión: Eres un intrépido explorador estelar en una misión crucial para mapear un sistema solar recién descubierto. Tu objetivo es determinar la velocidad orbital instantánea de un planeta potencialmente habitable para evaluar su idoneidad para la vida.

Desafío: La extraña interferencia cósmica en esta región del espacio dificulta la obtención de lecturas exitosas de tu escáner de largo alcance.

Datos Clave: Cuando el escáner funciona, te proporciona:

distance: La distancia recorrida por el planeta en su órbita durante el período de observación (en unidades astronómicas).
time: El tiempo transcurrido durante la observación (en horas).
Objetivo: Calcular la velocidad orbital instantánea del planeta hasta el número entero más cercano.
*/
// integration test
func TestStealthProbeSolution(t *testing.T) {
	measurmentClient := NewMeasurementClient()
	probeUseCase := application.NewProbeUseCase(measurmentClient)
	response, statusCode, err := probeUseCase.Execute()
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
