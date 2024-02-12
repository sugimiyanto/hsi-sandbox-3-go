package soal_1_2

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGenerateNIKHSI(t *testing.T) {
	var actual []string
	var expected []string

	current_time := time.Now()
	threshold := time.Date(time.Now().Year(), time.July, 1, 0, 0, 0, 0, time.Local)
	var academic_year int
	if current_time.After(threshold) || current_time.Equal(threshold) {
		academic_year = 2
	} else {
		academic_year = 1
	}

	expected = []string {
		fmt.Sprintf("ARN%d19-00001", academic_year),
		fmt.Sprintf("ARN%d19-00002", academic_year),
		fmt.Sprintf("ARN%d19-00003", academic_year),
	}
	actual = GenerateNIKHSI("male", 2019, 3)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected value: %v\nactual: %v", expected, actual)
	}

	// testing female students
	expected = []string {
		fmt.Sprintf("ART%d24-00001", academic_year),
		fmt.Sprintf("ART%d24-00002", academic_year),
		fmt.Sprintf("ART%d24-00003", academic_year),
		fmt.Sprintf("ART%d24-00004", academic_year),
	}
	actual = GenerateNIKHSI("female", 2024, 4)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected value: %v\nactual: %v", expected, actual)
	}
}
