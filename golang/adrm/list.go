package adrm

import (
	"log"
)

/*
List retorna a representação de uma tabela do banco sendo que cada índice da slice é uma linha que contem dentro um mapa no qual a chave é o nome da coluna e o seu valor é o valor da linha corrente na slice nessa chave do mapa que é coluna, caso algo falhe, retorne o erro.

Informe a query completa a ser executada para esta função.
*/
func (database *Database) List(query string) ([]map[string]string, error) {

	// recebe somente o *sql.DB da representação do banco.
	representation := database.Connection

	// executa a query, caso falhe retorne o erro.
	table, err := representation.Query(query)
	if err != nil {
		log.Println("query select error:", err, "query:", query)
		return nil, err
	}
	defer table.Close()
	log.Println("successfully query select!")

	// cria a variável que guardará a slice de rows/columns/values e que será retornada pelo função.
	var rowsMapSlice []map[string]string

	// pega o nome das colunas da tabela, caso falhe retorne o erro.
	columnsNames, err := table.Columns()
	if err != nil {
		log.Println("get column names error:", err)
		return nil, err
	}
	length := len(columnsNames)

	// cria a slice de interfaces que receberá os valores das linhas.
	rowInterfaceSlice := make([]any, length)
	for index := range rowInterfaceSlice {
		var rowPointer any
		rowInterfaceSlice[index] = &rowPointer
	}

	// itera sobre cada linha retornada do banco
	for table.Next() {

		// pega a linha da iteração atual e popula o valor de cada coluna na slice referente, caso falhe, retorne o erro.
		err := table.Scan(rowInterfaceSlice...)
		if err != nil {
			log.Println("error on read row:", err)
			return nil, err
		}

		// cria a slice temporária que será populada com o nome da coluna e seu respectivo valor com base na linha atual da iteração.
		rowMap := make(map[string]string, length)
		for index, columnName := range columnsNames {
			rowMap[columnName] = func() string {
				if rowByte, ok := (*rowInterfaceSlice[index].(*any)).([]byte); ok {
					return string(rowByte)
				}
				return "null"
			}()
		}

		// apenda a slice temporária de row na principal que será retornada.
		rowsMapSlice = append(rowsMapSlice, rowMap)

	}

	return rowsMapSlice, nil

}
