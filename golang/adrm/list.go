package adrm

import (
	"log"
)

/*
ListRow retorna uma linha do banco de dados que é representada por um mapa no qual a chave é o nome da coluna e o seu valor é o valor da linha, caso algo falhe, retorne o erro.

Informe a query completa a ser executada para esta função que pode ou não conter parâmetros representados por "?".
*/
func (database *Database) ListRow(query string, args ...any) (map[string]string, error) {

	// recebe somente o *sql.DB da representação do banco.
	representation := database.Connection

	// faz preparação da query para dentro de stmt, caso falhe, retorne o erro.
	stmt, err := representation.Prepare(query)
	if err != nil {
		log.Println("query statement failed!")
		return nil, err
	}

	// cria o map de interface que será populado em bytes pela linha retornada do banco na execução da query, caso falhe, retorne o erro.
	rowInterfaceMap := make(map[any]any)
	err = stmt.QueryRow(args...).Scan(rowInterfaceMap)
	if err != nil {
		log.Println("error on read row!")
		return nil, err
	}

	// cria o mapa de colunas e valores que será populado com o nome da coluna e seu respectivo valor.
	var rowsMap map[string]string
	for columnName, value := range rowInterfaceMap {
		rowsMap[columnName.(string)] = func() string {
			if stringValue, ok := value.(string); ok {
				return string(stringValue)
			}
			return "null"
		}()
	}

	return rowsMap, nil

}

/*
List retorna múltiplas linhas do banco de dados que são representadas por uma slice de map sendo que cada índice da slice é uma linha do banco, a chave do mapa é o nome da coluna e seu valor é o valor da linha, caso algo falhe, retorne o erro.

Informe a query completa ser executada para esta função que pode ou não conter parâmetros representados por "?".
*/
func (database *Database) List(query string, args ...any) ([]map[string]string, error) {

	// recebe somente o *sql.DB da representação do banco.
	representation := database.Connection

	// faz preparação da query para dentro de stmt.
	stmt, err := representation.Prepare(query)
	if err != nil {
		log.Println("query statement failed!")
		return nil, err
	}

	// executa a query, caso falhe retorne o erro.
	table, err := stmt.Query(args...)
	if err != nil {
		log.Println("error query select!")
		return nil, err
	}
	defer table.Close()

	// cria a variável que guardará a slice de mapas que será retornada pelo função.
	var rowsMapSlice []map[string]string

	// pega o nome das colunas da tabela, caso falhe retorne o erro.
	columnsNames, err := table.Columns()
	if err != nil {
		log.Println("error get column names!")
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
			log.Println("error on read row!")
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
