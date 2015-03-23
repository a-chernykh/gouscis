package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type database struct {
	Path string
}

func (db *database) caseStatusPath(caseNumber string) string {
	hash := sha256.New()
	io.WriteString(hash, caseNumber)
	hashedCaseNumber := fmt.Sprintf("%x", hash.Sum(nil))

	return filepath.Join(db.Path, hashedCaseNumber)
}

func (db *database) create() {
	err := os.Mkdir(db.Path, 0755)
	checkFileError(err)
}

func (db *database) saveCaseStatus(caseNumber string, caseStatus *caseStatus) {
	filePath := db.caseStatusPath(caseNumber)
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()

	_, err = f.WriteString(caseStatus.Status)
	check(err)
}

func (db *database) loadCaseStatus(caseNumber string) caseStatus {
	filePath := db.caseStatusPath(caseNumber)
	dat, err := ioutil.ReadFile(filePath)
	checkFileError(err)

	return caseStatus{Status: string(dat)}
}
