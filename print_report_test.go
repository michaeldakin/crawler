// Shamelessly stolen from Boot.dev summary after completing this module
package main

import (
	"reflect"
	"testing"
)

func TestSortLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []Link
	}{
		{
			name: "order count descending",
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []Link{
				{URL: "url4", Hits: 10},
				{URL: "url5", Hits: 7},
				{URL: "url1", Hits: 5},
				{URL: "url3", Hits: 3},
				{URL: "url2", Hits: 1},
			},
		},
		{
			name: "alphabetize",
			input: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
			expected: []Link{
				{URL: "a", Hits: 1},
				{URL: "b", Hits: 1},
				{URL: "c", Hits: 1},
				{URL: "d", Hits: 1},
				{URL: "e", Hits: 1},
			},
		},
		{
			name: "order count then alphabetize",
			input: map[string]int{
				"d": 2,
				"a": 1,
				"e": 3,
				"b": 1,
				"c": 2,
			},
			expected: []Link{
				{URL: "e", Hits: 3},
				{URL: "c", Hits: 2},
				{URL: "d", Hits: 2},
				{URL: "a", Hits: 1},
				{URL: "b", Hits: 1},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []Link{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: []Link{},
		},
		{
			name: "one key",
			input: map[string]int{
				"url1": 1,
			},
			expected: []Link{
				{URL: "url1", Hits: 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortLinks(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL:\nexpected URL: %v\nactual:       %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

