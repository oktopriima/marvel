package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm/schema"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/google/uuid"
)

// ReadFromYAML reads the YAML file and pass to the object
// args:
//
//	path: file path location
//	target: object which will hold the value
//
// returns:
//
//	error: operation state error
func ReadFromYAML(path string, target interface{}) error {
	yf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}

// PathExist check the path directory if exist
func PathExist(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// IsSameType Check Same Type
func IsSameType(src, dest interface{}) bool {
	return reflect.TypeOf(src) == reflect.TypeOf(dest)
}

const (
	letterBytes   = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	dateFormat     = `2006-01-02`
	dateTimeFormat = `2006-01-02 15:04:05`
)

// Pagination utility
func Pagination(limit, page uint64) (uint64, uint64) {
	if page <= 1 {
		return limit, 0
	}
	return limit, (page - 1) * limit
}

// Replacer string replacer helper
func Replacer(r map[string]string, msg string) string {
	for k, v := range r {
		msg = strings.ReplaceAll(msg, k, v)
	}
	return msg
}

// GenerateRandomNumberString generate random string number
// args:
//
//	integer: total randomization
//
// output:
//
//	string
func GenerateRandomNumberString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// GenerateReferenceID generates reference ID
func GenerateReferenceID(prefix string) string {
	now := time.Now().Format("20060102030405")
	buff := bytes.NewBufferString(now)
	buff.WriteString(GenerateRandomNumberString(8))

	return fmt.Sprintf("%s%s", prefix, buff.String())
}

// GenerateAppID generates reference ID
func GenerateAppID(prefix string) string {
	now := time.Now().Format("20060102030405")
	buff := bytes.NewBufferString(now)
	buff.WriteString(GenerateRandomNumberString(6))

	return fmt.Sprintf("%s%s", prefix, buff.String())
}

// GenerateUUID method
func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateUniqueIdMember(email string, requestPayload any) (string, error) {
	var err error
	var result string

	var requestEncode bytes.Buffer
	err = json.NewEncoder(&requestEncode).Encode(requestPayload)
	if err != nil {
		return result, err
	}

	md5 := md5.Sum([]byte(email + requestEncode.String()))
	result = hex.EncodeToString(md5[:])

	return result, err
}

func StringToDate(s string) time.Time {
	tm, _ := StringToDateE(s)
	return tm
}

func StringToDateE(s string) (time.Time, error) {
	tm, err := parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.000",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		"02/01/2006 15:04:05",     // indonesian date time
		"02/01/2006 15:04:05.000", // indonesian date time
		"02/01/2006",              // indonesian date
	})

	return tm, err
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}

var envArr = map[string]string{
	"production":  "prod",
	"staging":     "stg",
	"development": "dev",
	"prod":        "prod",
	"stg":         "stg",
	"dev":         "dev",
	"local":       "loc",
	"loc":         "loc",
	"prd":         "prod",
}

// EnvironmentTransform transformer
func EnvironmentTransform(s string) string {
	v, ok := envArr[strings.ToLower(strings.Trim(s, " "))]

	if !ok {
		return ""
	}

	return v
}

// DumpToString interface to string
func DumpToString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		json.NewEncoder(buff).Encode(v)
		return buff.String()
	}

	return str
}

func DebugPrint(v interface{}) {
	fmt.Println(DumpToString(v))
}

// InArray check if an element is existed in the array
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func Environment() string {
	return EnvironmentTransform(os.Getenv("APP_ENVIRONMENT"))
}

func ParseAccessToken(rawToken string) ([]byte, string, error) {
	token := strings.Split(rawToken, " ")
	if len(token) <= 1 {
		return nil, "", errors.New("invalid bearer token not valid")
	}
	if !InArray(strings.ToLower(token[0]), []string{"bearer", "basic"}) {
		return nil, "", errors.New("invalid token type")
	}

	tokenPayloads := strings.Split(token[1], ".")
	if len(tokenPayloads) <= 1 {
		return nil, "", errors.New("invalid bearer token not valid")
	}

	payloadIssue, err := base64.RawStdEncoding.DecodeString(tokenPayloads[1])
	if err != nil {
		return nil, "", errors.New("invalid issued profile token")
	}
	return payloadIssue, token[1], nil
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(value, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}

// StringToInt ..
func StringToInt(str string) int {
	i1, _ := strconv.Atoi(str)
	return i1
}

// StringToInt64 ...
func StringToInt64(str string) int64 {
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

// StringToBool ...
func StringToBool(str string) bool {
	castBool, _ := strconv.ParseBool(str)
	return castBool
}

// StrToUint64 ...
func StrToUint64(value string) uint64 {
	res, _ := strconv.ParseUint(value, 10, 64)
	return res
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

// ToSnakeCase ...
func ToSnakeCase(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}
	return string(b.r)
}

func FormatCurrencyRupiah(amount int64) string {
	formatted := "Rp" + commaSeparate(strconv.FormatInt(amount, 10))

	return formatted
}

func FormatCurrency(amount int64) string {
	formatted := commaSeparate(strconv.FormatInt(amount, 10))

	return formatted
}

func commaSeparate(s string) string {
	runes := []rune(s)
	var result string
	l := len(runes)
	for i, r := range runes {
		result += string(r)
		if (l-i-1)%3 == 0 && i != l-1 && string(r) != "-" {
			result += "."
		}
	}
	return result
}

func CensorEmail(email string) string {
	parts := strings.Split(email, "@")
	local := parts[0]
	domain := parts[1]

	if len(local) <= 3 {
		return email
	}

	censoredLocal := local[:1] + strings.Repeat("*", len(local)-2) + local[len(local)-1:]
	return censoredLocal + "@" + domain
}

func SetPercentage(percentage float64) string {
	var res string
	percentage = math.Floor(percentage*10) / 10
	if percentage == float64(int64(percentage)) {
		res = fmt.Sprintf("%.0f %%", percentage)
	} else {
		res = fmt.Sprintf("%.1f %%", percentage)
	}
	return res
}

func ToSearchableMap(attrs ...interface{}) (result interface{}) {
	if len(attrs) > 1 {
		if str, ok := attrs[0].(string); ok {
			result = map[string]interface{}{str: attrs[1]}
		}
	} else if len(attrs) == 1 {
		if attr, ok := attrs[0].(map[string]interface{}); ok {
			result = attr
		}

		if attr, ok := attrs[0].(interface{}); ok {
			s := structs.New(attr)
			s.TagName = "json"
			m := s.Map()

			value := make(map[string]interface{}, len(m))
			var ns schema.NamingStrategy
			for col, val := range m {
				dbCol := ns.ColumnName("", col)
				value[dbCol] = val
			}
			result = value
		}
	}
	return
}
