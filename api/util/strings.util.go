package apiutil

import (
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

func StringToInt(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return i
}

func StringToInt8(s string, d int8) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return d
	}
	return int8(i)
}

func StringToInt16(s string, d int16) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return d
	}
	return int16(i)
}

func StringToInt32(s string, d int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return d
	}
	return int32(i)
}

func StringToInt64(s string, d int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return d
	}
	return i
}

func StringToUint(s string, d uint) uint {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return d
	}
	return uint(i)
}

func StringToUint8(s string, d uint8) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return d
	}
	return uint8(i)
}

func StringToUint16(s string, d uint16) uint16 {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return d
	}
	return uint16(i)
}

func StringToUint32(s string, d uint32) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return d
	}
	return uint32(i)
}

func StringToUint64(s string, d uint64) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return d
	}
	return i
}

func StringToFloat32(s string, d float32) float32 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return d
	}
	return float32(i)
}

func StringToFloat64(s string, d float64) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return d
	}
	return i
}

func StringToBool(s string, d bool) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return d
	}
	return i
}

func StringToUUID(s string) uuid.UUID {
	u, err := uuid.Parse(s)

	if err != nil {
		return uuid.Nil
	}

	return u
}

func StringToUUIDList(s string) []uuid.UUID {
	if s == "" {
		return []uuid.UUID{}
	}

	ids := strings.Split(s, ",")

	uuids := make([]uuid.UUID, 0)
	for _, id := range ids {
		uid := StringToUUID(id)

		if uid != uuid.Nil {
			uuids = append(uuids, uid)
		}
	}

	return uuids
}

func StringToList(s string, sep string) []string {
	if s == "" {
		return []string{}
	}

	return strings.Split(s, sep)
}

func MustStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func MustStringToInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}

func MustStringToInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}

func MustStringToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}

func MustStringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func MustStringToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint(i), err
}

func MustStringToUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	return uint8(i), err
}

func MustStringToUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	return uint16(i), err
}

func MustStringToUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint32(i), err
}

func MustStringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func MustStringToFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func MustStringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func MustStringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func MustStringToList(s string, sep string) ([]string, error) {
	if s == "" {
		return []string{}, nil
	}

	return strings.Split(s, sep), nil
}

// RemovePrefix removes the prefix from the string if it exists, otherwise it returns the original string
func RemovePrefix(s string, prefix string) string {
	if len(s) >= len(prefix) && s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}

// RemoveSuffix removes the postfix from the string if it exists, otherwise it returns the original string
func RemoveSuffix(s string, postfix string) string {
	if len(s) >= len(postfix) && s[len(s)-len(postfix):] == postfix {
		return s[:len(s)-len(postfix)]
	}
	return s
}

// CountOccurrences counts the number of occurrences of a substring in a string
func CountOccurrences(s string, substr string) int {
	return strings.Count(s, substr)
}

func RemoveContiguousDuplicates(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] != s[i-1] {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// RemoveSpaces removes all spaces from the string
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// PadLeft pads the string with the given pad string until the string is at least the given length
// If the string is empty or already longer than the given length, the original string is returned
func PadLeft(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}

	return strings.Repeat(pad, length-len(s)) + s
}

// PadRight pads the string with the given pad string until the string is at least the given length
// If the string is empty or already longer than the given length, the original string is returned
func PadRight(s string, length int, pad string) string {
	if s == "" || len(s) >= length {
		return s
	}

	return s + strings.Repeat(pad, length-len(s))
}

// FindAllMatches finds all matches of the given regex in the string
// If the regex is invalid, an error is returned
func FindAllMatches(s string, regex string) ([]string, error) {
	re, err := regexp.Compile(regex)

	if err != nil {
		return nil, err
	}

	return re.FindAllString(s, -1), nil
}

// ToCamelCase converts the string to camel case
// If initCase is true, the first letter is capitalized
// If acronyms is not empty, the given acronyms are used to convert the string
func ToCamelCase(s string, initCase bool, acronyms ...map[string]string) string {
	defaultAcronyms := map[string]string{
		"ID":    "id",
		"URL":   "url",
		"HTTP":  "http",
		"HTTPS": "https",
	}

	if len(acronyms) == 0 {
		acronyms = append(acronyms, defaultAcronyms)
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if len(acronyms) > 0 {
		acronym := acronyms[0]
		if a, ok := acronym[s]; ok {
			s = a
		}
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNum := v >= '0' && v <= '9'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
