package main

import (
	"fmt"
)

const (
	Ungulates Types = "Ungulates"
	Birds     Types = "Birds"
	Primates  Types = "Primates"

	Brush  Purpose = "Grooming Brush"
	Bucket Purpose = "Water Bucket"
	Syring Purpose = "Veterinary Syring"
)

type Zoo struct {
	Areas
}

func (z Zoo) FindAnimalByName(name string) *Animal {
	for _, area := range z.Areas {
		for sectorName, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					fmt.Printf("Animal %s with ID %d is located in area %s, sector %s\n", animal.Name, animal.ID, area.Name, sectorName)
					return animal
				}
			}
		}
	}
	return nil
}

func NewZoo() *Zoo {
	temp := make(map[string]Area)
	return &Zoo{temp}
}

func (z Zoo) AddArea(t string, area *Area) {
	z.Areas[t] = *area
}

type Types string

type Area struct {
	Name    string
	Type    Types
	Sectors map[string]*Sector
}

func NewArea(n string, t Types) *Area {
	return &Area{n, t, make(map[string]*Sector)}
}

type Areas map[string]Area

type Sector struct {
	Subtype     string
	Animals     []*Animal
	UtilityRoom UtilityRoom
}

func (s *Sector) AddAnimal(animal *Animal) {
	s.Animals = append(s.Animals, animal)
}

func (s Sector) Describe() {
	fmt.Printf("Sector subtype is %s.\n", s.Subtype)
	for _, animal := range s.Animals {
		animal.Describe()
	}
	s.UtilityRoom.Describe()
}

func (s Sector) FeedAnimal(a Animal) {
	fmt.Printf("%s sector is feeding the %s animal ....\n", s.Subtype, a.Name)
	fmt.Printf("Animal %s is not hungry anymore\n", a.Name)
}

func (a *Area) NewSector(subtype string, u UtilityRoom) *Sector {
	sector := Sector{Subtype: subtype, UtilityRoom: u}
	a.Sectors[subtype] = &sector
	return &sector
}

type UtilityRoom struct {
	//The key is a tool's name
	Tools map[string]Tool
}

func (u UtilityRoom) Describe() {
	for k, v := range u.Tools {
		fmt.Printf("Tool is %s and purpose is %s\n", k, v.Purpose)
	}
}

func NewUtilityRoom(t map[string]Tool) *UtilityRoom {
	return &UtilityRoom{t}
}

type Purpose string

type Tool struct {
	Purpose Purpose
}

func NewTool(p Purpose) *Tool {
	return &Tool{p}
}

type Animal struct {
	ID      int
	Name    string
	Subtype string
}

func (a Animal) Describe() {
	fmt.Printf("Animal ID: %d, Name: %s, Subtype: %s\n", a.ID, a.Name, a.Subtype)
}

func NewAnimal(id int, name string, subtype string) *Animal {
	return &Animal{id, name, subtype}
}

func main() {
	zoo := NewZoo()
	ungulatesArea := NewArea("A1", Ungulates)
	birdsArea := NewArea("A2", Birds)
	primatesArea := NewArea("A3", Primates)
	zoo.AddArea("Ungulates", ungulatesArea)
	zoo.AddArea("Birds", birdsArea)
	zoo.AddArea("Primates", primatesArea)
	brush := NewTool(Brush)
	bucket := NewTool(Bucket)
	syring := NewTool(Syring)
	allTools := map[string]Tool{
		"Brush":  *brush,
		"Bucket": *bucket,
		"Syring": *syring,
	}
	ungulatesUtilityRoom := NewUtilityRoom(allTools)
	birdsUtilityRoom := NewUtilityRoom(allTools)
	primatesUtilityRoom := NewUtilityRoom(allTools)
	deerSector := ungulatesArea.NewSector("deer", *ungulatesUtilityRoom)
	birdsSector := birdsArea.NewSector("eagle", *birdsUtilityRoom)
	gorillasSector := primatesArea.NewSector("gorilla", *primatesUtilityRoom)
	deer := NewAnimal(1, "Bambi", "deer")
	eagle := NewAnimal(2, "Zazu", "eagle")
	gorilla := NewAnimal(3, "Kong", "gorilla")
	deerSector.AddAnimal(deer)
	birdsSector.AddAnimal(eagle)
	gorillasSector.AddAnimal(gorilla)
	deerSector.Describe()
	birdsSector.Describe()
	gorillasSector.Describe()
	deerSector.FeedAnimal(*deer)
	birdsSector.FeedAnimal(*eagle)
	gorillasSector.FeedAnimal(*gorilla)
	animal := zoo.FindAnimalByName("Kong")
	if animal != nil {
		fmt.Printf("Let's describe found animal by name: %s\n", animal.Name)
		animal.Describe()
	}
}
