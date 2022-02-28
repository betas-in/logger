package logger

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type CLITable struct {
	headerFmt func(format string, a ...interface{}) string
	columnFmt func(format string, a ...interface{}) string
	table     table.Table
}

func NewCLITable(columns ...interface{}) *CLITable {
	t := CLITable{}
	t.headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
	t.columnFmt = color.New(color.FgYellow).SprintfFunc()
	t.table = table.New(columns...)
	t.table.WithHeaderFormatter(t.headerFmt).WithFirstColumnFormatter(t.columnFmt)
	return &t
}

func (t *CLITable) Row(vals ...interface{}) {
	t.table.AddRow(vals...)
}

func (t *CLITable) Display() {
	t.table.Print()
}
