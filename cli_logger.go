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
	// bold        = color.New(color.Bold).SprintfFunc()
	// white       = color.New(color.FgWhite).SprintfFunc()
	// whiteBold   = color.New(color.FgWhite, color.Bold).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	// yellowBold  = color.New(color.FgYellow, color.Bold).SprintfFunc()
	green     = color.New(color.FgGreen).SprintfFunc()
	greenBold = color.New(color.FgGreen, color.Bold).SprintfFunc()
	red       = color.New(color.FgRed).SprintfFunc()
	// redBold     = color.New(color.FgRed, color.Bold).SprintfFunc()
	cyan = color.New(color.FgCyan).SprintfFunc()
	// cyanBold    = color.New(color.FgCyan, color.Bold).SprintfFunc()
	// magenta     = color.New(color.FgMagenta).SprintfFunc()
	magentaBold = color.New(color.FgMagenta, color.Bold).SprintfFunc()
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
		whereString = cyan(whereString)
	case 3: // WARN
		whereString = yellow(whereString)
	case 2: // ERROR
		whereString = red(whereString)
	case 1: // FATAL
		whereString = red(whereString)
	case 0: // PANIC
		whereString = red(whereString)
	case -1: // Highlight
		whereString = green(whereString)
	case -2: // Success
		whereString = greenBold(whereString)
	case -3: // Announce
		whereString = magentaBold(whereString)
	default:
	}

	return fmt.Sprintf("%s ›", whereString)
}

func (c *CLILoggerExec) Msgf(format string, v ...interface{}) {
	switch {
	case c.level == -1: // Highlight
		fmt.Printf("%s %s\n", c.Prefix(), green(format, v...))
	case c.level == -2: // Success
		fmt.Printf("%s %s\n", c.Prefix(), greenBold(format, v...))
	case c.level == -3: // Announce
		fmt.Printf("%s %s\n", c.Prefix(), magentaBold(format, v...))
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
