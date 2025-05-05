package validates

import "regexp"

// MatchURL ...
func (v *validatorLib) MatchURL(value string) bool {
	exp := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z A-Z 0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	match, _ := regexp.Compile(exp)
	if !match.MatchString(value) {
		return false
	}
	return true
}

// MatchSpace ...
func (v *validatorLib) MatchSpace(value string) bool {
	exp := `^[a-zA-Z0-9_.]+$`
	match, _ := regexp.Compile(exp)
	if !match.MatchString(value) {
		return false
	}
	return true
}

// MatchEmail ...
func (v *validatorLib) MatchEmail(value string) bool {
	exp := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$"
	match, _ := regexp.Compile(exp)
	if !match.MatchString(value) {
		return false
	}
	return true
}
