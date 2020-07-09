package mykad

import (
	"time"
)

type PlaceOfBirth struct {
	PossibleCountry string
	Location        string
}

type CitizenType int

const (
	CitizenTypeMalaysian CitizenType = 1
	CitizenTypeForeigner CitizenType = 2
)

type Gender int

const (
	GenderMale   Gender = 1
	GenderFemale Gender = 2
)

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
