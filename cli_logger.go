package logger

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

var cliLogLevel = 3

func NewCLILogger(logLevel string) {
	if logLevel == "" {
		cliLogLevel = 3
		return
	}

	switch logLevel {
	case "TRACE":
		cliLogLevel = 6
	case "DEBUG":
		cliLogLevel = 5
	case "INFO":
		cliLogLevel = 4
	case "WARN":
		cliLogLevel = 3
	case "ERROR":
		cliLogLevel = 2
	case "FATAL":
		cliLogLevel = 1
	case "PANIC":
		cliLogLevel = 0
	}
}

func Tracef(format string, v ...interface{}) {
	if cliLogLevel >= 6 {
		fmt.Printf(format+"\n", v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if cliLogLevel >= 5 {
		fmt.Printf(format+"\n", v...)
	}
}

func Infof(format string, v ...interface{}) {
	if cliLogLevel >= 4 {
		code := color.New(color.FgCyan)
		code.Printf(format+"\n", v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if cliLogLevel >= 3 {
		code := color.New(color.FgYellow, color.Bold)
		code.Printf(format+"\n", v...)
	}
}

func MinorSuccessf(format string, v ...interface{}) {
	if cliLogLevel >= 2 {
		code := color.New(color.FgGreen)
		code.Printf(format+"\n", v...)
	}
}

func Successf(format string, v ...interface{}) {
	if cliLogLevel >= 2 {
		code := color.New(color.FgGreen, color.Bold)
		code.Printf(format+"\n", v...)
	}
}

func Announcef(format string, v ...interface{}) {
	if cliLogLevel >= 2 {
		code := color.New(color.FgCyan, color.Bold)
		code.Printf(format+"\n", v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if cliLogLevel >= 2 {
		code := color.New(color.FgRed, color.Bold)
		code.Printf(format+"\n", v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if cliLogLevel >= 1 {
		code := color.New(color.FgRed, color.Bold, color.Underline)
		code.Printf(format+"\n", v...)
	}
}

func Panicf(format string, v ...interface{}) {
	if cliLogLevel >= 0 {
		code := color.New(color.FgRed, color.Bold, color.Underline)
		code.Printf(format+"\n", v...)
	}
}

func Questionf(format string, v ...interface{}) string {
	reader := bufio.NewReader(os.Stdin)

	code := color.New(color.FgMagenta, color.Bold)
	code.Printf(format+" : ", v...)
	text, err := reader.ReadString('\n')
	if err != nil {
		Errorf("could not read input: %v", err)
	}
	return text
}

func Dump(v ...interface{}) {
	spew.Dump(v...)
}
