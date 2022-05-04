package ex1math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	// require lib from [Testify](https://github.com/stretchr/testify)
	require.Equal(t, 3, result)
}

func TestSubtract(t *testing.T) {
	//TODO
	testcases := []struct {
		name   string
		arg1   int
		arg2   int
		expect int
	}{
		{
			name: "ok WHEN arg1, arg2 are both positive",
			arg1: 10, arg2: 12,
			expect: 24,
		},
		{ //IMPORTANT: type int will be referred to int32 or int64 according to your runtime machine
			name: "overflow when arg1 + arg2 too large",
			arg1: 0, arg2: 0, //TODO make arg1 + arg2 large enough to run overflow
			expect: -9223372036854775808,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.arg1, tc.arg2)
			require.Equal(t, tc.expect, got)
		})
	}
}
