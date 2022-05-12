package logger

import (
	"log"
)

func Info(message string) {
	log.Println("-------------------- Info -------------------- \n", message)
}

func Error(err error) {
	log.Println("-------------------- Err -------------------- \n", err)
}
