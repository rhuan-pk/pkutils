package adrm

import (
	"database/sql"
	"strconv"

	// username:password@(address)/databasename
	_ "github.com/go-sql-driver/mysql"
)

// NewDatabase retorna a instância de uma estrutura database.
func NewDatabase(user, password, host, name string, port int) (*Database, error) {

	// string de conexão.
	connection := user + ":" + password + "@(" + host + ":" + strconv.Itoa(port) + ")/" + name

	// retorna a representação do banco e valida erros de credênciais.
	representation, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	// valida se a conexão do banco pode ser estabelecida.
	err = representation.Ping()
	if err != nil {
		return nil, err
	}

	// caso tudo sucesse, retorne o ponteiro do banco e nil.
	return &Database{
		User:       user,
		Password:   password,
		Host:       host,
		Name:       name,
		Port:       port,
		Connection: representation,
	}, nil

}
