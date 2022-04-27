package assignment

import (
	"math"
	"reflect"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	var sumOfParams = uint64(x) + uint64(y)
	remainU32Max := sumOfParams % uint64(math.MaxUint32+1)
	return uint32(remainU32Max), sumOfParams != remainU32Max
}

func CeilNumber(f float64) float64 {
	ceiling := 0.25
	if mod := math.Mod(f, ceiling); mod == 0 {
		return f
	} else {
		return f + ceiling - mod
	}
}

func AlphabetSoup(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func StringMask(s string, n uint) string {
	if lengthOfString := len(s); n >= uint(lengthOfString) {
		if lengthOfString == 0 {
			lengthOfString = 1
		}
		return strings.Repeat("*", lengthOfString)
	} else {
		return s[:n] + strings.Repeat("*", lengthOfString-int(n))
	}
}

func WordSplit(arr [2]string) string {
	text, words := arr[0], strings.Split(arr[1], ",")
	scrappedWords := [2]string{"", ""}

	for _, word := range words {
		if strings.HasPrefix(text, word) {
			scrappedWords[0] = word
		} else if strings.HasSuffix(text, word) {
			scrappedWords[1] = word
		}
	}

	if scrappedWords[0] == "" || scrappedWords[1] == "" {
		return "not possible"
	}
	return strings.Join(scrappedWords[:], ",")
}

func VariadicSet(i ...interface{}) []interface{} {
	set := make([]interface{}, 0)

	for _, v := range i {
		isContain := false
		for _, setV := range set {
			if reflect.TypeOf(v) == reflect.TypeOf(setV) && setV == v {
				isContain = true
				break
			}
		}
		if !isContain {
			set = append(set, v)
		}
	}
	return set
}
