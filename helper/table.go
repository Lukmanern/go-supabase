package helper

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// id todo (s:status) (d:deleted_at)

func GetTableStructure() *tabwriter.Writer {
	table := tabwriter.NewWriter(os.Stdout, 8, 8, 2, '\t', tabwriter.AlignRight)
	fmt.Fprintln(table, "ID\tTodo\tStatus\tDeleted at")
	fmt.Fprintln(table, "--\t----\t------\t----------")

	return table
}