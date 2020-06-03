package mykad

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const minAge = 12
const maxAge = 112 // Based off guinness world record.

// Generate will return a new random MyKAD number.
func Generate() string {
	// Generate a random date for the year component.
	n := time.Now()
	rand.Seed(n.UnixNano())

	e := n.AddDate(-minAge, 0, 0).Unix()
	s := n.AddDate(-maxAge, 0, 0).Unix()

	sec := rand.Int63n(e-s) + s
	ds := time.Unix(sec, 0).Format("060102")

	// Generate a random place.
	p := rand.Intn(99)

	// Generate a special number.
	sn := rand.Intn(9999)

	return fmt.Sprintf("%v-%02d-%04d", ds, p, sn)
}

// Validate will validate a mykad number and return an error for
// an invalid mykad number.
func Validate(mykad string) error {
	r := regexp.MustCompile(`^(\d{6})-?(\d{2})-?(\d{3})(\d{1})$`)
	s := r.FindStringSubmatch(mykad)

	if len(s) != 5 {
		return errors.New("invalid mykad number")
	}

	_, err := time.Parse("060102", s[1])
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	pob, err := strconv.Atoi(s[2])
	if err != nil {
		return fmt.Errorf("error parsing location: %v", err)
	}

	// TODO: This should go into a isMalaysian and isForeigner check.
	if (pob > 1 && pob < 16) || (pob > 21 && pob < 56) {
	} else if _, ok := countryCodePairs[s[2]]; ok {
	} else {
		return errors.New("invalid location")
	}

	fmt.Printf("%v", s)
	return nil
}
