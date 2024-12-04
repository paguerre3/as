package infrastructure

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	comm "github.com/paguerre3/as/internal/common"
)

const (
	gryffindorCookieName = "gryffindor="
)

func (c *clientHandlerImpl) OpenDoor(body interface{}, gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error) {
	uri := comm.BuildASApiUri(1, "s1/e8/actions/door")
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(body).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	if gryffindorCookies != nil {
		cookie, err := decodeGryffindorCookie(resp)
		if err != nil {
			return nil, resp.StatusCode(), err
		}
		if cookie != "" {
			*gryffindorCookies = append(*gryffindorCookies, cookie)
		}
	}
	return c.handleResponse(resp)
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
func (c *clientHandlerImpl) FirstClues() (response map[string]interface{}, statusCode int, err error) {
	j := 0
	gryffindorCookies := make([]string, 0)
	for {
		if j > 100 {
			break
		}
		r, s, e := c.OpenDoor(nil, &gryffindorCookies)
		if e != nil {
			fmt.Printf("firstClue %d error: %s\n", j, e.Error())
		}
		fmt.Printf("firstClue %d statusCode %d: %+v\n", j, s, r)
		if s == 200 && len(r) > 0 {
			resp, ok := r["response"]
			if ok && strings.Contains(resp.(string), "revelio") {
				fmt.Printf("attempting FourthClue ... %d\n", j)
				for {
					rr, ss, ee := c.FourthClue(&gryffindorCookies)
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
func (c *clientHandlerImpl) FourthClue(gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error) {
	var gryffindorCookiesStr string
	if len(*gryffindorCookies) > 0 {
		// Convert the slice to a single string with a delimiter
		gryffindorCookiesStr = strings.Join(*gryffindorCookies, " ")
	}
	// Altwarts revela cómo la magia surge mediante perseverancia, precisión y esmero al enfrentar desafíos. Cada detalle importa; auténtica destreza se refleja con dedicación para mejorar continuamente
	fmt.Printf("FourthClue gryffindorCookiesStr: %s\n", gryffindorCookiesStr)

	// At this point, the Revelio spell reveals what is hidden
	response, statusCode, err = c.OpenDoor(map[string]string{
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

func (c *clientHandlerImpl) HiddenMessageSolution(hiddenMessagePayload map[string]interface{}) (map[string]interface{}, int, error) {
	requestBody := hiddenMessagePayload
	uri := comm.BuildASApiUri(1, "s1/e8/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return c.handleResponse(resp)
}

func decodeGryffindorCookie(resp *resty.Response) (string, error) {
	if resp == nil {
		return "", fmt.Errorf("response is nil")
	}
	// Get all Set-Cookie headers
	setCookieHeaders := resp.Header().Values("Set-Cookie")

	// Iterate through headers to find the "gryffindor" cookie
	for _, cookie := range setCookieHeaders {
		if strings.Contains(cookie, gryffindorCookieName) {
			// Extract the gryffindor value
			parts := strings.Split(cookie, ";") // Split by attributes
			for _, part := range parts {
				part = strings.TrimSpace(part) // Clean up spaces
				if strings.HasPrefix(part, gryffindorCookieName) {
					encodedValue := strings.TrimPrefix(part, gryffindorCookieName)
					encodedValue = strings.Trim(encodedValue, `"`) // Remove quotes if present

					// Decode the Base64 value
					decodedValue, err := base64.StdEncoding.DecodeString(encodedValue)
					if err != nil {
						return "", fmt.Errorf("failed to decode Base64: %w", err)
					}

					str := string(decodedValue)
					fmt.Printf("Found gryffindor cookie: %s\n", str)
					return str, nil
				}
			}
		}
	}
	return "", nil
}
