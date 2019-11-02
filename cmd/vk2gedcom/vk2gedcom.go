/*

Example: ./socialtree -id 233293686

TreeWalk https://gist.github.com/abhat/71332461951830f9a0f5
https://gist.github.com/zyxar/2317744
https://golang.org/doc/play/tree.go
https://github.com/alonsovidales/go_graph

FIXME:
		id181554010 (зацикливается)
		id140022470 (зацикливается)

*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type RawResponse struct {
	Person []json.RawMessage `json:"response"`
}

type Person struct {
	ID              int        `json:"id"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Sex             int        `json:"sex"`
	Bdate           string     `json:"bdate"`
	Photo_50        string     `json:"photo_50"`
	Relation        int        `json:"relation"`
	Relatives       []Relative `json:"relatives"`
	RelationPartner RelPartner `json:"relation_partner"`
	Generation      int
}

type RelPartner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Relative struct {
	ID   int    `json:"id"`   // идентификатор пользователя
	Type string `json:"type"` // тип родственной связи
	Name string `json:"name"` // имя родственника
}

func ValueInSlice(n int, list []int) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}

func process_person(person Person) []Person {

	var tree []Person
	var p Person
	if len(person.Relatives) != 0 {
		for _, r := range person.Relatives {
			p = Person{}
			log.Println(r.ID, r.Type, r.Name)
			id := strconv.Itoa(r.ID)
			if r.ID > 0 {
				log.Println("Person is in VK", id)
				p = get_profile(id)
			} else {
				log.Println("Person is absent in VK", id)
				p.FirstName = r.Name
				p.LastName = r.Name
			}

			switch r.Type {
			case "child":
				p.Generation = person.Generation - 1
			case "sibling":
				p.Generation = person.Generation
			case "parent":
				p.Generation = person.Generation + 1
			case "grandparent":
				p.Generation = person.Generation - 2
			case "grandchild":
				p.Generation = person.Generation + 2
			}
			tree = append(tree, p)
			process_person(p)
		}
	}

	relation := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if ValueInSlice(person.Relation, relation) {
		log.Println("no spouse")
	} else {
		p = Person{}
		id := strconv.Itoa(person.RelationPartner.ID)
		p = get_profile(id)
		tree = append(tree, p)
		process_person(p)
	}

	return tree
}

func get_profile(id string) Person {

	url := "https://api.vk.com/method/users.get?user_ids="
	url = url + id
	url = url + "&fields=first_name,last_name,photo_50,relatives,relation,bdate,sex&v=5.67"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var raw RawResponse
	err = json.Unmarshal(body, &raw)
	if err != nil {
		log.Fatal("Error parsing json: ", err)
	}

	var person Person
	err = json.Unmarshal(raw.Person[0], &person)
	if err != nil {
		log.Fatal("Error parsing json: ", err)
	}

	var sex string
	switch person.Sex {
	case 1:
		sex = "Ж"
	case 2:
		sex = "М"
	default:
		sex = "неизвестно"
	}

	log.Println(person.FirstName, person.LastName, sex, person.Bdate)
	log.Println("\t", person.Relatives)

	return person
}

func main() {

	flag.Usage = func() {
		fmt.Println("\nsocialtree is a tool for getting relations via VK.")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
	}

	var person_id = flag.String("id", "", "person id")

	flag.Parse()

	if *person_id == "" {
		flag.Usage()
		os.Exit(1)
	}

	var Tree []Person

	person := get_profile(*person_id)
	person.Generation = 0
	Tree = append(Tree, person)
	Tree = append(Tree, process_person(person)...) // new

	/*
		var p Person
		if len(person.Relatives) != 0 {
			for _, r := range person.Relatives {
				p = Person{}
				log.Println(r.ID, r.Type, r.Name)
				id := strconv.Itoa(r.ID)
				if r.ID > 0 {
					log.Println("Person is in VK", id)
					p = get_profile(id)
				} else {
					log.Println("Person is absent in VK", id)
					p.FirstName = r.Name
					p.LastName = r.Name
				}

				switch r.Type {
				case "child":
					p.Generation = person.Generation - 1
				case "sibling":
					p.Generation = person.Generation
				case "parent":
					p.Generation = person.Generation + 1
				case "grandparent":
					p.Generation = person.Generation - 2
				case "grandchild":
					p.Generation = person.Generation + 2
				}
				Tree = append(Tree, p)
				process_person(p)
			}
		}

		relation := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		if ValueInSlice(person.Relation, relation) {
			log.Println("no spouse")
		} else {
			p = Person{}
			id := strconv.Itoa(person.RelationPartner.ID)
			p = get_profile(id)
			Tree = append(Tree, p)
			process_person(p)
		}
	*/

	prettyJSON, err := json.MarshalIndent(Tree, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Println(string(prettyJSON))
}
