package adrm

import "database/sql"

// Database é a estrutra para manipulação de processos do banco de dados.
type Database struct {
	User       string
	Password   string
	Host       string
	Name       string
	Port       int
	Connection *sql.DB
}
