package aarm

import "strconv"

// API é a estrutura que representa a api.
type API struct {
	Port int
}

// GetPortString retorna o valor da porta convertido em string.
func (api *API) GetPortString() string {
	return strconv.Itoa(api.Port)
}

// NewAPI retorna o potência de uma nova instância de uma estrutura API.
func NewAPI(port int) *API {
	return &API{
		Port: port,
	}
}
