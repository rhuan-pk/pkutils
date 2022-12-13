package adrm

import "database/sql"

// Database é a estrutra para manipulação de processos do banco de dados.
type Database struct {
	Connection *sql.DB
}
