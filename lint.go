package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/iand/gedcom"
	"io/ioutil"
	"os"
	"reflect"
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

var (

	// rules are borrowed from gigatrees.com and lifelines

	// Individual checks:
	EI100 = "EI100: person's age at death is older than _oldage_"
	EI101 = "EI101: person is baptized before birth"
	EI102 = "EI102: person dies before birth"
	EI103 = "EI103: person is buried before birth"
	EI104 = "EI104: person dies before baptism"
	EI105 = "EI105: person is buried before baptism"
	EI106 = "EI106: person is buried before death"
	EI107 = "EI107: person is baptised after birth year"
	EI108 = "EI108: person is buried after death year"
	EI109 = "EI109: person has unkown gender"
	EI110 = "EI110: person has ambiguous gender"
	EI111 = "EI111: person has multiple parentage"
	EI112 = "EI112: person has no family pointers"

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

func lint_I100(person *gedcom.IndividualRecord) {
	// person's age at death is older than _oldage_
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	deat_date = eventDate(person, "DEAT")
	if deat_date.Sub(birt_date) < OLDAGE {
		fmt.Println(errors.New(EI100), person.Name)
	}
}

func lint_I101(person *gedcom.IndividualRecord) {
	// person is baptized before birth
	// TODO
}

func lint_I102(person *gedcom.IndividualRecord) {
	// person dies before birth
	var birt_date, deat_date time.Time
	birt_date = eventDate(person, "BIRT")
	deat_date = eventDate(person, "DEAT")
	if deat_date.Sub(birt_date) < 0 {
		fmt.Println(errors.New(EI102))
	}
}

func lint_I103(person *gedcom.IndividualRecord) {
	// person is buried before birth
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) < 0 {
		fmt.Println(errors.New(EI103))
	}
}

func lint_I104(person *gedcom.IndividualRecord) {
	// person dies before baptism
	// TODO
}

func lint_I105(person *gedcom.IndividualRecord) {
	// person is buried before baptism
	// TODO
}

func lint_I106(person *gedcom.IndividualRecord) {
	// person is buried before death
	var buri_date, deat_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) > 0 {
		fmt.Println(errors.New(EI106))
	}
}

func lint_I107(person *gedcom.IndividualRecord) {
	// person is baptised after birth year
	// TODO
}

func lint_I108(person *gedcom.IndividualRecord) {
	// person is buried after death year
	var deat_date, buri_date time.Time
	deat_date = eventDate(person, "DEAT")
	buri_date = eventDate(person, "BURI")
	if deat_date.Sub(buri_date) >= 1 {
		fmt.Println(errors.New(EI108))
	}
}

func lint_I109(person *gedcom.IndividualRecord) {
	// person has unkown gender
	if (person.Sex == "") {
		fmt.Println(errors.New(EI109), person.Name)
	}
}

func lint_I110(person *gedcom.IndividualRecord) {
	// person has ambiguous gender
	// TODO
}

func lint_I111(person *gedcom.IndividualRecord) {
	// person has multiple parentage
	// TODO
}

func lint_I112(person *gedcom.IndividualRecord) {
	// person has no family pointers
	// TODO
}

func lint_F100(family *gedcom.FamilyRecord) {
	// family has no members
}

func lint_F101(family *gedcom.FamilyRecord) {
	// family has no parents
}

func lint_F102(family *gedcom.FamilyRecord) {
	// husband missing pointer to family
}

func lint_F103(family *gedcom.FamilyRecord) {
	// family missing pointer to husband
}

func lint_F104(family *gedcom.FamilyRecord) {
	// wife missing pointer to family
}

func lint_F105(family *gedcom.FamilyRecord) {
	// family missing pointer to wife
}
func lint_F106(family *gedcom.FamilyRecord) {
	// child missing pointer to family"
}

func lint_F107(family *gedcom.FamilyRecord) {
	// family missing pointer to child
}
func lint_F108(family *gedcom.FamilyRecord) {
	// family has multiple husbands
}

func lint_F109(family *gedcom.FamilyRecord) {
	// family has multiple wives
}

func lint_F110(family *gedcom.FamilyRecord) {
	// child is in family multiple times
}

func lint_M100(family *gedcom.FamilyRecord) {
	// person marries before birth
}

func lint_M101(family *gedcom.FamilyRecord) {
	// person marries after death
}
func lint_M102(family *gedcom.FamilyRecord) {
	// person has more than _wedder_ spouses
}
func lint_M103(family *gedcom.FamilyRecord) {
	// person marries someone more than _maydec_ years older
}
func lint_M104(family *gedcom.FamilyRecord) {
	// person marries younger than _yngmar_
}
func lint_M105(family *gedcom.FamilyRecord) {
	// person marries older than _oldmar_
}
func lint_M106(family *gedcom.FamilyRecord) {
	// marriage out of order
}
func lint_M107(family *gedcom.FamilyRecord) {
	// marriage before birth from previous marriage
}
func lint_M108(family *gedcom.FamilyRecord) {
	// marriage after birth from subsequent marriage
}
func lint_M109(family *gedcom.FamilyRecord) {
	// homosexual marriage
}
func lint_M110(family *gedcom.FamilyRecord) {
	// person is a female husband
}
func lint_M111(family *gedcom.FamilyRecord) {
	// person is a male wife
}
func lint_M112(family *gedcom.FamilyRecord) {
	// person was a widow(er) longer than _lngwdw_ years
}
func lint_M113(family *gedcom.FamilyRecord) {
	// person lived more than _oldunm_ years and never married
}
func lint_M114(family *gedcom.FamilyRecord) {
	// person has multiple marriages, this one with no spouse and no children
}
func lint_M115(family *gedcom.FamilyRecord) {
	// person has same surname as spouse
}

func lint_P100(family *gedcom.FamilyRecord) {
	// mother has more than _fecmom_ children
}

func lint_P101(family *gedcom.FamilyRecord) {
	// mother is older than _oldmom_ at time of birth of child
}
func lint_P102(family *gedcom.FamilyRecord) {
	// child is born before mother
}
func lint_P103(family *gedcom.FamilyRecord) {
	// mother is younger than _yngmom_
}
func lint_P104(family *gedcom.FamilyRecord) {
	// mother is dead at birth of child
}
func lint_P105(family *gedcom.FamilyRecord) {
	// same as above, but for father
}
func lint_P106(family *gedcom.FamilyRecord) {
	// child doesn't inherit father's surname
}

func lint_C100(family *gedcom.FamilyRecord) {
	// child is born out of order with respect to a previous child
}

func lint_C101(family *gedcom.FamilyRecord) {
	// child is born in the same year as a previous child
}

func lint_C102(family *gedcom.FamilyRecord) {
	// child is born more than _cspace_ years after previous child
}

func lint_C103(family *gedcom.FamilyRecord) {
	// children's births span more than _cbspan_ years
}

func lint_C104(family *gedcom.FamilyRecord) {
	// child is born before parents' marriage
}

func lint_C105(family *gedcom.FamilyRecord) {
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

func contains(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("contains() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {

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

		lint_I100(rec)
		lint_I101(rec)
		lint_I102(rec)
		lint_I103(rec)
		lint_I104(rec)
		lint_I105(rec)
		lint_I106(rec)
		lint_I107(rec)
		lint_I108(rec)
		lint_I109(rec)
		lint_I110(rec)
		lint_I111(rec)
		lint_I112(rec)
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

		lint_F100(rec)
		lint_F101(rec)
		lint_F102(rec)
		lint_F103(rec)
		lint_F104(rec)
		lint_F105(rec)
		lint_F106(rec)
		lint_F107(rec)
		lint_F108(rec)
		lint_F109(rec)
		lint_F110(rec)

		for _, child := range rec.Child {
			if *verbose {
				fmt.Printf("\t%s, %s\n", rec.Xref, child.Xref)
			}
		}
	}
}
