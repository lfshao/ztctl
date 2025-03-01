package output

import (
	"fmt"
	"strings"
)

// Color 定义输出颜色
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// Level 定义输出级别
type Level int

const (
	LevelInfo Level = iota
	LevelWarning
	LevelError
	LevelSuccess
)

// DefaultPrinter 默认的全局打印器实例
var DefaultPrinter Printer = NewConsolePrinter()

// Printer 定义输出接口
type Printer interface {
	PrintTable(headers []string, rows [][]string)
	PrintMsg(level Level, msg string, args ...interface{})
}

// ConsolePrinter 控制台输出实现
type ConsolePrinter struct{}

// NewConsolePrinter 创建一个新的控制台输出器
func NewConsolePrinter() Printer {
	return &ConsolePrinter{}
}

// PrintTable 打印表格
func (p *ConsolePrinter) PrintTable(headers []string, rows [][]string) {
	// 计算每列的最大宽度
	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h)
	}

	// 计算数据行中每列的最大宽度
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// 打印表头
	printRow(headers, colWidths)
	printSeparator(colWidths)

	// 打印数据行
	for _, row := range rows {
		printRow(row, colWidths)
	}
}

// PrintMsg 打印消息
func (p *ConsolePrinter) PrintMsg(level Level, msg string, args ...interface{}) {
	var color string
	switch level {
	case LevelWarning:
		color = ColorYellow
	case LevelError:
		color = ColorRed
	case LevelSuccess:
		color = ColorGreen
	default:
		color = ColorReset
	}

	fmt.Printf(color+msg+ColorReset+"\n", args...)
}

// 打印表格行
func printRow(cells []string, colWidths []int) {
	for i, cell := range cells {
		fmt.Printf("%-*s", colWidths[i]+2, cell)
	}
	fmt.Println()
}

// 打印分隔线
func printSeparator(colWidths []int) {
	for _, width := range colWidths {
		fmt.Print(strings.Repeat("-", width+2))
	}
	fmt.Println()
}
