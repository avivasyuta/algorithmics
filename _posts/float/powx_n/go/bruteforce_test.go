package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPowBruteforce(t *testing.T) {
	testCases := []struct {
		name string
		x    float64
		n    int
		out  float64
	}{
		{
			name: "2^10",
			x:    2.0,
			n:    10,
			out:  1024,
		},
		{
			name: "2.1^3",
			x:    2.1,
			n:    3,
			out:  9.26100,
		},
		{
			name: "2^-2",
			x:    2,
			n:    -2,
			out:  0.25,
		},
		{
			name: "1^123456",
			x:    1,
			n:    123456,
			out:  1,
		},
		{
			name: "-1^123456",
			x:    -1,
			n:    123456,
			out:  1,
		},
		{
			name: "-1^123457",
			x:    -1,
			n:    123457,
			out:  -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := myPowBruteforce(testCase.x, testCase.n)

			//  Обрезаем кол-во знаков после запятой, чтобы нивелировать проблемы с точностью чисел с плавающей точкой
			assert.Equal(t, fmt.Sprintf("%.10f", testCase.out), fmt.Sprintf("%.10f", res))
		})
	}

}
