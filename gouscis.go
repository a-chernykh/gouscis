package main

import (
	"fmt"
	"log"
	"os"
)

const (
	databasePath = "cases"
)

func main() {
	requiredVars := []string{"CASE_NUMBER", "EMAIL", "SMTP_SERVER"}
	for _, key := range requiredVars {
		val := os.Getenv(key)
		if val == "" {
			log.Fatal(fmt.Sprintf("Please specify %s environment variable", key))
		}
	}

	db := database{Path: databasePath}
	db.create()

	schedulerChan := make(chan bool)

	notifier := emailNotifier{SMTP: os.Getenv("SMTP_SERVER"),
		Sender:    "noreply@gouscis",
		Recipient: os.Getenv("EMAIL")}

	go schedule(schedulerChan)
	statusUpdate(schedulerChan, os.Getenv("CASE_NUMBER"), db, notifier)
}
