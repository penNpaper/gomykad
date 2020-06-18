package mykad

import (
	"fmt"
	"math/rand"
	"time"
)

const minAge = 12
const maxAge = 112 // Based off guinness world record.

// Generate will return a new random MyKAD number.
// TODO: This may break validate.
func Generate() string {
	// Generate a random date for the year component.
	n := time.Now()
	rand.Seed(n.UnixNano())

	e := n.AddDate(-minAge, 0, 0).Unix()
	s := n.AddDate(-maxAge, 0, 0).Unix()

	sec := rand.Int63n(e-s) + s
	ds := time.Unix(sec, 0).Format("060102")

	// Generate a random place.
	var p int
	// TODO(shiraaz): Simplify
	for p == 0 || (p > 16 && p < 21) || (p > 68 && p < 71) || p == 73 || (p > 79 && p < 82) || (p > 93 && p < 98) {
		p = rand.Intn(99)
	}

	// Generate a special number.
	sn := rand.Intn(9999)

	return fmt.Sprintf("%v-%02d-%04d", ds, p, sn)
}

// Validate will validate a NRIC number.
func Validate(nric string) error {
	_, err := NewMyKAD(nric)
	return err
}
