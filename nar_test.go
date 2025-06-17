package nar

import (
	"slices"
	"testing"
)

func TestZip(t *testing.T) {
	type Pair struct {
		A int
		B string
	}

	tests := []struct {
		name string
		s1   []int
		s2   []string
		want []Pair
	}{
		{
			name: "zip two slices of different lengths",
			s1:   []int{1, 2, 3},
			s2:   []string{"a", "b"},
			want: []Pair{
				{1, "a"},
				{2, "b"},
			},
		},
		{
			name: "zip two slices of equal lengths",
			s1:   []int{1, 2, 3},
			s2:   []string{"a", "b", "c"},
			want: []Pair{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
		},
		{
			name: "zip two empty slices",
			s1:   []int{},
			s2:   []string{},
			want: []Pair{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []Pair
			for a, b := range Zip(tt.s1, tt.s2) {
				result = append(result, Pair{a, b})
			}

			if !slices.Equal(result, tt.want) {
				t.Errorf("Zip() got %v, want %v", result, tt.want)
			}
		})
	}

}

func TestZipLongest(t *testing.T) {
	type Pair struct {
		A int
		B string
	}

	tests := []struct {
		name string
		s1   []int
		s2   []string
		want []Pair
	}{
		{
			name: "zip two slices of different lengths",
			s1:   []int{1, 2, 3},
			s2:   []string{"a", "b"},
			want: []Pair{
				{1, "a"},
				{2, "b"},
				{3, ""},
			},
		},
		{
			name: "zip two slices of equal lengths",
			s1:   []int{1, 2, 3},
			s2:   []string{"a", "b", "c"},
			want: []Pair{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
		},
		{
			name: "zip two empty slices",
			s1:   []int{},
			s2:   []string{},
			want: []Pair{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []Pair
			for a, b := range ZipLongest(tt.s1, tt.s2) {
				result = append(result, Pair{a, b})
			}

			if !slices.Equal(result, tt.want) {
				t.Errorf("ZipLongest() got %v, want %v", result, tt.want)
			}
		})
	}

}
