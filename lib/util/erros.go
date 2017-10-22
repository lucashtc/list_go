package util

import (
	"log"
)

func VerifyErr(err error) {
	if err != nil {
		log.Fatalf("Error ===> %s \n", err)
	}
}