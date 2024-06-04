package model

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

// Printer 结构
type Printer struct {
	blowColor         *color.Color
	authColor         *color.Color
	pollReportColor   *color.Color
	logReportColor    *color.Color
	statusReportColor *color.Color
}

var PrinterInstance *Printer
var PrinterOnce sync.Once

func GetPrinterInstance() *Printer {
	PrinterOnce.Do(func() {
		PrinterInstance = &Printer{
			blowColor:         color.New(color.FgGreen).Add(color.Bold),
			authColor:         color.New(color.FgBlue).Add(color.Bold),
			pollReportColor:   color.New(color.FgYellow).Add(color.Bold),
			logReportColor:    color.New(color.FgRed).Add(color.Bold),
			statusReportColor: color.New(color.FgCyan).Add(color.Bold),
		}
	})
	return PrinterInstance
}

// Print 打印彩色信息
func (p *Printer) Print(category, roomId, message string) {
	var colorPrinter *color.Color

	switch category {
	case "blow":
		colorPrinter = p.blowColor
	case "auth":
		colorPrinter = p.authColor
	case "pollReport":
		colorPrinter = p.pollReportColor
	case "logReport":
		colorPrinter = p.logReportColor
	case "statusReport":
		colorPrinter = p.statusReportColor
	default:
		colorPrinter = color.New(color.FgWhite)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	colorPrinter.Printf("[%s] [%s] [Room %s]: %s\n", timestamp, category, roomId, message)
}
