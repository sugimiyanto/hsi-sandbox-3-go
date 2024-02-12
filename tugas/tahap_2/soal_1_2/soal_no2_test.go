package soal_1_2

import (
	"errors"
	"reflect"
	"testing"
)

func TestContinueGenerateNIKHSIValid(t *testing.T) {
	expected := []string {
		"ARN219-00123",
		"ARN219-00124",
		"ARN219-00125",
	}
	actual, err := ContinueGenerateNIKHSI("ARN219-00122", 3)
	if !reflect.DeepEqual(actual, expected) || err != nil {
		t.Fatalf("Expected value: %v\nactual: %v", expected, actual)
	}
}

func TestContinueGenerateNIKHSIInvalid(t *testing.T) {
	expected_err := errors.New("Input NIK should be in format AR[N|T]xxx-xxxxx")
	_, err := ContinueGenerateNIKHSI("N192-075", 3)
	if err.Error() != expected_err.Error() {
		t.Fatalf("Expected error: '%s'. Yours: '%s'", expected_err, err)
	}
}