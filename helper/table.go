package helper

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// id todo (s:status) (d:deleted_at)

func GetTableStructure() *tabwriter.Writer {
	// see https://pkg.go.dev/text/tabwriter#NewWriter
	table := tabwriter.NewWriter(os.Stdout, 8, 8, 2, '\t', tabwriter.AlignRight)
	fmt.Fprintln(table, "ID\tTodo\tStatus\tDeleted at")
	fmt.Fprintln(table, "--\t----\t------\t----------")

	return table
}