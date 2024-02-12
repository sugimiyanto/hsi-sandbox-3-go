package soal_1_2

import (
	"fmt"
	"strings"
	"time"
)

func GenerateNIKHSI(gender string, tahun int, jumlah_yang_digenerate int) []string {
	var nik []string
	var new_nik string
	gender_infix := ""
	if strings.ToLower(gender) == "male" {
		gender_infix = "N"
	} else {
		gender_infix = "T"
	}
	current_time := time.Now()
	threshold := time.Date(time.Now().Year(), time.July, 1, 0, 0, 0, 0, time.Local)
	var academic_year int
	if current_time.After(threshold) || current_time.Equal(threshold) {
		academic_year = 2
	} else {
		academic_year = 1
	}
	short_year := tahun % 100

	for i := 1; i <= jumlah_yang_digenerate; i++ {
		new_nik = fmt.Sprintf("AR%s%d%d-%05d", gender_infix, academic_year, short_year, i)
		nik = append(nik, new_nik)
	}
	return nik
}
