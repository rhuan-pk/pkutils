package aarm

import (
	"log"
	"net/http"
)

// StatusCodeReturn é a função comum para código de status de requisição
func StatusCodeReturn(endpoint string, writer *http.ResponseWriter, httpStatusCode int) {
	level := map[bool]string{true: "error", false: "hit"}[httpStatusCode != http.StatusOK]
	httpErrorMessage := http.StatusText(httpStatusCode)
	http.Error(*writer, httpErrorMessage, httpStatusCode)
	log.Println("endpoint \""+endpoint+"\" "+level+":", httpErrorMessage+"!")
}
