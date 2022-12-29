package adrm

import (
	"log"
	"strconv"
)

/*
ListRow retorna uma linha do banco de dados que é representada por um mapa no qual a chave é o nome da coluna e o seu valor é o valor da linha, caso algo falhe, retorne o erro.

Informe a query completa a ser executada para esta função que pode ou não conter parâmetros representados por "?".
*/
func (database *Database) ListRow(query string, args ...any) (map[string]string, error) {

	// row irá retornar somente o mapa da primeira linha da slice de linhas, caso falhe retorne o erro.
	sliceRow, err := database.List(query, args...)
	if err != nil {
		return nil, err
	}
	if len(sliceRow) > 0 {
		return sliceRow[0], nil
	}
	return nil, nil

}

/*
List retorna múltiplas linhas do banco de dados que são representadas por uma slice de mapa sendo que cada índice da slice é uma linha do banco, a chave do mapa é o nome da coluna e seu valor é o valor da linha, caso algo falhe, retorne o erro.

Informe a query completa ser executada para esta função que pode ou não conter parâmetros representados por "?".
*/
func (database *Database) List(query string, args ...any) ([]map[string]string, error) {

	// recebe somente o *sql.DB da representação do banco.
	representation := database.Connection

	// faz preparação da query para dentro de stmt seta defer para fechar a stmt.
	stmt, err := representation.Prepare(query)
	if err != nil {
		log.Println("query statement failed!")
		return nil, err
	}
	defer stmt.Close()

	// executa a query, caso falhe retorne o erro.
	table, err := stmt.Query(args...)
	if err != nil {
		log.Println("error query select!")
		return nil, err
	}
	defer table.Close()

	// cria a variável que guardará a slice de mapas que será retornada pelo função.
	var rowsMapSlice []map[string]string

	// pega o nome das colunas da tabela e pega a quantidade de colunas, caso falhe retorne o erro.
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
				row := *rowInterfaceSlice[index].(*any)
				if rowByte, isByte := row.([]byte); isByte {
					return string(rowByte)
				} else if rowInt, isInt := row.(int64); isInt {
					return strconv.Itoa(int(rowInt))
				}
				return "null"
			}()
		}

		// apenda a slice temporária de row na principal que será retornada.
		rowsMapSlice = append(rowsMapSlice, rowMap)

	}

	// caso tudo sucesse, retorna a slice de mapas das rows.
	return rowsMapSlice, nil

}
