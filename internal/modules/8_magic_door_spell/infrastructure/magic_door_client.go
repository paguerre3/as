package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type MagicDoorClient interface {
	FirstClues() (response map[string]interface{}, statusCode int, err error)
	HiddenMessageSolution(hiddenMessagePayload map[string]interface{}) (map[string]interface{}, int, error)
}

type GryffindorCookie interface {
	DecodeGryffindorCookie(setCookieHeaders []string) (string, error)
} // GryffindorCookie helper interface duplicaed from domain in order to avoid DDD violation

type magicDoorClientImpl struct {
	clientHandler    common_infra.ClientHandler
	griffindorCookie GryffindorCookie
}

func NewMagicDoorClient(gryffindorCookie GryffindorCookie) MagicDoorClient {
	return &magicDoorClientImpl{
		clientHandler:    common_infra.NewClientHandlerDebug(), // Must be in debug mode for logging cookie headers!
		griffindorCookie: gryffindorCookie,
	}
}

func (c *magicDoorClientImpl) openDoor(body interface{}, gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error) {
	uri := common_infra.BuildASApiUri(1, "s1/e8/actions/door")
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(body).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	if gryffindorCookies != nil {
		cookie, err := c.decodeGryffindorCookie(resp)
		if err != nil {
			return nil, resp.StatusCode(), err
		}
		if cookie != "" {
			*gryffindorCookies = append(*gryffindorCookies, cookie)
		}
	}
	return c.clientHandler.HandleResponse(resp)
}

// First clue: The Second Clock
/*
Primera pista: El Reloj de los Segundos

En la Torre del Reloj de Altwarts, el tiempo es tu mayor aliado y tu peor enemigo.
Cada segundo marca el paso de una clave que abre una puerta.
Elige el momento correcto, y la llave será tuya.
Pero recuerda, solo en el primer segundo puedes encontrar la llave de la primera puerta.
*/
// Second clue: Action and Reaction
/*
Segunda pista: Acción y reacción, las puertas traen revelación

En el aula de Encantamientos, el profesor Flitwick guarda un pergamino antiguo que habla de palabras ocultas,
no en los hechizos mismos, sino en las consecuencias de su lanzamiento.
Cada acción mágica genera una reacción, y en esa reacción yacen los secretos que buscan los magos astutos.
Solo aquellos que observan más allá de lo evidente, que desentrañan los hilos invisibles que conectan causa y efecto,
podrán descifrar el verdadero mensaje.
*/

// Third clue: N magical doors
/*
Tercera pista:

En el Gran Comedor, N puertas mágicas están alineadas, cada una desbloqueada solo por una palabra.
Un hechizo pronunciado mal te perderá, pero el correcto te guiará.
Solo avanzando en orden y en el momento adecuado, alcanzarás tu objetivo.
*/
func (c *magicDoorClientImpl) FirstClues() (response map[string]interface{}, statusCode int, err error) {
	j := 0
	gryffindorCookies := make([]string, 0)
	for {
		if j > 100 {
			break
		}
		r, s, e := c.openDoor(nil, &gryffindorCookies)
		if e != nil {
			fmt.Printf("firstClue %d error: %s\n", j, e.Error())
		}
		fmt.Printf("firstClue %d statusCode %d: %+v\n", j, s, r)
		if s == 200 && len(r) > 0 {
			resp, ok := r["response"]
			if ok && strings.Contains(resp.(string), "revelio") {
				fmt.Printf("attempting FourthClue ... %d\n", j)
				for {
					rr, ss, ee := c.fourthClue(&gryffindorCookies)
					if ee != nil {
						fmt.Printf("fourthClue %d error: %s\n", j, ee.Error())
					}
					hm, ok := rr["hidden_message"]
					if !ok {
						hm = ""
					}
					if ss == 200 && len(resp.(string)) > 0 && len(hm.(string)) > 0 {
						return rr, ss, nil
					}
				}
			}
		}
		time.Sleep(1 * time.Second)
		j++
	}
	return nil, 0, fmt.Errorf("could not open the door")
}

// Fourth clue: The Revelio spell
/*
Cuarta pista:

Cada puerta te llevará a la siguiente, pero la respuesta no se encuentra al final,
sino que el "camino mismo" es la clave.
Recuerda usar el hechizo Revelio para ver lo que está oculto.
*/
func (c *magicDoorClientImpl) fourthClue(gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error) {
	var gryffindorCookiesStr string
	if len(*gryffindorCookies) > 0 {
		// Convert the slice to a single string with a delimiter
		gryffindorCookiesStr = strings.Join(*gryffindorCookies, " ")
	}
	// Altwarts revela cómo la magia surge mediante perseverancia, precisión y esmero al enfrentar desafíos. Cada detalle importa; auténtica destreza se refleja con dedicación para mejorar continuamente
	fmt.Printf("FourthClue gryffindorCookiesStr: %s\n", gryffindorCookiesStr)

	// At this point, the Revelio spell reveals what is hidden
	response, statusCode, err = c.openDoor(map[string]string{
		"revelio": gryffindorCookiesStr,
	}, nil)
	if err != nil {
		fmt.Printf("FourthClue error: %s\n", err.Error())
		return nil, statusCode, err
	}
	fmt.Printf("FourthClue statusCode %d: %+v\n", statusCode, response)
	if statusCode == 200 && len(gryffindorCookiesStr) > 0 {
		// override response:
		response = map[string]interface{}{
			"hidden_message": gryffindorCookiesStr,
		}
	}
	return response, statusCode, nil
}

func (c *magicDoorClientImpl) HiddenMessageSolution(hiddenMessagePayload map[string]interface{}) (map[string]interface{}, int, error) {
	requestBody := hiddenMessagePayload
	uri := common_infra.BuildASApiUri(1, "s1/e8/solution")

	// Send the POST request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}

func (c *magicDoorClientImpl) decodeGryffindorCookie(resp *resty.Response) (string, error) {
	if resp == nil {
		return "", fmt.Errorf("response is nil")
	}
	// Get all Set-Cookie headers
	setCookieHeaders := resp.Header().Values("Set-Cookie")
	return c.griffindorCookie.DecodeGryffindorCookie(setCookieHeaders)
}
