package main

import (
	"log"
	"time"
)

const (
	CheckDuration = "24h"
)

func schedule(schedulerChan chan bool) {
	duration, err := time.ParseDuration(CheckDuration)
	if err != nil {
		log.Fatal(err)
	}

	for {
		schedulerChan <- true
		time.Sleep(duration)
	}
}
