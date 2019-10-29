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
        DROP TABLE IF EXISTS famchild;
		CREATE TABLE famchild (
		   famID varchar(40) NOT NULL DEFAULT '',
		   child varchar(40) NOT NULL DEFAULT '',
		   PRIMARY KEY (famID, child)
		);

		DROP TABLE IF EXISTS family;
		CREATE TABLE family (
		   famID varchar(40) NOT NULL DEFAULT '',
		   husband varchar(40) DEFAULT NULL,
		   wife varchar(40) DEFAULT NULL,
		   marr_date varchar(255) DEFAULT NULL,
		   marr_plac varchar(255) DEFAULT NULL,
		   marr_sour varchar(255) DEFAULT NULL,
		   marb_date varchar(255) DEFAULT NULL,
		   marb_plac varchar(255) DEFAULT NULL,
		   marb_sour varchar(255) DEFAULT NULL,
		   PRIMARY KEY (famID)
		);

		DROP TABLE IF EXISTS person_st;
		CREATE TABLE person_st (
		   persID varchar(40) NOT NULL DEFAULT '',
		   name varchar(255) DEFAULT NULL,
		   vorname varchar(255) DEFAULT NULL,
		   marname varchar(255) DEFAULT NULL,
		   sex char(1) DEFAULT NULL,
		   birt_date varchar(255) DEFAULT NULL,
		   birt_plac varchar(255) DEFAULT NULL,
		   birt_sour varchar(255) DEFAULT NULL,
		   taufe_date varchar(255) DEFAULT NULL,
		   taufe_plac varchar(255) DEFAULT NULL,
		   taufe_sour varchar(255) DEFAULT NULL,
		   deat_date varchar(255) DEFAULT NULL,
		   deat_plac varchar(255) DEFAULT NULL,
		   deat_sour varchar(255) DEFAULT NULL,
		   buri_date varchar(255) DEFAULT NULL,
		   buri_plac varchar(255) DEFAULT NULL,
		   buri_sour varchar(255) DEFAULT NULL,
		   occupation varchar(255) DEFAULT NULL,
		   occu_date varchar(255) DEFAULT NULL,
		   occu_plac varchar(255) DEFAULT NULL,
		   occu_sour varchar(255) DEFAULT NULL,
		   religion varchar(80) DEFAULT NULL,
		   confi_date varchar(255) DEFAULT NULL,
		   confi_plac varchar(255) DEFAULT NULL,
		   confi_sour varchar(255) DEFAULT NULL,
		   note longtext,
		   PRIMARY KEY (persID)
		);
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
		fmt.Println("\ngedcom2sql is a tool for conversion of GEDCOM to SQLite.")
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
		fmt.Println("\nPersons (Xref, Name, Sex):")
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
			stmt, err := db.Prepare("INSERT INTO person_st (persID, name, birt_date, birt_plac, deat_date, deat_plac, buri_date, buri_plac, sex) values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
			checkErr(err)
			_, err = stmt.Exec(rec.Xref, rec.Name[0].Name, birt_date, birt_plac, deat_date, deat_plac, buri_date, buri_plac, rec.Sex)
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
