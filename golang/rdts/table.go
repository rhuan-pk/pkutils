package rdts

// Table é a estrutura que guardará as informações na representação do banco.
type Table struct {
	Name    string            `json:"name"`
	Columns []string          `json:"columns"`
	Rows    []*Row            `json:"rows"`
	Options map[string]string `json:"options"`
}

// NewTable retorna a instânci de uma estrutura Table.
func NewTable() *Table {
	return &Table{}
}
