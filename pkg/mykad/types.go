package mykad

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// All place of birth mappings are obtained from https://en.wikipedia.org/wiki/Malaysian_identity_card#Place_of_birth

// TODO(shiraaz): This was copy pasted from somwhere and needs to be cleaned up.
const REGION_SOUTHEAST_ASIA = "SOUTHEAST_ASIA"
const REGION_BRITISH_ISLES = "BRITISH_ISLES"
const REGION_SOVIET_REPUBLIC = "SOVIET_REPUBLIC"
const REGION_EAST_ASIA = "EAST_ASIA"
const REGION_SOUTH_ASIA = "SOUTH_ASIA"
const REGION_AFRICA = "AFRICA"
const REGION_SOUTH_AMERICA = "SOUTH_AMERICA"
const REGION_CENTRAL_AMERICA = "CENTRAL_AMERICA"
const REGION_NORTH_AMERICA = "NORTH_AMERICA"
const REGION_OCEANIA = "OCEANIA"
const REGION_MIDDLE_EAST = "MIDDLE_EAST"
const REGION_EUROPE = "EUROPE"
const REGION_MISCELLANEOUS = "MISCELLANEOUS"

type region struct {
	country string
	region  string
}

// TODO(shiraaz): This was copy pasted from somwhere and needs to be cleaned up.
var countryCodePairs = map[string]region{
	"60": {country: "BN", region: REGION_SOUTHEAST_ASIA},
	"61": {country: "ID", region: REGION_SOUTHEAST_ASIA},
	"62": {country: "KH", region: REGION_SOUTHEAST_ASIA},
	"63": {country: "LA", region: REGION_SOUTHEAST_ASIA},
	"64": {country: "MM", region: REGION_SOUTHEAST_ASIA},
	"65": {country: "PH", region: REGION_SOUTHEAST_ASIA},
	"66": {country: "SG", region: REGION_SOUTHEAST_ASIA},
	"67": {country: "TH", region: REGION_SOUTHEAST_ASIA},
	"68": {country: "VN", region: REGION_SOUTHEAST_ASIA},

	"71": {country: "FOREIGN_UNKNOWN", region: ""},
	"72": {country: "FOREIGN_UNKNOWN", region: ""},

	"74": {country: "CN", region: REGION_EAST_ASIA},
	"75": {country: "IN", region: REGION_SOUTH_ASIA},
	"76": {country: "PK", region: REGION_SOUTH_ASIA},
	"77": {country: "SA", region: REGION_MIDDLE_EAST},
	"78": {country: "LK", region: REGION_SOUTH_ASIA},
	"79": {country: "BD", region: REGION_SOUTH_ASIA},

	"83": {
		country: "AS|AU|CX|CC|CK|FJ|PF|GU|HM|MH|FM|NC|NZ|NU|NF|PG|TL|TK|UM|WF",
		region:  REGION_OCEANIA},
	"84": {
		country: "AI|AR|AW|BO|BR|CL|CO|EC|GF|GP|GY|PY|PE|GS|ST|UY|VE",
		region:  REGION_SOUTH_AMERICA},
	"85": {
		country: "DZ|AO|BW|BI|CM|CF|CG|CD|DG|EG|ER|ET|GA|GM|GN|KE|LR|MW|ML|MR|YT|" +
			"MA|MZ|NA|NE|NG|RW|RE|SN|SL|SO|SD|SZ|TZ|TG|TO|TN|UG|ME|ZR|ZM|ZW",
		region: REGION_AFRICA},
	"86": {
		country: "AM|AT|BE|CY|DK|FO|FR|FI|DE|DD|GR|VA|IT|LU|" +
			"MK|MT|MC|NL|NO|PT|MD|SK|SI|ES|SE|CH|GG|JE|IM",
		region: REGION_EUROPE},
	"87": {
		country: "GB|IE",
		region:  REGION_BRITISH_ISLES},
	"88": {
		country: "BH|IR|IQ|PS|JO|KW|OM|QA|YE|SY|TR|YE|YD|",
		region:  REGION_MIDDLE_EAST},
	"89": {
		country: "JP|KP|KR|TW",
		region:  REGION_EAST_ASIA},
	"90": {
		country: "BS|BB|BZ|CR|CU|DM|DO|SV|GD|GT|HT|HN|" +
			"JM|MQ|MX|NI|PA|PR|KN|LC|VC|TT|TC|VI",
		region: REGION_CENTRAL_AMERICA},
	"91": {
		country: "CA|GL|AN|PM|US",
		region:  REGION_NORTH_AMERICA},
	"92": {
		country: "AL|BY|BA|BG|HR|CZ|CS|EE|GE|HU|LV|LT|ME|PL|XK|RO|RU|RS|UA",
		region:  REGION_SOVIET_REPUBLIC},
	"93": {
		country: "AF|AD|AQ|AG|AZ|BJ|BM|BT|IO|BF|CV|KY|KM|DY|GQ|TF|GI|GW|HK|" +
			"IS|CI|KZ|KI|KG|LS|LY|LI|MO|MG|MV|MU|MN|MS|NR|NP|MP|PW|PS|" +
			"PN|SH|LC|VC|WS|SM|ST|SC|SB|SJ|TJ|TM|TV|HV|UZ|VU|VA|VG|YU",
		region: REGION_MISCELLANEOUS},

	"98": {country: "STATELESS", region: ""},
	"99": {country: "UNSPECIFIED", region: ""},
}

type PlaceOfBirth struct {
	CountryISO string
	Location   string
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

// NewMyKAD returns a new malaysian identity from a provided NRIC number.
func NewMyKAD(nric string) (*MyKAD, error) {
	r := regexp.MustCompile(`^(\d{6})-?(\d{2})-?(\d{3})(\d{1})$`)
	s := r.FindStringSubmatch(nric)

	if len(s) != 5 {
		return nil, errors.New("invalid mykad number")
	}

	dob, err := time.Parse("060102", s[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	pob, err := strconv.Atoi(s[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing location: %v", err)
	}

	// TODO: This should go into a isMalaysian and isForeigner check.
	var c CitizenType
	var pb PlaceOfBirth
	if (pob > 1 && pob < 17) || (pob > 20 && pob < 60) {
		c = CitizenTypeMalaysian
		pb = PlaceOfBirth{
			CountryISO: "MY",
			Location:   "",
		}
	} else if p, ok := countryCodePairs[s[2]]; ok {
		c = CitizenTypeForeigner
		pb = PlaceOfBirth{
			CountryISO: p.country,
			Location:   p.region,
		}
	} else {
		return nil, errors.New("invalid location")
	}

	gender := GenderMale
	g, err := strconv.Atoi(s[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing gender: %v", err)
	}

	if g%2 == 0 {
		gender = GenderFemale
	}

	return &MyKAD{
		NRIC:         nric,
		DateOfBirth:  dob,
		PlaceOfBirth: pb,
		CitizenType:  c,
		Gender:       gender,
	}, nil
}
