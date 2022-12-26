package adrm

import (
	"log"
	"strings"
)

/*
Save insere ou atualizad um novo dado no banco retornando a quantidade de linhas modificadas caso seja um update ou a última linha inserida caso seja um create.

Informe a query completa a ser executada para esta função.
*/
func (database *Database) Save(query string, args ...any) (int, error) {

	// recebe somente o *sql.DB da representação do banco.
	representation := database.Connection

	// executa a query, caso falhe retorne o erro se não continue.
	result, err := representation.Exec(query, args...)
	if err != nil {
		log.Println("query exec failed!")
		return -1, err
	}

	// pega a quantidade de linhas modificadas ou a última inserida para usar como retorno.
	var rowsOrID int64
	if strings.Split(query, " ")[0] != "INSERT" {
		rowsOrID, _ = result.RowsAffected()
	} else {
		rowsOrID, _ = result.LastInsertId()
	}

	return int(rowsOrID), nil

}
