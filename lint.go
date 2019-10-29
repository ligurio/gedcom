package main

import (
	"bytes"
	_"errors"
	"flag"
	"fmt"
	"github.com/iand/gedcom"
	"io/ioutil"
	"os"
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

type fn func(person *gedcom.IndividualRecord)

var (
	// rules are borrowed from gigatrees.com and lifelines
	/*
		Persons Born After They Were Baptized
		Persons Born After They Were Married
		Persons Born After Their Children
		Persons Born After They Died
		Persons Born After Being Buried
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
		Persons With Multiple Parents
		Persons Having an Ancestral Loop
		Persons Baptized Past the Age of 5
		Persons Married Before the Age of 14
		Wives Married After the Age of 50
		Persons Living Past the Age of 100
		Persons at Least 30 Years Older Than Their Spouse
		Persons Having Children Before the Age of 15
		Mothers Having Children Past the Age of 50
		Persons Having Similarly Named Children
		Persons Having a Degree of Kinship of 4 or Less
		Persons Having a Non-biological Parent
		Persons Having Unknown Genders
		Families Having Swapped Spouses
		Persons Whose Birth Dates Could Not Be Estimated
		Persons with Invalid Dates
	*/

	// Marriage checks:
	EM100 = "EM100: person marries before birth"
	EM101 = "EM101: person marries after death"
	EM102 = "EM102: person has more than _wedder_ spouses"
	EM103 = "EM103: person marries someone more than _maydec_ years older"
	EM104 = "EM104: person marries younger than _yngmar_"
	EM105 = "EM105: person marries older than _oldmar_"
	EM106 = "EM106: marriage out of order"
	EM107 = "EM107: marriage before birth from previous marriage"
	EM108 = "EM108: marriage after birth from subsequent marriage"
	EM109 = "EM109: homosexual marriage"
	EM110 = "EM110: person is a female husband"
	EM111 = "EM111: person is a male wife"
	EM112 = "EM112: person was a widow(er) longer than _lngwdw_ years"
	EM113 = "EM113: person lived more than _oldunm_ years and never married"
	EM114 = "EM114: person has multiple marriages, this one with no spouse and no children"
	EM115 = "EM115: person has same surname as spouse"

	// Parentage checks:
	EP100 = "EP100: mother has more than _fecmom_ children"
	EP101 = "EP101: mother is older than _oldmom_ at time of birth of child"
	EP102 = "EP102: child is born before mother"
	EP103 = "EP103: mother is younger than _yngmom_"
	EP104 = "EP104: mother is dead at birth of child"
	EP105 = "EP105: same as above, but for father"
	EP106 = "EP106: child doesn't inherit father's surname"

	// Children checks:
	EC100 = "EC100: child is born out of order with respect to a previous child"
	EC101 = "EC101: child is born in the same year as a previous child"
	EC102 = "EC102: child is born more than _cspace_ years after previous child"
	EC103 = "EC103: children's births span more than _cbspan_ years"
	EC104 = "EC104: child is born before parents' marriage"
	EC105 = "EC105: child has same given name as sibling"

	// Family checks:
	EF100 = "EF100: family has no members"
	EF101 = "EF101: family has no parents"
	EF102 = "EF102: husband missing pointer to family"
	EF103 = "EF103: family missing pointer to husband"
	EF104 = "EF104: wife missing pointer to family"
	EF105 = "EF105: family missing pointer to wife"
	EF106 = "EF106: child missing pointer to family"
	EF107 = "EF107: family missing pointer to child"
	EF108 = "EF108: family has multiple husbands"
	EF109 = "EF109: family has multiple wives"
	EF110 = "EF110: child is in family multiple times"
)

func rule_I100(person *gedcom.IndividualRecord) {
	// person's age at death is older than _oldage_
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	deat_date = eventDate(person, "DEAT")
	if deat_date.Sub(birt_date) < OLDAGE {
		//fmt.Println(errors.New(EI100), person.Name)
		fmt.Println(person.Name)
	}
}

func rule_I101(person *gedcom.IndividualRecord) {
	// person is baptized before birth
	// TODO
}

func rule_I102(person *gedcom.IndividualRecord) {
	// person dies before birth
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	deat_date = eventDate(person, "DEAT")
	if deat_date.Sub(birt_date) < 0 {
		fmt.Println()
		//fmt.Println(errors.New(EI102))
	}
}

func rule_I103(person *gedcom.IndividualRecord) {
	// person is buried before birth
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) < 0 {
		fmt.Println()
		//fmt.Println(errors.New(EI103))
	}
}

func rule_I104(person *gedcom.IndividualRecord) {
	// person dies before baptism
	// TODO
}

func rule_I105(person *gedcom.IndividualRecord) {
	// person is buried before baptism
	// TODO
}

func rule_I106(person *gedcom.IndividualRecord) {
	// person is buried before death
	var buri_date, deat_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) > 0 {
		fmt.Println()
		//fmt.Println(errors.New(EI106))
	}
}

func rule_I107(person *gedcom.IndividualRecord) {
	// person is baptised after birth year
	// TODO
}

func rule_I108(person *gedcom.IndividualRecord) {
	// person is buried after death year
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) >= 1 {
		fmt.Println()
		//fmt.Println(errors.New(EI108))
	}
}

func rule_I109(person *gedcom.IndividualRecord) {
	// person has unkown gender
	if person.Sex == "" {
		fmt.Println(person.Name)
		//fmt.Println(errors.New(EI109), person.Name)
	}
}

func rule_I110(person *gedcom.IndividualRecord) {
	// person has ambiguous gender
	// TODO
}

func rule_I111(person *gedcom.IndividualRecord) {
	// person has multiple parentage
	// TODO
}

func rule_I112(person *gedcom.IndividualRecord) {
	// person has no family pointers
	// TODO
}

func rule_F100(family *gedcom.FamilyRecord) {
	// family has no members
}

func rule_F101(family *gedcom.FamilyRecord) {
	// family has no parents
}

func rule_F102(family *gedcom.FamilyRecord) {
	// husband missing pointer to family
}

func rule_F103(family *gedcom.FamilyRecord) {
	// family missing pointer to husband
}

func rule_F104(family *gedcom.FamilyRecord) {
	// wife missing pointer to family
}

func rule_F105(family *gedcom.FamilyRecord) {
	// family missing pointer to wife
}
func rule_F106(family *gedcom.FamilyRecord) {
	// child missing pointer to family"
}

func rule_F107(family *gedcom.FamilyRecord) {
	// family missing pointer to child
}
func rule_F108(family *gedcom.FamilyRecord) {
	// family has multiple husbands
}

func rule_F109(family *gedcom.FamilyRecord) {
	// family has multiple wives
}

func rule_F110(family *gedcom.FamilyRecord) {
	// child is in family multiple times
}

func rule_M100(family *gedcom.FamilyRecord) {
	// person marries before birth
}

func rule_M101(family *gedcom.FamilyRecord) {
	// person marries after death
}
func rule_M102(family *gedcom.FamilyRecord) {
	// person has more than _wedder_ spouses
}
func rule_M103(family *gedcom.FamilyRecord) {
	// person marries someone more than _maydec_ years older
}
func rule_M104(family *gedcom.FamilyRecord) {
	// person marries younger than _yngmar_
}
func rule_M105(family *gedcom.FamilyRecord) {
	// person marries older than _oldmar_
}
func rule_M106(family *gedcom.FamilyRecord) {
	// marriage out of order
}
func rule_M107(family *gedcom.FamilyRecord) {
	// marriage before birth from previous marriage
}
func rule_M108(family *gedcom.FamilyRecord) {
	// marriage after birth from subsequent marriage
}
func rule_M109(family *gedcom.FamilyRecord) {
	// homosexual marriage
}
func rule_M110(family *gedcom.FamilyRecord) {
	// person is a female husband
}
func rule_M111(family *gedcom.FamilyRecord) {
	// person is a male wife
}
func rule_M112(family *gedcom.FamilyRecord) {
	// person was a widow(er) longer than _lngwdw_ years
}
func rule_M113(family *gedcom.FamilyRecord) {
	// person lived more than _oldunm_ years and never married
}
func rule_M114(family *gedcom.FamilyRecord) {
	// person has multiple marriages, this one with no spouse and no children
}
func rule_M115(family *gedcom.FamilyRecord) {
	// person has same surname as spouse
}

func rule_P100(family *gedcom.FamilyRecord) {
	// mother has more than _fecmom_ children
}

func rule_P101(family *gedcom.FamilyRecord) {
	// mother is older than _oldmom_ at time of birth of child
}
func rule_P102(family *gedcom.FamilyRecord) {
	// child is born before mother
}
func rule_P103(family *gedcom.FamilyRecord) {
	// mother is younger than _yngmom_
}
func rule_P104(family *gedcom.FamilyRecord) {
	// mother is dead at birth of child
}
func rule_P105(family *gedcom.FamilyRecord) {
	// same as above, but for father
}
func rule_P106(family *gedcom.FamilyRecord) {
	// child doesn't inherit father's surname
}

func rule_C100(family *gedcom.FamilyRecord) {
	// child is born out of order with respect to a previous child
}

func rule_C101(family *gedcom.FamilyRecord) {
	// child is born in the same year as a previous child
}

func rule_C102(family *gedcom.FamilyRecord) {
	// child is born more than _cspace_ years after previous child
}

func rule_C103(family *gedcom.FamilyRecord) {
	// children's births span more than _cbspan_ years
}

func rule_C104(family *gedcom.FamilyRecord) {
	// child is born before parents' marriage
}

func rule_C105(family *gedcom.FamilyRecord) {
	// child has same given name as sibling
}

func eventDate(person *gedcom.IndividualRecord, eventTag string) time.Time {
	var date time.Time
	for _, event := range person.Event {
		if event.Tag == eventTag {
			date, _ = parse(event.Date)
		}
	}

	return date
}

func parse(date string) (time.Time, error) {
	// "2006-01-02T15:04:05.000Z"
	// http://www.gedcomx.org/GEDCOM-5.5.1.pdf
	layout := "02 JAN 2006"
	t, err := time.Parse(layout, date)
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}

func main() {

	rules := map[string]fn{
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nlint is a tool for checking of GEDCOM to errors.")
		fmt.Fprintf(os.Stderr, "\n\nFlags:\n")
		flag.PrintDefaults()
	}

	var gedfile = flag.String("file", "", "GEDCOM filename")
	var ignorelist = flag.String("ignore", "", "errors ignore list")
	var verbose = flag.Bool("verbose", false, "verbose mode")

	flag.Parse()

	if *gedfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	var ignores []string
	ignores = strings.Split(*ignorelist, ",")
	fmt.Println("Errors to ignore:", ignores)

	data, _ := ioutil.ReadFile(*gedfile)
	d := gedcom.NewDecoder(bytes.NewReader(data))
	g, _ := d.Decode()

	if *verbose {
		fmt.Println("\nPersons (Xref, Name, Sex):")
		fmt.Println("--------------------------")
		fmt.Println("Found persons:", len(g.Individual))
	}

	for _, rec := range g.Individual {
		if len(rec.Name) > 0 {
			if *verbose {
				fmt.Printf("%s, %s, %s\n", rec.Xref, rec.Name[0].Name, rec.Sex)
			}
		}

		for err_msg, fn := range rules {
			fmt.Printf("key[%s] value[%s]\n", err_msg, fn)
			fn(rec)
		}
	}

	if *verbose {
		fmt.Println("\nFamily (Xref, Husband, Wife):")
		fmt.Println("-------------------------------")
		fmt.Println("Found families:", len(g.Family))
	}
	for _, rec := range g.Family {
		if *verbose {
			fmt.Printf("%s\n", rec.Xref)
			fmt.Printf("%s, %s, %s\n", rec.Xref, rec.Husband.Xref, rec.Wife.Xref)
		}

		/*
			rule_F100(rec)
			rule_F101(rec)
			rule_F102(rec)
			rule_F103(rec)
			rule_F104(rec)
			rule_F105(rec)
			rule_F106(rec)
			rule_F107(rec)
			rule_F108(rec)
			rule_F109(rec)
			rule_F110(rec)
		*/

		for _, child := range rec.Child {
			if *verbose {
				fmt.Printf("\t%s, %s\n", rec.Xref, child.Xref)
			}
		}
	}
}
