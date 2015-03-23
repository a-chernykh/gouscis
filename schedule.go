package main

import (
	"log"
	"time"
)

const (
	checkDuration = "24h"
)

func schedule(schedulerChan chan bool) {
	duration, err := time.ParseDuration(checkDuration)
	if err != nil {
		log.Fatal(err)
	}

	for {
		schedulerChan <- true
		time.Sleep(duration)
	}
}
