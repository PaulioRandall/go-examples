package logfmt

import (
	"fmt"
	"log"
)

func Debug(msg string, args ...interface{}) string {
	return FmtMsg("debug", msg, args...)
}

func LogDebug(msg string, args ...interface{}) {
	m := Debug(msg, args...)
	log.Println(m)
}

func Info(msg string, args ...interface{}) string {
	return FmtMsg("info", msg, args...)
}

func LogInfo(msg string, args ...interface{}) {
	m := Info(msg, args...)
	log.Println(m)
}

func Error(msg string, args ...interface{}) string {
	return FmtMsg("error", msg, args...)
}

func LogError(msg string, args ...interface{}) {
	m := Error(msg, args...)
	log.Println(m)
}

func FmtMsg(tag string, msg string, args ...interface{}) string {
	return fmt.Sprintf("["+tag+"] "+msg, args...)
}
