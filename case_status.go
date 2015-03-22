package main

import (
	"fmt"
)

type CaseStatus struct {
	Status      string
	Description string
}

func (cs CaseStatus) String() string {
	return fmt.Sprintf("%s\n%s", cs.Status, cs.Description)
}
