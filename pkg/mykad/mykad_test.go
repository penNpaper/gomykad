package mykad

import (
	"fmt"
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
				fmt.Println("Result ", res)
			}
		})
	}
}
