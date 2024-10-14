package utils

import (
	"fmt"
	"log"
)

func FatalOnError(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func HadleError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %s", msg, err)
	}
	return nil
}

func LogErr(msg string, err error) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}
