package mykad

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name       string
		sampleSize int
	}{
		{"Run of ten", 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < test.sampleSize; i++ {
				res := Generate()
				fmt.Printf("%v\n", res)
				myKAD, err := NewMyKAD(res)
				require.NoError(t, err)
				fmt.Printf("%v\n", myKAD)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		mykad       string
		expectError bool
	}{
		{"Valid mykad", "840701-79-3686", false},
		{"Invalid mykad - wrong date", "0701-79-3686", true},
		{"Invalid mykad - invalid place of birth", "840701-69-3686", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Validate(test.mykad)
			if test.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
