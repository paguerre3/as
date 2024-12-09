package domain

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	gryffindorCookieName = "gryffindor="
)

type GryffindorCookie interface {
	DecodeGryffindorCookie(setCookieHeaders []string) (string, error)
}

type gryffindorCookieImpl struct{}

func NewGryffindorCookie() GryffindorCookie {
	return &gryffindorCookieImpl{}
}

// Iterate through headers to find the "gryffindor" cookie
// Extract the gryffindor value
// Split by attributes
// Clean up spaces
// Remove quotes if present
// Decode the Base64 value
func (g *gryffindorCookieImpl) DecodeGryffindorCookie(setCookieHeaders []string) (string, error) {
	if setCookieHeaders == nil {
		return "", nil
	}
	for _, cookie := range setCookieHeaders {
		if strings.Contains(cookie, gryffindorCookieName) {

			parts := strings.Split(cookie, ";")
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if strings.HasPrefix(part, gryffindorCookieName) {
					encodedValue := strings.TrimPrefix(part, gryffindorCookieName)
					encodedValue = strings.Trim(encodedValue, `"`)

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
