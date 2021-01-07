package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/iand/gedcom"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	OLDAGE = 99
	WEDDER = 3
	MAYDEC = 5
	YNGMAR = 20
	OLDMAR = 80
	LNGWDW = 80
	OLDUNM = 67
	FECMOM = 45
	OLDMOM = 55
	YNGMOM = 16
	CSPACE = 1
	CBSPAN = 1
)

type fn_person func(person *gedcom.IndividualRecord) bool
type fn_family func(family *gedcom.FamilyRecord) bool

func rule_I100(person *gedcom.IndividualRecord) bool {
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	if birt_date.IsZero() {
		return true
	}
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	if deat_date.Sub(birt_date) > OLDAGE {
		fmt.Println(deat_date, birt_date)
		return false
	}

	return true
}

func rule_I101(person *gedcom.IndividualRecord) bool {
	var birt_date, bap_date time.Time
	birt_date = eventDate(person, "BIRT")
	if birt_date.IsZero() {
		return true
	}
	bap_date = eventDate(person, "BAP")
	if bap_date.IsZero() {
		return true
	}
	if bap_date.Sub(birt_date) > 0 {
		fmt.Println(birt_date, bap_date)
		return false
	}

	return true
}

func rule_I102(person *gedcom.IndividualRecord) bool {
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	if birt_date.IsZero() {
		return true
	}
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	if deat_date.Sub(birt_date) < 0 {
		fmt.Println(birt_date, deat_date)
		return false
	}

	return true
}

func rule_I103(person *gedcom.IndividualRecord) bool {
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	buri_date = eventDate(person, "BURI")
	if buri_date.IsZero() {
		return true
	}
	if deat_date.Sub(buri_date) < 0 {
		fmt.Println(deat_date, buri_date)
		return false
	}

	return true
}

func rule_I104(person *gedcom.IndividualRecord) bool {
	var deat_date, bap_date time.Time
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	bap_date = eventDate(person, "BAP")
	if bap_date.IsZero() {
		return true
	}
	if deat_date.Sub(bap_date) < 0 {
		fmt.Println(deat_date, bap_date)
		return false
	}

	return true
}

func rule_I105(person *gedcom.IndividualRecord) bool {
	var bap_date, buri_date time.Time
	bap_date = eventDate(person, "BAP")
	if bap_date.IsZero() {
		return true
	}
	buri_date = eventDate(person, "BURI")
	if buri_date.IsZero() {
		return true
	}
	if bap_date.Sub(buri_date) < 0 {
		fmt.Println(buri_date, bap_date)
		return false
	}

	return true
}

func rule_I106(person *gedcom.IndividualRecord) bool {
	var buri_date, deat_date time.Time
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	buri_date = eventDate(person, "BURI")
	if buri_date.IsZero() {
		return true
	}
	if deat_date.Sub(buri_date) > 0 {
		fmt.Println(buri_date, deat_date)
		return false
	}

	return true
}

func rule_I107(person *gedcom.IndividualRecord) bool {
	var bap_date, birt_date time.Time
	bap_date = eventDate(person, "BAP")
	if bap_date.IsZero() {
		return true
	}
	birt_date = eventDate(person, "BIRT")
	if birt_date.IsZero() {
		return true
	}
	if bap_date.Sub(birt_date) >= 1 {
		fmt.Println(bap_date, birt_date)
		return false
	}

	return true
}

func rule_I108(person *gedcom.IndividualRecord) bool {
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	if deat_date.IsZero() {
		return true
	}
	buri_date = eventDate(person, "BURI")
	if buri_date.IsZero() {
		return true
	}
	if deat_date.Sub(buri_date) >= 1 {
		fmt.Println(deat_date, buri_date)
		return false
	}

	return true
}

func rule_I109(person *gedcom.IndividualRecord) bool {
	if person.Sex == "" {
		return false
	}

	return true
}

func rule_I110(person *gedcom.IndividualRecord) bool {
	/* FIXME */
	return true
}

func rule_I111(person *gedcom.IndividualRecord) bool {
	if len(person.Parents) > 1 {
		return false
	}

	return true
}

func rule_I112(person *gedcom.IndividualRecord) bool {
	if len(person.Family) == 0 {
		return false
	}

	return true
}

func rule_F100(family *gedcom.FamilyRecord) bool {
	if family.Husband == nil {
		return false
	}
	if family.Wife == nil {
		return false
	}
	if len(family.Child) == 0 {
		return false
	}

	return true
}

func rule_F101(family *gedcom.FamilyRecord) bool {
	if family.Husband == nil {
		return false
	}
	if family.Wife == nil {
		return false
	}

	return true
}

func rule_F102(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_F103(family *gedcom.FamilyRecord) bool {
	if family.Husband == nil {
		return false
	}

	return true
}

func rule_F104(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_F105(family *gedcom.FamilyRecord) bool {
	if family.Wife == nil {
		return false
	}

	return true
}

func rule_F106(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_F107(family *gedcom.FamilyRecord) bool {
	if len(family.Child) == 0 {
		return false
	}

	return true
}

func rule_F108(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_F109(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_F110(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M100(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M101(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M102(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M103(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M104(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M105(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M106(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M107(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M108(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M109(family *gedcom.FamilyRecord) bool {
	if family.Husband.Sex == family.Wife.Sex {
		return false
	}

	return true
}

func rule_M110(family *gedcom.FamilyRecord) bool {
	if family.Husband.Sex == "F" {
		return false
	}

	return true
}

func rule_M111(family *gedcom.FamilyRecord) bool {
	if family.Wife.Sex == "M" {
		return false
	}

	return true
}

func rule_M112(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M113(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M114(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_M115(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_P100(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_P101(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_P102(family *gedcom.FamilyRecord) bool {
	var mother_birt_date, child_birt_date time.Time
	mother_birt_date = eventDate(family.Wife, "BIRT")
	for _, child := range family.Child {
		child_birt_date = eventDate(child, "BIRT")
		if mother_birt_date.Sub(child_birt_date) < 0 {
			return false
		}
	}

	return true
}

func rule_P103(family *gedcom.FamilyRecord) bool {
	var mother_birt_date, child_birt_date time.Time
	mother_birt_date = eventDate(family.Wife, "BIRT")
	if mother_birt_date.IsZero() {
		return true
	}
	for _, child := range family.Child {
		child_birt_date = eventDate(child, "BIRT")
		if child_birt_date.IsZero() {
			continue
		}
		if mother_birt_date.Sub(child_birt_date) < YNGMOM {
			fmt.Println(mother_birt_date, child_birt_date)
			return false
		}
	}

	return true
}

func rule_P104(family *gedcom.FamilyRecord) bool {
	var mother_deat_date, child_birt_date time.Time
	mother_deat_date = eventDate(family.Wife, "DEAT")
	if mother_deat_date.IsZero() {
		return true
	}
	for _, child := range family.Child {
		child_birt_date = eventDate(child, "BIRT")
		if child_birt_date.IsZero() {
			continue
		}
		if mother_deat_date == child_birt_date {
			fmt.Println(mother_deat_date, child_birt_date)
			return false
		}
	}

	return true
}

func rule_P105(family *gedcom.FamilyRecord) bool {
	var father_deat_date, child_birt_date time.Time
	father_deat_date = eventDate(family.Husband, "DEAT")
	if father_deat_date.IsZero() {
		return true
	}
	for _, child := range family.Child {
		child_birt_date = eventDate(child, "BIRT")
		if child_birt_date.IsZero() {
			continue
		}
		if father_deat_date == child_birt_date {
			fmt.Println(father_deat_date, child_birt_date)
			return false
		}
	}

	return true
}

func rule_P106(family *gedcom.FamilyRecord) bool {
	var father_surname []string
	var child_surname []string
	surname_re := regexp.MustCompile(`.+ /(.+)/`)
	for _, name := range family.Husband.Name {
		surname := surname_re.FindStringSubmatch(name.Name)
		if surname[1] != "" {
			father_surname = append(father_surname, surname[1])
		}
	}
	for _, child := range family.Child {
		for _, name := range child.Name {
			surname := surname_re.FindStringSubmatch(name.Name)
			if surname[1] != "" {
				child_surname = append(child_surname, surname[1])
			}
		}
	}

	for _, surname := range child_surname {
		if !contains(father_surname, surname) {
			fmt.Println(father_surname, surname)
			return false
		}
	}

	return true
}

func rule_C100(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_C101(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_C102(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_C103(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_C104(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func rule_C105(family *gedcom.FamilyRecord) bool {
	/* FIXME */
	return true
}

func eventDate(person *gedcom.IndividualRecord, eventTag string) time.Time {
	var date time.Time
	for _, event := range person.Event {
		if event.Tag == eventTag {
			date, _ = parse_date_string(event.Date)
		}
	}

	return date
}

func parse_date_string(date string) (time.Time, error) {
	// "2006-01-02T15:04:05.000Z"
	layout := "02 JAN 2006"
	t, err := time.Parse(layout, date)
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PrintIndividualRecord(record *gedcom.IndividualRecord) {
	fmt.Printf("Person (%s)", record.Xref)
	if len(record.Name) > 0 {
		fmt.Printf(" %s\n", record.Name[0].Name)
	} else {
		fmt.Printf("\n")
	}
}

func PrintFamilyRecord(record *gedcom.FamilyRecord) {
	fmt.Printf("Family (%s)\n", record.Xref)
}

var (
	/*
		persons born after they were married
		persons born after their children
		persons having similarly named children
		persons having an ancestral loop

		Persons Born After One of Their Parents Died
		Persons Born After One of Their Parents Was Buried
		Persons Baptized After Being Buried
		Persons Baptized After Being Married
		Persons Married After Being Buried
		Persons Baptized After They Died
		Persons Married After They Died
		Persons Died Before Children Were Born
		Persons Died After Being Buried
		Persons Buried Before Having Children
		Persons Baptized Past the Age of 5
		Persons Married Before the Age of 14
		Wives Married After the Age of 50
		Persons Living Past the Age of 100
		Persons at Least 30 Years Older Than Their Spouse
		Persons Having a Degree of Kinship of 4 or Less
		Persons Having a Non-biological Parent
		Persons Whose Birth Dates Could Not Be Estimated
		Persons with Invalid Dates
		Families Having Swapped Spouses
	*/

	// rules are borrowed from gigatrees.com and lifelines
	person_rules = map[string]fn_person{
		"EI100: person's age at death is older than _oldage_": rule_I100,
		"EI101: person is baptized before birth":              rule_I101,
		"EI102: person dies before birth":                     rule_I102,
		"EI103: person is buried before birth":                rule_I103,
		"EI104: person dies before baptism":                   rule_I104,
		"EI105: person is buried before baptism":              rule_I105,
		"EI106: person is buried before death":                rule_I106,
		"EI107: person is baptised after birth year":          rule_I107,
		"EI108: person is buried after death year":            rule_I108,
		"EI109: person has unkown gender":                     rule_I109,
		"EI110: person has ambiguous gender":                  rule_I110,
		"EI111: person has multiple parentage":                rule_I111,
		"EI112: person has no family pointers":                rule_I112,
	}

	family_rules = map[string]fn_family{
		"EF100: family has no members":             rule_F100,
		"EF101: family has no parents":             rule_F101,
		"EF102: husband missing pointer to family": rule_F102,
		"EF103: family missing pointer to husband": rule_F103,
		"EF104: wife missing pointer to family":    rule_F104,
		"EF105: family missing pointer to wife":    rule_F105,
		"EF106: child missing pointer to family":   rule_F106,
		"EF107: family missing pointer to child":   rule_F107,
		"EF108: family has multiple husbands":      rule_F108,
		"EF109: family has multiple wives":         rule_F109,
		"EF110: child is in family multiple times": rule_F110,
		// marriage rules
		"EM100: person marries before birth":                                            rule_M100,
		"EM101: person marries after death":                                             rule_M101,
		"EM102: person has more than WEDDER spouses":                                    rule_M102,
		"EM103: person marries someone more than MAYDEC years older":                    rule_M103,
		"EM104: person marries younger than YNGMAR":                                     rule_M104,
		"EM105: person marries older than OLDMAR":                                       rule_M105,
		"EM106: marriage out of order":                                                  rule_M106,
		"EM107: marriage before birth from previous marriage":                           rule_M107,
		"EM108: marriage after birth from subsequent marriage":                          rule_M108,
		"EM109: homosexual marriage":                                                    rule_M109,
		"EM110: person is a female husband":                                             rule_M110,
		"EM111: person is a male wife":                                                  rule_M111,
		"EM112: person was a widow(er) longer than LNGWDW years":                        rule_M112,
		"EM113: person lived more than OLDUNM years and never married":                  rule_M113,
		"EM114: person has multiple marriages, this one with no spouse and no children": rule_M114,
		"EM115: person has same surname as spouse":                                      rule_M115,
		// parentage rules
		"EP100: mother has more than FECMOM children":                    rule_P100,
		"EP101: mother is older than OLDMOM at time of birth of child":   rule_P101,
		"EP102: child is born before mother":                             rule_P102,
		"EP103: mother is younger than YNGMOM at time of birth of child": rule_P103,
		"EP104: mother is dead at birth of child":                        rule_P104,
		"EP105: father is dead at birth of child":                        rule_P105,
		"EP106: child doesn't inherit father's surname":                  rule_P106,
		// children rules
		"EC100: child is born out of order with respect to a previous child": rule_C100,
		"EC101: child is born in the same year as a previous child":          rule_C101,
		"EC102: child is born more than CSPACE years after previous child":   rule_C102,
		"EC103: children's births span more than CBSPAN years":               rule_C103,
		"EC104: child is born before parent's marriage":                      rule_C104,
		"EC105: child has same given name as sibling":                        rule_C105,
	}
)

func printErrors(fn *string, ignoresList *string, verbose bool) {

	var ignores []string
	if len(*ignorelist) != 0 {
		ignores = strings.Split(*ignorelist, ",")
		fmt.Println("These rules are ignored:", ignores)
	}

	data, _ := ioutil.ReadFile(*gedfile)
	d := gedcom.NewDecoder(bytes.NewReader(data))
	g, _ := d.Decode()

	if *verbose {
		fmt.Printf("Found %d persons:\n\n", len(g.Individual))
	}

	found_errors := false
	for _, record := range g.Individual {
		for err_msg, fn := range person_rules {
			if contains(ignores, err_msg) {
				continue
			}
			if !fn(record) {
				PrintIndividualRecord(record)
				fmt.Println(err_msg)
				found_errors = true
			}
		}
	}

	if *verbose {
		fmt.Printf("\nFound %d families:\n\n", len(g.Family))
	}
	for _, record := range g.Family {
		for err_msg, fn := range family_rules {
			if contains(ignores, err_msg) {
				continue
			}
			if !fn(record) {
				PrintFamilyRecord(record)
				fmt.Println(err_msg)
				found_errors = true
			}
		}
	}

	if found_errors {
		os.Exit(1)
	}
}
