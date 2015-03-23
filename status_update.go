package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	caseStatusURL = "https://egov.uscis.gov/casestatus/mycasestatus.do"

	caseStatusContainerSelector   = ".main-row"
	caseStatusSelector            = "h1"
	caseStatusDescriptionSelector = ".rows p"
)

func statusUpdate(schedulerChan chan bool, caseNumber string, db database, notifier notifier) {
	previousCaseStatus := db.loadCaseStatus(caseNumber)

	for {
		<-schedulerChan
		status := getLatestStatus(caseNumber)

		if previousCaseStatus.Status != status.Status {
			notifier.Notify(fmt.Sprintf("Status was changed from %s to %s\n", previousCaseStatus.Status, status.Status))
		}

		db.saveCaseStatus(caseNumber, &status)
		previousCaseStatus = status
	}
}

func getLatestStatus(caseNumber string) caseStatus {
	resp, err := http.PostForm(caseStatusURL, url.Values{"appReceiptNum": {caseNumber}})

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		log.Fatal(err)
	}

	mainRow := doc.Find(caseStatusContainerSelector)
	status := mainRow.Find(caseStatusSelector).First().Text()
	description := mainRow.Find(caseStatusDescriptionSelector).First().Text()

	return caseStatus{status, description}
}
