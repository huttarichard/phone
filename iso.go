package phone

import (
	"fmt"
	"regexp"
	"strings"
)

// GetDefaultISO3166 ...
func GetDefaultISO3166() *Data {
	return ISO3166Data[0]
}

// GetISO3166 ...
func GetISO3166(country string) *Data {
	var res *Data
	switch len(country) {
	case 0:
		res = ISO3166Data[0]
		break
	case 2:
		for _, v := range ISO3166Data {
			if country == v.CountryData.Alpha2 {
				res = v
				break
			}
		}
	case 3:
		for _, v := range ISO3166Data {
			if country == v.CountryData.Alpha3 {
				res = v
				break
			}
		}
	default:
		lower := strings.ToLower(country)
		for _, v := range ISO3166Data {
			if lower == strings.ToLower(v.CountryData.Name) {
				res = v
				break
			}
		}
	}
	return res
}

// GetISO3166ByPhone ...
func GetISO3166ByPhone(phone string) (*Data, error) {
	lenOfPhoneNum := len(phone)

	for _, v := range ISO3166Data {
		if reg, err := regexp.Compile(fmt.Sprintf("^%s", v.CountryData.CountryCode)); err == nil && reg.MatchString(phone) {

			for _, x := range v.PhoneNumberLengths {
				if lenOfPhoneNum == (len(v.CountryData.CountryCode) + x) {
					// it match.. but may have more than one result.
					// e.g. USA and Canada. need to check mobileBeginWith

					for _, y := range v.MobileBeginsWith {
						if reg, err := regexp.Compile(fmt.Sprintf("^%s%s", v.CountryData.CountryCode, y)); err == nil {
							if reg.MatchString(phone) {
								return v, nil
							}
						} else {
							return nil, err
						}
					}
				}
			}

		} else if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// IsValidPhoneISO3166 ...
func IsValidPhoneISO3166(p string, data *Data) bool {
	if reg, err := regexp.Compile(fmt.Sprintf("^%s", data.CountryData.CountryCode)); err == nil {
		phone := reg.ReplaceAllString(p, "")
		lenOfPhoneNum := len(phone)
		for _, v := range data.PhoneNumberLengths {
			if lenOfPhoneNum == v {
				for _, x := range data.MobileBeginsWith {
					if reg, err := regexp.Compile(fmt.Sprintf("^%s", x)); err == nil {
						if reg.MatchString(phone) {
							return true
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}
