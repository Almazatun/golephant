package logger

import (
	"fmt"
)

func Info(message string) {
	fmt.Println("-------------------- logger -------------------- \n", message)
}

func InfoError(err error) {
	fmt.Println("-------------------- err msg -------------------- \n", err)
}
