package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dollarsignteam/go-utils"
)

func TestPointerOf(t *testing.T) {
	value := 0
	result := utils.PointerOf(value)
	assert.Equal(t, &value, result)
}

func TestPackageName(t *testing.T) {
	name := utils.PackageName()
	assert.Equal(t, "go-utils_test", name)
}

func TestUniqueOf(t *testing.T) {
	tests := []struct {
		input    []any
		expected []any
	}{
		{
			input:    []any{1, 2, 3, 2, 1},
			expected: []any{1, 2, 3},
		},
		{
			input:    []any{"foo", "bar", "baz", "foo"},
			expected: []any{"foo", "bar", "baz"},
		},
		{
			input:    []any{1, "foo", true, 2, "foo"},
			expected: []any{1, "foo", true, 2},
		},
	}
	for _, test := range tests {
		result := utils.UniqueOf(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestValueOf(t *testing.T) {
	var i *int
	result := utils.ValueOf(i)
	assert.Equal(t, 0, result)

	value := 5
	result = utils.ValueOf(&value)
	assert.Equal(t, value, result)

	var s *string
	resultStr := utils.ValueOf(s)
	assert.Equal(t, "", resultStr)

	str := "hello"
	resultStr = utils.ValueOf(&str)
	assert.Equal(t, str, resultStr)

	var f *float64
	resultFloat := utils.ValueOf(f)
	assert.Equal(t, 0.0, resultFloat)

	floatVal := 3.14
	resultFloat = utils.ValueOf(&floatVal)
	assert.Equal(t, floatVal, resultFloat)
}

func BenchmarkUniqueOf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.UniqueOf([]any{1, "foo", true, 2, "foo"})
	}
}
