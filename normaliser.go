package phone

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// PlusSignRegex ...
	PlusSignRegex = `(?i)^\+`
	// NonDigitChar ...
	NonDigitChar = `(?i)\D`
	// ZeroDigitChars ...
	ZeroDigitChars = `^0+`
	// EightDigitChars ...
	EightDigitChars = `^8+`
	// RusStartDigitChars ...
	RusStartDigitChars = `^89`
)

var (
	// NonLeadingZeroCountries ...
	NonLeadingZeroCountries = []string{"GAB", "CIV", "COG"}
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("no valid format found")
	// ErrPhoneMiss ...
	ErrPhoneMiss = fmt.Errorf("unable to locate data from phone number")
)

// Normalize ...
func Normalize(p, c string) (*PhoneResult, error) {
	var (
		phone   string
		country string
		err     error
	)

	hasPlusSign := containsPlusSign(p)

	phone, err = normalizePhoneNumber(p)
	if err != nil {
		return nil, err
	}

	country, err = normalizeCountry(c)
	if err != nil {
		return nil, err
	}

	// if no country, default is USA
	var iso3166 *Data
	if country != "" {
		iso3166 = GetISO3166(country)
	}

	var result *Result
	if iso3166 != nil {
		result, err = normalizeWithCountry(phone, iso3166, hasPlusSign)
	} else {
		result, err = normalizeWithoutCountry(phone, hasPlusSign)
	}

	if err != nil {
		return nil, err
	}

	if result.PhoneData == nil {
		return nil, ErrPhoneMiss
	}

	if IsValidPhoneISO3166(result.PhoneNumber, result.PhoneData) {
		return result.WithPlusSign().AsPhoneResult(), nil
	}

	return result.AsPhoneResult(), ErrNotFound
}

func normalizePhoneNumber(phone string) (string, error) {
	reg, err := regexp.Compile(NonDigitChar)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(strings.TrimSpace(phone), ""), nil
}

func normalizeCountry(country string) (string, error) {
	return strings.ToUpper(strings.TrimSpace(country)), nil
}

func containsPlusSign(phone string) bool {
	if matched, err := regexp.MatchString(PlusSignRegex, phone); err == nil {
		return matched
	}
	return false
}

func containsString(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func containsInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func removeZeros(phone string) (string, error) {
	reg, err := regexp.Compile(ZeroDigitChars)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(phone, ""), nil
}

func isRussianPhoneNum(phone, country string) bool {
	if country == "RUS" && len(phone) == 11 {
		if matched, err := regexp.MatchString(RusStartDigitChars, phone); err == nil {
			return matched
		}
		return false
	}
	return false
}

func cleanRussianPhoneNum(phone string) (string, error) {
	reg, err := regexp.Compile(EightDigitChars)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(phone, ""), nil
}

func normalizeWithCountry(phone string, data *Data, hasPlusSign bool) (*Result, error) {
	var err error
	if !containsString(NonLeadingZeroCountries, data.CountryData.Alpha3) {
		if phone, err = removeZeros(phone); err != nil {
			return nil, err
		}
	}

	if isRussianPhoneNum(phone, data.CountryData.Alpha3) {
		if phone, err = cleanRussianPhoneNum(phone); err != nil {
			return nil, err
		}
	}

	if !hasPlusSign {
		lenOfPhoneNum := len(phone)
		if containsInt(data.PhoneNumberLengths, lenOfPhoneNum) {
			phone = fmt.Sprintf("%s%s", data.CountryData.CountryCode, phone)
		}
	}

	return NewResult(phone, data), nil
}

func normalizeWithoutCountry(phone string, hasPlusSign bool) (*Result, error) {
	// Let's try and brute force it anyway
	if possible, err := GetISO3166ByPhone(phone); err == nil && possible != nil {
		return NewResult(phone, possible), nil
	}

	if hasPlusSign {
		return nil, ErrPhoneMiss
	}

	// We really have no idea what this phone number is!
	var (
		data          = GetDefaultISO3166()
		lenOfPhoneNum = len(phone)
	)
	if containsInt(data.PhoneNumberLengths, lenOfPhoneNum) {
		phone = fmt.Sprintf("1%s", phone)
	}

	return NewResult(phone, data), nil
}
