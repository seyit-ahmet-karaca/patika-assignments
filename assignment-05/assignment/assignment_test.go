package assignment

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	cases := []struct {
		givenX, givenY   uint32
		expectedSum      uint32
		expectedOverflow bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
	}

	for _, v := range cases {
		t.Run(strconv.Itoa(int(v.givenX))+" "+strconv.Itoa(int(v.givenY)), func(t *testing.T) {
			sum, overflow := AddUint32(v.givenX, v.givenY)
			assert.Equal(t, v.expectedSum, sum)
			assert.Equal(t, v.expectedOverflow, overflow)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	cases := []struct {
		givenF        float64
		expectedPoint float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("%v", v.givenF), func(t *testing.T) {
			point := CeilNumber(v.givenF)
			assert.Equal(t, v.expectedPoint, point)
		})
	}
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	cases := []struct {
		givenWord  string
		sortedWord string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, v := range cases {
		t.Run(v.givenWord, func(t *testing.T) {
			result := AlphabetSoup(v.givenWord)

			assert.Equal(t, v.sortedWord, result)
		})
	}
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	cases := []struct {
		givenWord string
		givenN    uint
		expected  string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 3, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, v := range cases {
		t.Run(v.givenWord, func(t *testing.T) {
			result := StringMask(v.givenWord, v.givenN)

			assert.Equal(t, v.expected, result)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/

	cases := []struct {
		givenElements [2]string
		expected      string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}

	for _, v := range cases {
		t.Run(v.givenElements[0], func(t *testing.T) {
			result := WordSplit(v.givenElements)

			assert.Equal(t, v.expected, result)
		})
	}

}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	cases := []struct {
		given []interface{}
		expected []interface{}
	}{
		{[]interface{}{4,2,5,4,2,4} , []interface{}{4,2,5}},
		{[]interface{}{"bootcamp","rocks!","really","rocks!"} , []interface{}{"bootcamp","rocks!","really"}},
		{[]interface{}{1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first"} , []interface{}{1,uint32(1),"first",2,uint32(2),"second"}},
	}

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			set := VariadicSet(v.given...)

			assert.Equal(t, v.expected, set)
		})
	}
}
