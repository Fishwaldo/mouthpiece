package telegram

import (
	"strings"
	"fmt"

	"github.com/go-logr/logr"
)


type telegramLogger struct {
	log logr.Logger
	token string
}

func (l *telegramLogger) Debug(msg ...interface{}) {
	l.log.V(1).Info(l.sanitize(fmt.Sprint(msg...)))
}

func (l *telegramLogger) Debugf(format string, args ...interface{}) {
	l.log.V(1).Info(l.sanitize(fmt.Sprintf(format, args...)))
}	

func (l *telegramLogger) Error(msg ...interface{}) {
	l.log.Error(fmt.Errorf(l.sanitize(fmt.Sprint(msg...))), "Telegram Error")	
}

func (l *telegramLogger) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Errorf(l.sanitize(fmt.Sprintf(format, args...))), "Telegram Error")
}

func (l *telegramLogger) sanitize(msg string) string {
	return strings.NewReplacer(l.token, "TOKEN").Replace(msg)
	
}