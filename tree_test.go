package honey

import (
	"regexp"
	"testing"
)

func TestParamsGet(t *testing.T) {
	params := Params{
		Param{
			Key:   "name",
			Value: "hj",
		},
		Param{
			Key:   "age",
			Value: "44",
		},
	}
	for i, param := range params {
		if value, ok := params.Get(param.Key); !ok || value != param.Value {
			t.Errorf("params[%d] is %s, but get %s", i, param.Value, value)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	a := "abcde"
	b := "abde"
	i := longestCommonPrefix(a, b)
	t.Error(i, a[:i])
}

func TestRegex(t *testing.T) {
	regexStr := "\\d+"
	r, _ := regexp.Compile(regexStr)
	b := r.MatchString("a")
	t.Errorf("b is %t", b)
}
func TestCountParams(t *testing.T) {
	paths := map[string]uint16{
		"/hello":            0,
		"/hello/:name":      1,
		"/hello/*name":      1,
		"/hello/{}name":     1,
		"/hello/*name/:age": 2,
	}

	for path, count := range paths {
		if c := countParams(path); count != c {
			t.Errorf("%s path should be contains %d params,but find %d params", path, count, c)
		}
	}
}
func TestFindWildcard(t *testing.T) {
	path := "/hello/{\\w+}name/:age"
	t.Error(findWildcard(path))
}

func TestParseWildcard(t *testing.T) {
	// t.Error(parseWildcard(":name"))
	// t.Error(parseWildcard("*name"))
	_, _, regex := parseWildcard("{^\\w{4}$}name")

	r := regexp.MustCompile(regex)
	t.Error(r.FindString("3dds"))
}
