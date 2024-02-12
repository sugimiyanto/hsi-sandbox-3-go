package tahap_2

import (
	"fmt"
)

func SortNIP(values []string) []string {
	index_sorted := len(values) - 1
	for ; index_sorted > 0; {
		for i := 0; i < index_sorted; i++ {
			if values[i] > values[i+1] {
				values[i], values[i+1] = values[i+1], values[i]
			}
		}
		index_sorted -= 1
	}	
	return values
}

func GroupAndSortNip(nips []string) map[string][]string {
	groupedStudents := make(map[string][]string)
	var prefix_gender, gender string
	for _, nip := range nips {
		prefix_gender = nip[0:3]
		if prefix_gender == "ARN" {
			gender = "male"
		} else {
			gender = "female"
		}

		if _, ok := groupedStudents[gender]; ok {
			groupedStudents[gender] = append(groupedStudents[gender], nip)
		} else {
			groupedStudents[gender] = []string{nip}
		}
	}

	for gender, nips := range groupedStudents {
		groupedStudents[gender] = SortNIP(nips)
	}
	fmt.Println(groupedStudents)
	return groupedStudents
}