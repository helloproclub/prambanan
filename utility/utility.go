package utility

import (
	"fmt"
	"github.com/helloproclub/prambanan/errors"
	"strconv"
	"time"
)

// ParseBirthDateCode parse birth date code on NIK into gender and birth date
func ParseBirthDateCode(birthDateCode string) (gender, birthDate string, err error) {
	day, _ := strconv.Atoi(birthDateCode[:2])
	if 71 < day {
		err = errors.ErrGender
		return
	}

	reducer := 0
	rawGender, _ := strconv.Atoi(birthDateCode[0:1])
	if rawGender < 4 {
		gender = "pria"
	} else {
		gender = "wanita"
		reducer = 4
	}
	day -= reducer * 10

	month, _ := strconv.Atoi(birthDateCode[2:4])
	if month < 1 || 12 < month {
		err = errors.ErrMonth
		return
	}

	year, _ := strconv.Atoi(birthDateCode[4:6])
	currentYear := time.Now().Year()
	if year <= (currentYear-17)%100 {
		year += currentYear / 100 * 100
	} else {
		year += currentYear/100*100 - 100
	}

	birthDate = fmt.Sprintf("%v-%v-%v", year, month, day)

	return
}
