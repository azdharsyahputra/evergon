package util

import "log"

func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Error(msg string, err error) {
	log.Println("[ERROR]", msg, err)
}
