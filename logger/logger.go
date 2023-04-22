package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)
var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func init() {
	Info = log.New(os.Stdout, colorCyan+"INFO\t"+colorReset, log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, colorRed+"ERROR\t"+colorReset, log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrorNew(str string) {
	Error.Printf(fmt.Sprintf("%s%s%s", colorRed, str, colorReset))

}
func InfoNew(str string) {
	Info.Printf(fmt.Sprintf("%s%s%s", colorGreen, str, colorReset))
}
