package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"github.com/iand/gedcom"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	sql_table = `
        DROP TABLE IF EXISTS note;
		CREATE TABLE note (
						  handle CHARACTER(25) PRIMARY KEY,
						  gid    CHARACTER(25),
						  text   TEXT,
						  format INTEGER,
						  note_type1   INTEGER,
						  note_type2   TEXT,
						  change INTEGER,
						  private BOOLEAN);
		DROP TABLE IF EXISTS name;
		CREATE TABLE name (
						  handle CHARACTER(25) PRIMARY KEY,
						  primary_name BOOLEAN,
						  private BOOLEAN,
						  first_name TEXT,
						  suffix TEXT,
						  title TEXT,
						  name_type0 INTEGER,
						  name_type1 TEXT,
						  group_as TEXT,
						  sort_as INTEGER,
						  display_as INTEGER,
						  call TEXT,
						  nick TEXT,
						  famnick TEXT);
		DROP TABLE IF EXISTS surname;
		CREATE TABLE surname (
						  handle CHARACTER(25),
						  surname TEXT,
						  prefix TEXT,
						  primary_surname BOOLEAN,
						  origin_type0 INTEGER,
						  origin_type1 TEXT,
						  connector TEXT);
		CREATE INDEX idx_surname_handle ON
						  surname(handle);
		DROP TABLE IF EXISTS date;
		CREATE TABLE date (
						  handle CHARACTER(25) PRIMARY KEY,
						  calendar INTEGER,
						  modifier INTEGER,
						  quality INTEGER,
						  day1 INTEGER,
						  month1 INTEGER,
						  year1 INTEGER,
						  slash1 BOOLEAN,
						  day2 INTEGER,
						  month2 INTEGER,
						  year2 INTEGER,
						  slash2 BOOLEAN,
						  text TEXT,
						  sortval INTEGER,
						  newyear INTEGER);
		DROP TABLE IF EXISTS person;
		CREATE TABLE person (
						  handle CHARACTER(25) PRIMARY KEY,
						  gid CHARACTER(25),
						  gender INTEGER,
						  death_ref_handle TEXT,
						  birth_ref_handle TEXT,
						  change INTEGER,
						  private BOOLEAN);
		DROP TABLE IF EXISTS family;
		CREATE TABLE family (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 father_handle CHARACTER(25),
						 mother_handle CHARACTER(25),
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS place;
		CREATE TABLE place (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 title TEXT,
						 value TEXT,
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 code TEXT,
						 long TEXT,
						 lat TEXT,
						 lang TEXT,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS place_ref;
		CREATE TABLE place_ref (
						   handle             CHARACTER(25) PRIMARY KEY,
						   from_place_handle  CHARACTER(25),
						   to_place_handle    CHARACTER(25));
		DROP TABLE IF EXISTS place_name;
		CREATE TABLE place_name (
						  handle        CHARACTER(25) PRIMARY KEY,
						  from_handle   CHARACTER(25),
						  value         CHARACTER(25),
						  lang          CHARACTER(25));
		DROP TABLE IF EXISTS event;
		CREATE TABLE event (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 description TEXT,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS citation;
		CREATE TABLE citation (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 confidence INTEGER,
						 page CHARACTER(25),
						 source_handle CHARACTER(25),
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS source;
		CREATE TABLE source (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 title TEXT,
						 author TEXT,
						 pubinfo TEXT,
						 abbrev TEXT,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS media;
		CREATE TABLE media (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 path TEXT,
						 mime TEXT,
						 desc TEXT,
						 checksum INTEGER,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS repository_ref;
		CREATE TABLE repository_ref (
						 handle CHARACTER(25) PRIMARY KEY,
						 ref CHARACTER(25),
						 call_number TEXT,
						 source_media_type0 INTEGER,
						 source_media_type1 TEXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS repository;
		CREATE TABLE repository (
						 handle CHARACTER(25) PRIMARY KEY,
						 gid CHARACTER(25),
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 name TEXT,
						 change INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS link;
		CREATE TABLE link (
						 from_type CHARACTER(25),
						 from_handle CHARACTER(25),
						 to_type CHARACTER(25),
						 to_handle CHARACTER(25));
		CREATE INDEX idx_link_to ON
						  link(from_type, from_handle, to_type);
		DROP TABLE IF EXISTS markup;
		CREATE TABLE markup (
						 handle CHARACTER(25) PRIMARY KEY,
						 markup0 INTEGER,
						 markup1 TEXT,
						 value INTEGER,
						 start_stop_list TEXT);
		DROP TABLE IF EXISTS event_ref;
		CREATE TABLE event_ref (
						 handle CHARACTER(25) PRIMARY KEY,
						 ref CHARACTER(25),
						 role0 INTEGER,
						 role1 TEXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS person_ref;
		CREATE TABLE person_ref (
						 handle CHARACTER(25) PRIMARY KEY,
						 description TEXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS child_ref;
		CREATE TABLE child_ref (
						 handle CHARACTER(25) PRIMARY KEY,
						 ref CHARACTER(25),
						 frel0 INTEGER,
						 frel1 CHARACTER(25),
						 mrel0 INTEGER,
						 mrel1 CHARACTER(25),
						 private BOOLEAN);
		DROP TABLE IF EXISTS lds;
		CREATE TABLE lds (
						 handle CHARACTER(25) PRIMARY KEY,
						 type INTEGER,
						 place CHARACTER(25),
						 famc CHARACTER(25),
						 temple TEXT,
						 status INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS media_ref;
		CREATE TABLE media_ref (
						 handle CHARACTER(25) PRIMARY KEY,
						 ref CHARACTER(25),
						 role0 INTEGER,
						 role1 INTEGER,
						 role2 INTEGER,
						 role3 INTEGER,
						 private BOOLEAN);
		DROP TABLE IF EXISTS address;
		CREATE TABLE address (
						handle CHARACTER(25) PRIMARY KEY,
						private BOOLEAN);
		DROP TABLE IF EXISTS location;
		CREATE TABLE location (
						 handle CHARACTER(25) PRIMARY KEY,
						 street TEXT,
						 locality TEXT,
						 city TEXT,
						 county TEXT,
						 state TEXT,
						 country TEXT,
						 postal TEXT,
						 phone TEXT,
						 parish TEXT);
		DROP TABLE IF EXISTS attribute;
		CREATE TABLE attribute (
						 handle CHARACTER(25) PRIMARY KEY,
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 value TEXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS url;
		CREATE TABLE url (
						 handle CHARACTER(25) PRIMARY KEY,
						 path TEXT,
						 desc TXT,
						 type0 INTEGER,
						 type1 TEXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS datamap;
		CREATE TABLE datamap (
						 from_handle CHARACTER(25),
						 the_type0 INTEGER,
						 the_type1 TEXT,
						 value_field TXT,
						 private BOOLEAN);
		DROP TABLE IF EXISTS tag;
		CREATE TABLE tag (
						 handle CHARACTER(25) PRIMARY KEY,
						 name TEXT,
						 color TEXT,
						 priority INTEGER,
						 change INTEGER);
        `
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func main() {

	flag.Usage = func() {
		fmt.Println("\ngedcom2sql is a tool for conversion of GEDCOM file to an SQLite DB.")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
	}

	var gedfile = flag.String("file", "", "GEDCOM filename")
	var verbose = flag.Bool("verbose", false, "verbose mode")

	flag.Parse()

	if *gedfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	/*
			if len(g.Header.SourceSystem.Version) > 0 {
				println(g.Header.SourceSystem.Version, g.Header.SourceSystem.Id)
		    }
			if len(g.Header.SourceSystem.Id) > 0 {
				println(g.Header.SourceSystem.Id)
		    }
	*/

	var dbname = strings.TrimSuffix(*gedfile, filepath.Ext(*gedfile)) + ".sqlite"
	db := InitDB(dbname)
	defer db.Close()
	CreateTable(db)

	data, _ := ioutil.ReadFile(*gedfile)
	d := gedcom.NewDecoder(bytes.NewReader(data))
	g, _ := d.Decode()

	if *verbose {
		fmt.Println("\nPersons (XRef, Name, Sex):")
		fmt.Println("--------------------------")
	}
	println("Found persons:", len(g.Individual))
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
// handle CHARACTER(25) PRIMARY KEY, gid CHARACTER(25), gender INTEGER, death_ref_handle TEXT, birth_ref_handle TEXT, change INTEGER, private BOOLEAN);
			stmt, err := db.Prepare("INSERT INTO person (handle, gid, gender, death_ref_handle, birth_ref_handle, change, private) values(?, ?, ?, ?, ?, ?, ?)")
			checkErr(err)
			_, err = stmt.Exec(nil, nil, rec.Sex, nil, nil, false)
			// rec.Name[0].Name, birt_date, birt_plac, deat_date, deat_plac, buri_date, buri_plac, rec.Sex)
			checkErr(err)
		}
	}

	if *verbose {
		fmt.Println("\nFamilies (Xref, Husband, Wife):")
		fmt.Println("-------------------------------")
	}
	println("Found families:", len(g.Family))
	for _, rec := range g.Family {
		if *verbose {
			fmt.Printf("%s\n", rec.Xref)
			fmt.Printf("%s, %s, %s\n", rec.Xref, rec.Husband.Xref, rec.Wife.Xref)
		}
		stmt, err := db.Prepare("INSERT INTO family (famID, husband, wife) values(?, ?, ?)")
		checkErr(err)
		_, err = stmt.Exec(rec.Xref, rec.Husband.Xref, rec.Wife.Xref)
		checkErr(err)

		for _, child := range rec.Child {
			if *verbose {
				fmt.Printf("\t%s, %s\n", rec.Xref, child.Xref)
			}
			stmt, err := db.Prepare("INSERT INTO famchild (famID, child) values(?, ?)")
			checkErr(err)
			_, err = stmt.Exec(rec.Xref, child.Xref)
			checkErr(err)
		}
	}
	println("Database path:", dbname)
}
