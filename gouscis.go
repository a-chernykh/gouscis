package main

import (
	"fmt"
	"log"
	"os"
)

const (
	DatabasePath = "cases"
)

func main() {
	requiredVars := []string{"CASE_NUMBER", "EMAIL", "SMTP_SERVER"}
	for _, key := range requiredVars {
		val := os.Getenv(key)
		if val == "" {
			log.Fatal(fmt.Sprintf("Please specify %s environment variable", key))
		}
	}

	db := Database{Path: DatabasePath}
	db.create()

	schedulerChan := make(chan bool)

	notifier := EmailNotifier{Smtp: os.Getenv("SMTP_SERVER"),
		Sender:    "noreply@gouscis",
		Recipient: os.Getenv("EMAIL")}

	go schedule(schedulerChan)
	statusUpdate(schedulerChan, os.Getenv("CASE_NUMBER"), db, notifier)
}
