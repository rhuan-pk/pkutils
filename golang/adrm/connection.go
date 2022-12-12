package adrm

import (
	"database/sql"
	"log"
	"strconv"

	// username:password@(address)/databasename
	_ "github.com/go-sql-driver/mysql"
)

// openConnection abre uma conexão com o banco e valida a mesma, caso tudo sucesse retorna a representação do banco, caso contrário retorne o erro referente.
func (database *Database) openConnection() (*sql.DB, error) {

	// variáveis que guarda as informações das configurações de credencial.
	user := database.User
	password := database.Password
	host := database.Host
	port := strconv.Itoa(database.Port)
	name := database.Name

	// string de conexão.
	connection := user + ":" + password + "@(" + host + ":" + port + ")/" + name

	// pega a representação do banco e valida erros de credênciais.
	representation, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("open connection failed:", err)
		return nil, err
	}
	log.Println("successfully connection!")

	// pega a conexão e valida erros da mesma.
	err = representation.Ping()
	if err != nil {
		log.Println("ping connection failed:", err)
		return nil, err
	}
	log.Println("successfully ping!")

	// caso tudo sucesse, retorna o banco e nil.
	return representation, nil

}
