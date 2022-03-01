package logger

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

type CLILogger struct {
	logLevel int
	wherePad int
}

type CLILoggerExec struct {
	where     string
	wherePad  int
	level     int
	threshold int
}

var (
	Bold        = color.New(color.Bold).SprintfFunc()
	White       = color.New(color.FgWhite).SprintfFunc()
	WhiteBold   = color.New(color.FgWhite, color.Bold).SprintfFunc()
	Yellow      = color.New(color.FgYellow).SprintfFunc()
	YellowBold  = color.New(color.FgYellow, color.Bold).SprintfFunc()
	Green       = color.New(color.FgGreen).SprintfFunc()
	GreenBold   = color.New(color.FgGreen, color.Bold).SprintfFunc()
	Red         = color.New(color.FgRed).SprintfFunc()
	RedBold     = color.New(color.FgRed, color.Bold).SprintfFunc()
	Cyan        = color.New(color.FgCyan).SprintfFunc()
	CyanBold    = color.New(color.FgCyan, color.Bold).SprintfFunc()
	Magenta     = color.New(color.FgMagenta).SprintfFunc()
	MagentaBold = color.New(color.FgMagenta, color.Bold).SprintfFunc()
	Blue        = color.New(color.FgBlue).SprintfFunc()
	BlueBold    = color.New(color.FgBlue, color.Bold).SprintfFunc()
)

func NewCLILogger(logLevel int, wherePad int) *CLILogger {
	cliLogger := CLILogger{
		logLevel: logLevel,
		wherePad: wherePad,
	}
	return &cliLogger
}

func (c *CLILogger) Trace(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 6, wherePad: c.wherePad}
}

func (c *CLILogger) Debug(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 5, wherePad: c.wherePad}
}

func (c *CLILogger) Info(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 4, wherePad: c.wherePad}
}

func (c *CLILogger) Warn(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 3, wherePad: c.wherePad}
}

func (c *CLILogger) Error(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 2, wherePad: c.wherePad}
}

func (c *CLILogger) Fatal(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 1, wherePad: c.wherePad}
}

func (c *CLILogger) Panic(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: 0, wherePad: c.wherePad}
}

func (c *CLILogger) Highlight(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: -1, wherePad: c.wherePad}
}

func (c *CLILogger) Success(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: -2, wherePad: c.wherePad}
}

func (c *CLILogger) Announce(where string) *CLILoggerExec {
	return &CLILoggerExec{where: where, threshold: c.logLevel, level: -3, wherePad: c.wherePad}
}

func (c *CLILoggerExec) Prefix() string {
	whereString := c.where
	if len(whereString) > c.wherePad {
		whereString = whereString[:c.wherePad]
	}
	if len(whereString) < c.wherePad {
		for len(whereString) != c.wherePad {
			whereString = " " + whereString
		}
	}

	switch c.level {
	case 6: // TRACE
	case 5: // DEBUG
	case 4: // INFO
		whereString = Cyan(whereString)
	case 3: // WARN
		whereString = Yellow(whereString)
	case 2: // ERROR
		whereString = Red(whereString)
	case 1: // FATAL
		whereString = Red(whereString)
	case 0: // PANIC
		whereString = Red(whereString)
	case -1: // Highlight
		whereString = Green(whereString)
	case -2: // Success
		whereString = GreenBold(whereString)
	case -3: // Announce
		whereString = MagentaBold(whereString)
	default:
	}

	return fmt.Sprintf("%s ›", whereString)
}

func (c *CLILoggerExec) Msgf(format string, v ...interface{}) {
	switch {
	case c.level == -1: // Highlight
		fmt.Printf("%s %s\n", c.Prefix(), Green(format, v...))
	case c.level == -2: // Success
		fmt.Printf("%s %s\n", c.Prefix(), GreenBold(format, v...))
	case c.level == -3: // Announce
		fmt.Printf("%s %s\n", c.Prefix(), MagentaBold(format, v...))
	default:
		if c.threshold >= c.level {
			newFormat := fmt.Sprintf("%s %s\n", c.Prefix(), format)
			fmt.Printf(newFormat, v...)
		}
	}
}

func Dump(v ...interface{}) {
	spew.Dump(v...)
}
