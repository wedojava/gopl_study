// Package console provides operations for the console.
package console

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"
)

// Clear clears the console.
func Clear() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// SprintTable takes a slice of a table row and returns a string representation.
func SprintTable(rows [][]string) string {
	if len(rows) == 0 || len(rows[0]) == 0 {
		return ""
	}
	var b bytes.Buffer
	lens := columnLens(rows)
	for i := range rows[0] {
		fmt.Fprintf(&b, "+-%s-", strings.Repeat("-", lens[i]))
	}
	fmt.Fprintf(&b, "+\n")

	for _, row := range rows {
		for i := range rows[0] {
			var val string
			if len(row) > i {
				val = row[i]
			}
			fmt.Fprintf(&b, "| % -*s ", lens[i], val)
		}
		fmt.Fprintf(&b, "|\n")
		for i := range rows[0] {
			fmt.Fprintf(&b, "+-%s-", strings.Repeat("-", lens[i]))
		}
		fmt.Fprintf(&b, "+\n")
	}
	return b.String()
}

// columnLens returns the maximum string length for each column.
func columnLens(rows [][]string) []int {
	result := []int{}
	if len(rows) == 0 || len(rows[0]) == 0 {
		return result
	}
	for i := range rows[0] {
		max := 0
		for _, row := range rows {
			if len(row) > i {
				val := utf8.RuneCountInString(row[i])
				if max < val {
					max = val
				}
			}
		}
		result = append(result, max)
	}
	return result
}
