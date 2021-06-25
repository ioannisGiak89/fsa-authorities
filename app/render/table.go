package render

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Table represents a standard table
type Table struct {
	rows      []table.Row
	header    table.Row
	subHeader table.Row
}

// NewTable creates a new Table
func NewTable(rows []table.Row, header table.Row, subHeader table.Row) *Table {
	return &Table{rows: rows, header: header, subHeader: subHeader}
}

// Render renders a table
func (t *Table) Render() {
	pt := table.NewWriter()
	pt.SetOutputMirror(os.Stdout)
	pt.AppendHeader(t.header, table.RowConfig{AutoMerge: true})
	pt.AppendHeader(t.subHeader)
	pt.AppendRows(t.rows)
	pt.AppendSeparator()
	pt.Render()
}
