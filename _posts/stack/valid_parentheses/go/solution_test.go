package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "()",
			in:   "()",
			out:  true,
		},
		{
			name: "()[]{}",
			in:   "()[]{}",
			out:  true,
		},
		{
			name: "(]",
			in:   "(]",
			out:  false,
		},
		{
			name: "((((",
			in:   "((((",
			out:  false,
		},
		{
			name: "))))",
			in:   "))))",
			out:  false,
		},
		{
			name: "(){",
			in:   "(){",
			out:  false,
		},
		{
			name: "}{",
			in:   "}{",
			out:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := isValid(testCase.in)

			assert.Equal(t, testCase.out, res)
		})
	}
}
