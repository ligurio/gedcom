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
	EI100 = errors.New("EI100: person's age at death is older than _oldage_")
	EI101 = errors.New("EI101: person is baptized before birth")
	EI102 = errors.New("EI102: person dies before birth")
	EI103 = errors.New("EI103: person is buried before birth")
	EI104 = errors.New("EI104: person dies before baptism")
	EI105 = errors.New("EI105: person is buried before baptism")
	EI106 = errors.New("EI106: person is buried before death")
	EI107 = errors.New("EI107: person is baptised after birth year")
	EI108 = errors.New("EI108: person is buried after death year")
	EI109 = errors.New("EI109: person has unkown gender")
	EI110 = errors.New("EI110: person has ambiguous gender")
	EI111 = errors.New("EI111: person has multiple parentage")
	EI112 = errors.New("EI112: person has no family pointers")

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
	EM100 = errors.New("EM100: person marries before birth")
	EM101 = errors.New("EM101: person marries after death")
	EM102 = errors.New("EM102: person has more than _wedder_ spouses")
	EM103 = errors.New("EM103: person marries someone more than _maydec_ years older")
	EM104 = errors.New("EM104: person marries younger than _yngmar_")
	EM105 = errors.New("EM105: person marries older than _oldmar_")
	EM106 = errors.New("EM106: marriage out of order")
	EM107 = errors.New("EM107: marriage before birth from previous marriage")
	EM108 = errors.New("EM108: marriage after birth from subsequent marriage")
	EM109 = errors.New("EM109: homosexual marriage")
	EM110 = errors.New("EM110: person is a female husband")
	EM111 = errors.New("EM111: person is a male wife")
	EM112 = errors.New("EM112: person was a widow(er) longer than _lngwdw_ years")
	EM113 = errors.New("EM113: person lived more than _oldunm_ years and never married")
	EM114 = errors.New("EM114: person has multiple marriages, this one with no spouse and no children")
	EM115 = errors.New("EM115: person has same surname as spouse")

	// Parentage checks:
	EP101 = errors.New("EP101: mother has more than _fecmom_ children")
	EP102 = errors.New("EP102: mother is older than _oldmom_ at time of birth of child")
	EP103 = errors.New("EP103: child is born before mother")
	EP104 = errors.New("EP104: mother is younger than _yngmom_")
	EP105 = errors.New("EP105: mother is dead at birth of child")
	EP106 = errors.New("EP106: same as above, but for father")
	EP107 = errors.New("EP107: child doesn't inherit father's surname")

	// Children checks:
	EC101 = errors.New("EC101: child is born out of order with respect to a previous child")
	EC102 = errors.New("EC102: child is born in the same year as a previous child")
	EC103 = errors.New("EC103: child is born more than _cspace_ years after previous child")
	EC104 = errors.New("EC104: children's births span more than _cbspan_ years")
	EC105 = errors.New("EC105: child is born before parents' marriage")
	EC106 = errors.New("EC106: child has same given name as sibling")

	// Family checks:
	EF101 = errors.New("EF101: family has no members")
	EF102 = errors.New("EF102: family has no parents")
	EF103 = errors.New("EF103: husband missing pointer to family")
	EF104 = errors.New("EF104: family missing pointer to husband")
	EF105 = errors.New("EF105: wife missing pointer to family")
	EF106 = errors.New("EF106: family missing pointer to wife")
	EF107 = errors.New("EF107: child missing pointer to family")
	EF108 = errors.New("EF108: family missing pointer to child")
	EF109 = errors.New("EF109: family has multiple husbands")
	EF110 = errors.New("EF110: family has multiple wives")
	EF111 = errors.New("EF111: child is in family multiple times")
)

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
			var birt_date, deat_date, buri_date string
			var birt_plac, deat_plac, buri_plac string
			for _, event := range rec.Event {
				if *verbose {
					fmt.Printf("\t%s, %s, %s\n", event.Tag, event.Date, event.Place.Name)
				}
				if event.Tag == "BIRT" {
					birt_date = event.Date
					birt_plac = event.Place.Name
				}
				if event.Tag == "DEAT" {
					deat_date = event.Date
					deat_plac = event.Place.Name
				}
				if event.Tag == "DEAT" {
					buri_date = event.Date
					buri_plac = event.Place.Name
				}
			}
			fmt.Println(birt_date, deat_date, buri_date)
			fmt.Println(birt_plac, deat_plac, buri_plac)
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

		for _, child := range rec.Child {
			if *verbose {
				fmt.Printf("\t%s, %s\n", rec.Xref, child.Xref)
			}
		}
	}
}
