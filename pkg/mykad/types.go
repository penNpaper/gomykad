package mykad

import (
	"time"
)

// PlaceOfBirth represents the literal place of birth of the MyKAD holder.
type PlaceOfBirth struct {
	PossibleCountry string
	Location        string
}

// CitizenType is the citizenship that the MyKAD holder holds in Malaysia.
type CitizenType int

const (
	// CitizenTypeMalaysian is a Malaysian born citizen.
	CitizenTypeMalaysian CitizenType = 1

	// CitizenTypeForeigner is a Malaysian MyKAD holder who was born in a country other than Malaysia.
	CitizenTypeForeigner CitizenType = 2
)

// Gender is the gender of the MyKAD holder.
type Gender int

const (
	// GenderMale is a male.
	GenderMale   Gender = 1

	// GenderFemale is a female.
	GenderFemale Gender = 2
)

// MyKAD is a typed representation of a MyKAD holder. It includes the original NRIC numer.
type MyKAD struct {
	NRIC         string
	DateOfBirth  time.Time
	PlaceOfBirth PlaceOfBirth
	Gender       Gender
	CitizenType  CitizenType
}

type placeOfBirthCode int

func (pob placeOfBirthCode) isMalaysianCode() bool {
	i := int(pob)
	return (i >= 1 && i <= 16) || (i >= 21 && i <= 59)
}

func (pob placeOfBirthCode) isForeignerCode() bool {
	i := int(pob)
	return (i >= 60 && i <= 68) || i == 71 || i == 72 ||
		(i >= 74 && i <= 79) || (i >= 82 && i <= 93) ||
		i == 98 || i == 99
}

func (pob placeOfBirthCode) valid() bool {
	return pob.isMalaysianCode() || pob.isForeignerCode()
}
