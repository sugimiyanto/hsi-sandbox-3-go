package soal_1_2

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ContinueGenerateNIKHSI(previous_nik string, jumlah_yang_digenerate int) ([]string, error){
	valid_pattern := "AR[N|T][0-9]{3}-[0-9]{5}"
	regex := regexp.MustCompile(valid_pattern)

	match := regex.MatchString(previous_nik)
	if !match {
		return []string{}, errors.New("Input NIK should be in format AR[N|T]xxx-xxxxx")
	}

	extract_number := strings.Split(previous_nik, "-")
	current_prefix_nik := extract_number[0]
	current_number := extract_number[len(extract_number) - 1]
	current_number_int, err := strconv.Atoi(current_number)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	var continue_nik []string
	var new_nik string
	for i := 1; i<= jumlah_yang_digenerate; i++ {
		current_number_int += 1
		new_nik = fmt.Sprintf("%s-%05d", current_prefix_nik, current_number_int)
		continue_nik = append(continue_nik, new_nik)
	}
	return continue_nik, nil
}