package main

import (
	"log"
	"os"
)

func checkFileError(e error) {
	if e != nil && !os.IsNotExist(e) && !os.IsExist(e) {
		log.Fatal(e)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
