package utils

import (
	"fmt"
	"log"
	"strings"
)

func FatalOnError(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func LogErr(msg string, err error) error {
	if err != nil {

		msg = strings.ToLower(msg)

		log.Printf("%s: %s\n", msg, err)

		return fmt.Errorf("%s: %s", msg, err)
	}

	return nil
}
