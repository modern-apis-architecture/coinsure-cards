package banklo

import "net/http"

func ProvideHttpCli() *http.Client {
	return &http.Client{}
}
