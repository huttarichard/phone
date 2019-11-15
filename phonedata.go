package phone

import (
	c "github.com/coinfinitygroup/phone/countries"
)

// Data ...
type Data struct {
	CountryData        c.CountryData
	PhoneCode          string
	MobileBeginsWith   []string
	PhoneNumberLengths []int
}

// NewData ...
func NewData() *Data {
	return &Data{
		CountryData:        c.CountryData{},
		MobileBeginsWith:   make([]string, 0, 0),
		PhoneNumberLengths: make([]int, 0, 0),
	}
}
