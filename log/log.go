package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

//log package control , colorized information and log  values
var (
	Info  = log.New(os.Stdout, color.New(color.FgGreen, color.Bold).SprintFunc()("INFO: "), log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, color.New(color.FgRed).SprintFunc()("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(os.Stdout, color.New(color.FgRed, color.Bold).SprintFunc()("FATAL: "), log.Ldate|log.Ltime|log.Lshortfile)
)
