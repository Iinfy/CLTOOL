package ui

import (
	"math"
	"strings"
)

func TableBuilder(columns []string, data [][]string) string {
	columnWidths := calcColumnsWidth(columns, data)
	lengthCells, widthCells := calcTableSize(columns, data)
	splitterLine := strings.Repeat("-", calcSpliterLineLength(columnWidths))
	table := strings.Repeat("_", calcSpliterLineLength(columnWidths))
	table += "\n"
	rowString := "|"
	for i, column := range columns {
		rowString += column
		rowString += strings.Repeat(" ", columnWidths[i]-len(column))
		rowString += "|"
	}
	rowString += "\n"
	table += rowString
	table += splitterLine
	table += "\n"
	for i := 0; i < lengthCells; i++ {
		rowString = "|"
		for j := 0; j < widthCells; j++ {
			rowString += data[i][j]
			rowString += strings.Repeat(" ", columnWidths[j]-len(data[i][j]))
			rowString += "|"
		}
		rowString += "\n"
		table += rowString
		rowString = ""
		table += splitterLine
		table += "\n"
	}
	return table
}

func calcTableSize(columns []string, data [][]string) (length int, width int) {
	length = len(data)
	width = len(columns)
	return length, width
}

func calcColumnsWidth(columns []string, data [][]string) []int {
	widths := make([]int, len(columns))
	for i, column := range columns {
		widths[i] = len(column)
	}
	for _, values := range data {
		for i, value := range values {
			widths[i] = int(math.Max(float64(widths[i]), float64(len(value))))
		}
	}
	return widths
}

func calcSpliterLineLength(columnsWidths []int) int {
	sumLength := 0
	for _, value := range columnsWidths {
		sumLength += value
	}
	return sumLength + len(columnsWidths) + 1
}
