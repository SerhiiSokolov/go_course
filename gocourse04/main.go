package main

import (
	"fmt"
)

const (
	Ungulates AnimalType = "Ungulates"
	Birds     AnimalType = "Birds"
	Primates  AnimalType = "Primates"

	Brush  Purpose = "Grooming Brush"
	Bucket Purpose = "Water Bucket"
	Syring Purpose = "Veterinary Syring"
)

type Zoo struct {
	areas
}

func (z Zoo) FindAnimalByName(name string) *Animal {
	for _, area := range z.areas {
		for _, sector := range area.sectors {
			for _, animal := range sector.animals {
				if animal.name == name {
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
	z.areas[t] = *area
}

type AnimalType string

type Area struct {
	name       string
	animalType AnimalType
	sectors    map[string]*Sector
}

func NewArea(n string, t AnimalType) *Area {
	return &Area{n, t, make(map[string]*Sector)}
}

type areas map[string]Area

type Sector struct {
	subtype     string
	animals     []*Animal
	utilityRoom UtilityRoom
}

func (s *Sector) AddAnimal(animal *Animal) {
	s.animals = append(s.animals, animal)
}

func (s Sector) Describe() {
	fmt.Printf("Sector subtype is %s.\n", s.subtype)
	for _, animal := range s.animals {
		animal.Describe()
	}
	s.utilityRoom.Describe()
}

func (s Sector) FeedAnimal(a Animal) {
	fmt.Printf("%s sector is feeding the %s animal ....\n", s.subtype, a.name)
	fmt.Printf("Animal %s is not hungry anymore\n", a.name)
}

func (a *Area) NewSector(subtype string, u UtilityRoom) *Sector {
	sector := Sector{subtype: subtype, utilityRoom: u}
	a.sectors[subtype] = &sector
	return &sector
}

type UtilityRoom struct {
	// The key is a tool's name
	tools map[string]Tool
}

func (u UtilityRoom) Describe() {
	for k, v := range u.tools {
		fmt.Printf("Tool is %s and purpose is %s\n", k, v.purpose)
	}
}

func NewUtilityRoom(tools map[string]Tool) *UtilityRoom {
	return &UtilityRoom{tools}
}

type Purpose string

type Tool struct {
	purpose Purpose
}

func NewTool(p Purpose) *Tool {
	return &Tool{p}
}

type Animal struct {
	id      int
	name    string
	subtype string
}

func (a Animal) Describe() {
	fmt.Printf("Animal id: %d, name: %s, subtype: %s\n", a.id, a.name, a.subtype)
}

func NewAnimal(id int, name string, subtype string) *Animal {
	return &Animal{id: id, name: name, subtype: subtype}
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
		fmt.Printf("Let's describe found animal by name: %s\n", animal.name)
		animal.Describe()
	}
}
