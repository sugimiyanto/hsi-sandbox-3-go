package tahap_2

import (
	"testing"
	"reflect"
	soal_1_2 "sugi.com/soal_1_2"
)

func TestSortNIP(t *testing.T) {
	input := []string{"ARN219-00034", "ARN220-00056", "ARN219-00005"}
	expectedOutput := []string{"ARN219-00005", "ARN219-00034", "ARN220-00056"}
	actual := SortNIP(input)
	if !reflect.DeepEqual(actual, expectedOutput) {
		t.Fatalf("Expected: %v. Actual: %v", expectedOutput, actual)
	}
}

func TestGroupAndSortNip(t *testing.T) {
	students := soal_1_2.GenerateNIKHSI("male", 2019, 2)
	students = append(students, soal_1_2.GenerateNIKHSI("male", 2020, 3)...)
	students = append(students, soal_1_2.GenerateNIKHSI("female", 2019, 2)...)
	continueStudents, _ := soal_1_2.ContinueGenerateNIKHSI("ARN119-00002", 2)
	students = append(students, continueStudents...)

	// assume this unit test is always running in first-half of the year,
	// otherwise need to modify this expectedOutput as ARN1|ART1 to ARN2|ART2
	expectedOutput := map[string][]string {
		"male": {
			"ARN119-00001",
			"ARN119-00002",
			"ARN119-00003",
			"ARN119-00004",
			"ARN120-00001",
			"ARN120-00002",
			"ARN120-00003",
		},
		"female": {
			"ART119-00001",
			"ART119-00002",
		},
	}
	
	actual := GroupAndSortNip(students)
	if !reflect.DeepEqual(actual, expectedOutput) {
		t.Fatalf("Expected: %v. Actual: %v", expectedOutput, actual)
	}
}