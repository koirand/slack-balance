package utils

import (
	"log"
	"os"
)

func LoggingSetting() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}
