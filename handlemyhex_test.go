package gohandlemyhex

import (
	"testing"
	"reflect"
)

func TestHexStringToByteSlice(t *testing.T) {
    tests := []struct {
        input string
        expectedOutput []byte
        expectError bool
    }{
        {
            input: "0x12AbCd",
            expectedOutput: []byte{0x12, 0xab, 0xcd},
            expectError: false,
        },
        {
            input: "0xabCdEf",
            expectedOutput: []byte{0xab, 0xcd, 0xef},
            expectError: false,
        },
        {
            input: "",
            expectedOutput: []byte{},
            expectError: false,
        },
        {
            input: "0x1",
            expectedOutput: nil,
            expectError: true,
        },
		{
			input: "{0x12, 0xAbCd}",
			expectedOutput: []byte{0x12, 0xab, 0xcd},
			expectError: false,
		},
		{
			input: "[0x12, 0xAb, 0xCd]",
			expectedOutput: []byte{0x12, 0xab, 0xcd},
			expectError: false,
		},
		{
			input: "ciao ciao",
			expectedOutput: []byte{0xca, 0xca},
			expectError: false,
		},
		{
			input: "byte{ciao ,,,,     ciao}",
			expectedOutput: []byte{0xbe, 0xca, 0xca},
			expectError: false,
		},
    }

    for _, test := range tests {
        output, err := HexStringToByteSlice(test.input)

        if test.expectError && err == nil {
            t.Errorf("Test failed: expected error, got no error")
        } else if !test.expectError && err != nil {
            t.Errorf("Test failed: unexpected error %s", err)
        } else if !reflect.DeepEqual(output, test.expectedOutput) {
            t.Errorf("Test failed: expected %s, got %s", test.expectedOutput, output)
        }
    }
}

func TestHexStringToIntSlice(t *testing.T) {
    tests := []struct {
        input string
        expectedOutput []int
        expectError bool
    }{
		{
            input: "0x12AbCd",
            expectedOutput: []int{1, 2, 10, 11, 12, 13},
            expectError: false,
        },
        {
            input: "0xabCdEf",
            expectedOutput: []int{10, 11, 12, 13, 14, 15},
            expectError: false,
        },
        {
            input: "",
            expectedOutput: []int{},
            expectError: false,
        },
		{
			input: "ciao ciao",
			expectedOutput: []int{12,10,12,10},
			expectError: false,
		},
		{
			input: "byte{ciao ciao}",
			expectedOutput: []int{11,14,12,10,12,10},
			expectError: false,
		},
    }

    for _, test := range tests {
        output, err := HexStringToIntSlice(test.input)

        if test.expectError && err == nil {
            t.Errorf("Test failed: expected error, got no error")
        } else if !test.expectError && err != nil {
            t.Errorf("Test failed: unexpected error %s", err)
        } else if !reflect.DeepEqual(output, test.expectedOutput) {
            t.Errorf("Test failed: expected %v, got %v", test.expectedOutput, output)
        }
    }
}

func TestHexStringToUnicode(t *testing.T) {
	tests := []struct {
		input string
		expectedOutput string
		expectError bool
	}{
		{
			input: "0x4369616f6369616f",
			expectedOutput: "Ciaociao",
			expectError: false,
	
		},
		{
			input: "",
			expectedOutput: "",
			expectError: false,
		},
	}

	for _, test := range tests {
		output, err := HexStringToUnicode(test.input)

		if test.expectError && err == nil {
			t.Errorf("Test failed: expected error, gor no error")
		} else if !test.expectError && err != nil {
			t.Errorf("Test failed unecpected error %s", err)
		} else if !reflect.DeepEqual(output, test.expectedOutput) {
			t.Errorf("Test failed: expected %s, got %s", test.expectedOutput, output)
		}
	}
}
