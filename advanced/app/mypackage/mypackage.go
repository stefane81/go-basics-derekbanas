package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Stefan"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

type Date struct {
	day, month, year int
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}

func (m *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("incorrect month value")
	}
	m.month = month
	return nil
}

func (y *Date) SetYear(year int) error {
	if (year < 1900) || (year > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	y.year = year
	return nil
}

func (d *Date) GetDay() int {
	return d.day
}

func (m *Date) GetMonth() int {
	return m.month
}

func (y *Date) GetYear() int {
	return y.year
}
