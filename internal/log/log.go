package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type LogLevel string

const (
	ERROR   LogLevel = "ERROR"
	INFO    LogLevel = "INFO"
	SUCCESS LogLevel = "SUCCESS"
	WARNING LogLevel = "WARNING"
)

func log(message string, level LogLevel) {
	var color string
	switch level {
	case ERROR:
		color = "#FF0000"
	case INFO:
		color = "#00FF00"
	case SUCCESS:
		color = "#00FF00"
	case WARNING:
		color = "#FFCC66"
	}

	levelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(color)).
		Render(string(level))

	fmt.Println(lipgloss.NewStyle().Render(levelStyle, message))
}

func Info(message string) {
	log(message, INFO)
}

func Error(message string, errValue error) {
	var errMessage string
	if errValue != nil {
		errMessage = fmt.Sprintf("%s: %e", message, errValue)
		log(errMessage, ERROR)
	}
	log(message, ERROR)
}

func Warning(message string) {
	log(message, WARNING)
}

func Success(message string) {
	log(message, SUCCESS)
}
