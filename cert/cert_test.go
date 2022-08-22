package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is nos valid. expected = 'GOLANG COURSE', got = '%v'", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-21")
	if err == nil {
		t.Error("Error should be returned on a empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azerttuoiehfoihzeifhzoibfoizhfoizheoizhfiouahnfoiuh"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course = %s", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-05-21")
	if err == nil {
		t.Error("Error should be returned on a empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "azerttuoiehfoihzeifhzoibfoizhfoizheoizhfiouahnfoiuh"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (name = %s", name)
	}
}
