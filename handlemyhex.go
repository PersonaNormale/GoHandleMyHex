package gohandlemyhex

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Remove all Hexadecimal Prefixes present in a string
func removeHexadecimalPrefixes(input string) string {
	input = strings.ToLower(input)

	remHexPrefRegex := regexp.MustCompile(`0x`)

	return remHexPrefRegex.ReplaceAllLiteralString(input, "")
}

// Pick a rune and return a string if it belongs to hex
var valid = map[rune] string {
		'0':"0",
		'1':"1",
		'2':"2",
		'3':"3",
		'4':"4",
		'5':"5",
		'6':"6",
		'7':"7",
		'8':"8",
		'9':"9",
		'a':"a",
		'b':"b",
		'c':"c",
		'd':"d",
		'e':"e",
		'f':"f",
}

// Filter and prepare the string so the result string only have hexadecimal values
// This is useful for passing arrays and other strings that aren't in a numerical format
func prepareHexString(input string) string {
	if len(input) == 0 {
		return ""
	}

	input = removeHexadecimalPrefixes(input)

	res := ""

	for i := 0; i < len(input); i++ {
		res += valid[rune(input[i])]
	}	

	return res
}

// HexStringToIntSlice convert a string that match 0[xX][0-9a-fA-F]+
// It returns a []byte slice, return an error if the string doesn't match the regexp
func HexStringToByteSlice(input string) ([]byte, error) {
	input = prepareHexString(input)

	if len(input) == 0 {
		return []byte{}, nil
	}

	if len(input) % 2 == 1 {
		return nil, errors.New("Input string has odd number of character")
	}

	convertedHexSlice := []byte{}
	
	for i := 0; i < len(input); i+=2 {
		tmp, err := strconv.ParseUint(input[i:i+2], 16, 8)
		check(err)
		convertedHexSlice = append(convertedHexSlice, byte(tmp))
	}

	return convertedHexSlice, nil
}

// HexStringToIntSlice convert a string that match 0[xX][0-9a-fA-F]+
// It returns a []int slice, return an error if the string doesn't match the regexp
func HexStringToIntSlice(input string) ([]int, error){
	input = prepareHexString(input)

	if len(input) == 0 {
		return []int{}, nil
	}

	convertedHexSlice := []int{}
	
	for i := 0; i < len(input); i++ {
		tmp, err := strconv.ParseUint(string(input[i]), 16, 4)
		check(err)
		convertedHexSlice = append(convertedHexSlice, int(tmp))
	}

	return convertedHexSlice, nil
}

// HexStringToUnicode convert a string that match 0[xX][0-9a-fA-F]+
// It returns a String Rapresenting in Unicode 
func HexStringToUnicode(input string) (string, error) {
	input = prepareHexString(input)

	if len(input) == 0 {
		return "", nil
	}

	if len(input) % 2 == 1 {
		return "", errors.New("Input string has odd number of character")
	}

	convertedHexSlice := ""
	for i := 0; i < len(input); i+=2 {
		tmp, err := strconv.ParseUint(input[i: i+2], 16, 16)
		check(err)
		convertedHexSlice += string(rune(tmp))
	}

	return convertedHexSlice, nil
}
