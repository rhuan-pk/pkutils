package adrm

// Database é a estrutra para manipulação de processos do banco de dados.
type Database struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// NewDatabase retorna a instância de uma estrutura database.
func NewDatabase(user, password, host, name string, port int) *Database {
	return &Database{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
	}
}
