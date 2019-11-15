package phone

import (
	"fmt"
)

// Result ...
type Result struct {
	PhoneNumber string
	PhoneData   *Data
}

// NewResult ...
func NewResult(phone string, data *Data) *Result {
	return &Result{
		PhoneNumber: phone,
		PhoneData:   data,
	}
}

// WithPlusSign ...
func (r Result) WithPlusSign() *Result {
	return NewResult(fmt.Sprintf("+%s", r.PhoneNumber), r.PhoneData)
}

// AsPhoneResult ...
func (r Result) AsPhoneResult() *PhoneResult {
	return NewPhoneResult(r.PhoneNumber, r.PhoneData.CountryData.Alpha3)
}

// PhoneResult ...
type PhoneResult struct {
	PhoneNumber string
	Country     string
}

// NewPhoneResult ...
func NewPhoneResult(phone, country string) *PhoneResult {
	return &PhoneResult{
		PhoneNumber: phone,
		Country:     country,
	}
}
