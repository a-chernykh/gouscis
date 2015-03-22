package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	CaseStatusUrl = "https://egov.uscis.gov/casestatus/mycasestatus.do"

	CaseStatusContainerSelector   = ".main-row"
	CaseStatusSelector            = "h1"
	CaseStatusDescriptionSelector = ".rows p"
)

func statusUpdate(schedulerChan chan bool, caseNumber string, db Database, notifier Notifier) {
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

func getLatestStatus(caseNumber string) CaseStatus {
	resp, err := http.PostForm(CaseStatusUrl, url.Values{"appReceiptNum": {caseNumber}})

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		log.Fatal(err)
	}

	mainRow := doc.Find(CaseStatusContainerSelector)
	status := mainRow.Find(CaseStatusSelector).First().Text()
	description := mainRow.Find(CaseStatusDescriptionSelector).First().Text()

	return CaseStatus{status, description}
}
