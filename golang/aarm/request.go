package aarm

import (
	"log"
	"net/http"
)

// StatusCodeReturn é a função comum para código de status de requisição
func StatusCodeReturn(endpoint string, response *http.ResponseWriter, httpStatusCode int) {
	level := map[bool]string{
		true:  "error",
		false: "hit",
	}[(httpStatusCode < 200 || httpStatusCode > 299)]
	responseMessage := http.StatusText(httpStatusCode)
	http.Error(*response, responseMessage, httpStatusCode)
	log.Println("endpoint \""+endpoint+"\" "+level+":", responseMessage+"!")
}
