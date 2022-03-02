package mberror

import "log"

type ErrorLevel int

const (
	Debug ErrorLevel = iota
	Info
	Warn
	Error
	Panic
	Fatal
)

func Check(err error) bool {
	return CheckLevel(err, Info)
}

func CheckLevel(err error, level ErrorLevel) bool {
	if err != nil {
		log.Println(err.Error())
		if level >= Error {
			return true
		}
	}
	return false
}
