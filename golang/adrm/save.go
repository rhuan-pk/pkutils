package adrm

import (
	"log"
	"strings"
)

/*
Save insere ou atualizad um novo dado no banco retornando a quantidade de linhas modificadas caso seja um update ou a última linha inserida caso seja um create.

Informe a query completa a ser executada para esta função.
*/
func (database *Database) Save(query string) (int, error) {

	// pega as configuraçẽos do pacote, tenta abrir conexão com o banco, caso falhe retorne o erro se não continue.
	representation, err := database.openConnection()
	if err != nil {
		log.Println("estabilish connection failed:", err)
		return -1, err
	}
	defer representation.Close()
	log.Println("successfully estabilish connection!")

	// executa a query, caso falhe retorne o erro se não continue.
	result, err := representation.Exec(query)
	if err != nil {
		log.Println("query exec failed:", err, "query:", query)
		return -1, err
	}
	log.Println("successfully query exec!")

	// pega a quantidade de linhas modificadas ou a última inserida para usar como retorno.
	var rowsOrID int64
	if strings.Split(query, " ")[0] != "INSERT" {
		rowsOrID, _ = result.RowsAffected()
	} else {
		rowsOrID, _ = result.LastInsertId()
	}

	return int(rowsOrID), nil

}
