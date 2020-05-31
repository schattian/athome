package normalize

import (
	"sort"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func stripAccents(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		return "", errors.Wrap(err, "transform.String")
	}
	return result, nil
}

func removeNonWord(s string) string {
	return clean([]byte(s))
}

func ToNormal(s string) (string, error) {
	s = strings.ToLower(s)
	s, err := stripAccents(s)
	if err != nil {
		return "", errors.Wrap(err, "stripAccents")
	}
	s = removeNonWord(s)
	s = strings.TrimSpace(s)
	return s, nil
}

// CompareSliceSoft compares, and sorts before, each word in a[i] with b[j]
func CompareSliceSoft(a, b []string) (eq bool, err error) {
	if len(b) > len(a) {
		eq, err = CompareSlice(b, a)
	}
	if len(a) > len(b) {
		eq, err = CompareSlice(a, b)
	}
	if len(a) == len(b) {
		eq, err = CompareSlice(a, b)
	}
	return
}

// CompareSlice compares, and sorts before, each word in a[i] with b[i]
func CompareSlice(a, b []string) (eq bool, err error) {
	sort.Strings(a)
	sort.Strings(b)

	for i, aword := range a {
		if len(b)-1 < i {
			break
		}
		eq, err = Compare(aword, b[i])
		if err != nil {
			return false, errors.Wrap(err, "compare")
		}
		if !eq {
			break
		}
	}
	return
}

func Compare(a, b string) (bool, error) {
	a, err := ToNormal(a)
	if err != nil {
		return false, errors.Wrap(err, "a.toNormal")
	}
	b, err = ToNormal(b)
	if err != nil {
		return false, errors.Wrap(err, "b.toNormal")
	}
	return a == b, nil
}

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}
