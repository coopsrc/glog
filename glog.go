package glog

import (
	"fmt"
	"time"
)

const (
	color_black   = uint8(iota + 90)
	color_red
	color_green
	color_yellow
	color_blue
	color_magenta
	color_cyan
	color_white
)

const (
	verbose = "[VERB]"
	trace   = "[TRAC]"
	errors  = "[ERRO]"
	warn    = "[WARN]"
	info    = "[INFO]"
	debug   = "[DBUG]"
	assert  = "[ASST]"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13

func Verbose(tag string, format string, a ...interface{}) {

	level := formatLevel(verbose)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Trace(tag string, format string, a ...interface{}) {

	level := formatLevel(trace)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Error(tag string, format string, a ...interface{}) {

	level := formatLevel(errors)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Warn(tag string, format string, a ...interface{}) {

	level := formatLevel(warn)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Info(tag string, format string, a ...interface{}) {

	level := formatLevel(info)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Debug(tag string, format string, a ...interface{}) {

	level := formatLevel(debug)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

func Asset(tag string, format string, a ...interface{}) {

	level := formatLevel(assert)
	colorTag := formatTag(level, tag)

	fmt.Println(formatPrefix(colorTag), fmt.Sprintf(format, a...))
}

// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func formatLevel(level string) string {

	switch level {
	case verbose:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_black, level)
	case trace:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_cyan, level)
	case errors:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, level)
	case warn:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, level)
	case info:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, level)
	case debug:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, level)
	case assert:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, level)
	default:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_black, level)
	}

}

func formatTag(level string, tag string) string {
	return fmt.Sprintf("%s%s", level, tag)
}

func formatPrefix(colorTag string) string {
	return fmt.Sprintf("%s%s:", time.Now().Format("2006-01-02 15:04:05.999"), colorTag)
}
