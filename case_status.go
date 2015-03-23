package main

import (
	"fmt"
)

type caseStatus struct {
	Status      string
	Description string
}

func (cs caseStatus) String() string {
	return fmt.Sprintf("%s\n%s", cs.Status, cs.Description)
}
