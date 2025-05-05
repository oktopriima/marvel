package validates

// Validator contract ...
type Validator interface {
	Request(value interface{}) error
	OperatorChecker(msisdn string) (*DataProvider, string)
	MatchURL(value string) bool
	MatchSpace(value string) bool
	MatchEmail(value string) bool
}
