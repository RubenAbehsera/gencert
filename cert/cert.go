package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 30

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

// Interface that helps us to save a certificat -> Cert.Saver
type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificat - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, "course") {
		c = c + " course"
	}

	return strings.ToTitle(c), nil
}

func validateName(name string) (string, error) {
	n, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}

	return strings.ToTitle(n), nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}

func validateStr(str string, maxLent int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("invalid string. got='%s', len='%d'", c, len(c))
	} else if len(c) >= maxLent {
		return c, fmt.Errorf("invalid string. got='%s', len='%d'", c, len(c))
	}
	return c, nil
}
