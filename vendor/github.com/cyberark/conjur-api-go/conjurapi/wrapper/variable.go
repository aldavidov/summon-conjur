package wrapper

import (
	"net/url"
	"fmt"
	"net/http"
	"strings"
)

func RetrieveSecretRequest(applianceURL, account, variableIdentifier string) (*http.Request, error) {
	return http.NewRequest(
		"GET",
		VariableURL(applianceURL, account, variableIdentifier),
		nil,
	)
}

func AddSecretRequest(applianceURL, account, variableIdentifier, secretValue string) (*http.Request, error) {
	return http.NewRequest(
		"POST",
		VariableURL(applianceURL, account, variableIdentifier),
		strings.NewReader(secretValue),
	)
}

func RetrieveSecretResponse(variableIdentifier string, resp *http.Response) ([]byte, error) {
	err := VariableResponse(variableIdentifier, resp)
	if err != nil {
		return nil, err
	}
	return ByteResponseTransformer(resp)
}

func AddSecretResponse(variableIdentifier string, resp *http.Response) error {
	return VariableResponse(variableIdentifier, resp)
}

func VariableResponse(variableIdentifier string, resp *http.Response) error {
	switch resp.StatusCode {
	case 404:
		return fmt.Errorf("%v: Variable '%s' not found", resp.StatusCode, variableIdentifier)
	case 403:
		return fmt.Errorf("%v: Invalid permissions on '%s'", resp.StatusCode, variableIdentifier)
	case 200, 201:
		return nil
	default:
		return fmt.Errorf("%v: %s", resp.StatusCode, resp.Status)
	}
}

func VariableURL(applianceURL, account, variableIdentifier string) string {
	return fmt.Sprintf("%s/secrets/%s/variable/%s", applianceURL, account, url.QueryEscape(variableIdentifier))
}
