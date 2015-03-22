package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Database struct {
	Path string
}

func (db *Database) caseStatusPath(caseNumber string) string {
	hash := sha256.New()
	io.WriteString(hash, caseNumber)
	hashedCaseNumber := fmt.Sprintf("%x", hash.Sum(nil))

	return filepath.Join(db.Path, hashedCaseNumber)
}

func (db *Database) create() {
	err := os.Mkdir(db.Path, 0755)
	checkFileError(err)
}

func (db *Database) saveCaseStatus(caseNumber string, caseStatus *CaseStatus) {
	filePath := db.caseStatusPath(caseNumber)
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()

	_, err = f.WriteString(caseStatus.Status)
	check(err)
}

func (db *Database) loadCaseStatus(caseNumber string) CaseStatus {
	filePath := db.caseStatusPath(caseNumber)
	dat, err := ioutil.ReadFile(filePath)
	checkFileError(err)

	return CaseStatus{Status: string(dat)}
}
