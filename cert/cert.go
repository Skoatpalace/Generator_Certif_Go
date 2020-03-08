package cert

import (
	"fmt"
	"time"
)

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {

	c := course
	n := name
	d := date
	
	cert := &Cert{
		Course : c,
		Name : n,
		LabelTitle : fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion : "Certificate of Completion", 
		LabelPresented : "This Certificate is Presented To",
		LabelParticipation : fmt.Sprintf("For participation in the %v",c),
		LabelDate : fmt.Sprintf("%v", d),		
	}
	return cert, nil
}
