package loggerinfo

import (
	"fmt"
)

func Logger(message string) {
	fmt.Println("-------------------- logger -------------------- \n", message)
}

func LoggerError(err error) {
	fmt.Println("-------------------- err msg -------------------- \n", err)
}
