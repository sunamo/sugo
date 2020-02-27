package sugo

import (
	"errors"
	"fmt"
	"strings"
)

// ra rune at
func ra(s string, p int) string {
	r := string([]rune(s))[p]
	return fmt.Sprintf("%v", r)
}

//Join strings
func Join(sep string, a ...interface{}) string {
	fs := make([]string, len(a))

	for i := 0; i < len(a); i++ {
		o := a[i]
		str := ToStr(o)
		fs[i] = str
	}
	return strings.Join(fs, sep)
}

// ToStr convert A1 to string
func ToStr(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

// FirstChar convert to rune
func FirstChar(a interface{}) string {
	r := fmt.Sprintf("%v", a)[0]
	return ToStr(r)
}

// reverse
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// PrintFactors return all factors of A1
func PrintFactors(nr int64) ([]int64, error) {
	fs := make([]int64, 1)
	if nr < 1 {

		return fs, errors.New(Join("", "\nFactors of ", nr, " not computed"))
	}

	fs[0] = 1
	apf := func(p int64, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := 0; j < n; j++ {
				fs = append(fs, fs[j]*pp)
			}
		}
	}
	e := 0
	for ; nr&1 == 0; e++ {
		nr >>= 1
	}
	apf(2, e)
	for d := int64(3); nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}

	return fs, nil
}

// Int convert []int to []interface{}
func Int(a []int) []interface{} {
	l := len(a)
	arr := make([]interface{}, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	return arr
}

// Int64 convert []int64 to []interface{}
func Int64(a []int64) []interface{} {
	l := len(a)
	arr := make([]interface{}, l)
	for i := 0; i < l; i++ {
		arr[i] = a[i]
	}
	return arr
}

// Contains whether A1 contains A1. Cant use []inteface for []int64 https://golang.org/doc/faq#convert_slice_of_interface, https://github.com/golang/go/issues/6804
func Contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if ToStr(a) == ToStr(e) {
			return true
		}
	}
	return false
}

// During problem with import packages simply entry really not exists. popup will show paths

// Exc print error if exists
func Exc(e error) {
	if e != nil {
		fmt.Println((e))
	}
}
