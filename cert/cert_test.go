package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2019-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err = %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got = nil")
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Expected = 'GOLANG COURSE' got = %v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T)  {
	_, err := New("", "Bob", "2019-05-31")
	if err == nil {
		t.Error("Error should be returned an empty course")
	}
}

func TestNameEmptyValue(t *testing.T)  {
	_, err := New("Golang", "", "2019-05-31")
	if err == nil {
		t.Error("Error should be returned an empty name")
	}
}

func TestCourseTooLong(t *testing.T)  {
	course := "azzzzerthgbfdqetryhtjhngvtreythrngbddrtrethgbfvdezgebdvcssrgfbvcxd"
	_, err := New(course, "bob", "2019-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course = %s)", course)
	}
}

func TestNameTooLong(t *testing.T)  {
	name := "azzzzerthgbfdqetryhtjhngvtreythrngbddrtrethgbfvdezgebdvcssrgfbvcxd"
	_, err := New("Golang",name, "2019-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long student name (name = %s)", name)
	}
}