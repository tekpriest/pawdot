package utils

import (
	"log"
	"os"
)

type Utils interface {
	AppError(err error, msg ...string)
}

func AppError(err error, msg ...string) {
	log.Fatalf("%s %s", err, msg)
	os.Exit(1)
}
