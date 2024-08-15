package main

import "fmt"

const (
	Lion    AnimalSpecies = "lion"
	Warthog AnimalSpecies = "warthog"
	Meerkat AnimalSpecies = "meerkat"
)

type AnimalSpecies string

type Animal struct {
	name    string
	species AnimalSpecies
	cage    *Cage
}

type Zoo struct {
	Animals []*Animal
	Cages   []*Cage
}

type Cage struct {
	name   string
	animal *Animal
}

type Zookeeper struct {
	name string
}

func NewAnimal(species AnimalSpecies, name string) *Animal {
	return &Animal{species: species, name: name}
}

func NewCage(name string) *Cage {
	return &Cage{name: name}
}

func (a *Animal) Escape() {
	fmt.Printf("Animal %s escaped from the cage %v\n", a.name, a.cage.name)
	a.cage.animal = nil
	a.cage = nil
}

func NewZookeeper() *Zookeeper {
	return &Zookeeper{name: "Bob"}
}

func (c Cage) Describe() {
	if c.animal == nil {
		fmt.Println("Cage is empty!")
	} else {
		fmt.Printf("Cage %s is occupied by animal: %s\n", c.name, c.animal.name)
	}
}

func (a Animal) Reproduce(species AnimalSpecies, name string) *Animal {
	return NewAnimal(species, name)
}

// AddAnimalToCage puts all animals to cages.
// Function returns nil if all good.
// Error is returned in case all cages are already occupied
func (zk Zookeeper) AddAnimalToCage(z *Zoo) error {

	for _, animal := range z.Animals {
		if animal.cage == nil {
			for j := range z.Cages {
				cage := z.Cages[j]
				if cage.animal == nil {
					cage.animal = animal
					animal.cage = cage
					fmt.Printf("Added %s to the cage %s\n", animal.name, cage.name)
					break
				}
			}
			if animal.cage == nil {
				return fmt.Errorf("All cages are occupied, %s is homeless", animal.name)
			}
		}
	}
	return nil
}

func main() {
	simba := NewAnimal("asdkjhasd", "Simba")
	fmt.Println(simba.species)
	nala := NewAnimal(Lion, "Nala")
	pumba := NewAnimal(Warthog, "Pumba")
	timon := NewAnimal(Meerkat, "Timon")
	cage1 := NewCage("cage1")
	cage2 := NewCage("cage2")
	cage3 := NewCage("cage3")
	cage4 := NewCage("cage4")
	cage5 := NewCage("cage5")
	zoopark := Zoo{
		Animals: []*Animal{
			simba,
			nala,
			pumba,
			timon,
		},
		Cages: []*Cage{
			cage1,
			cage2,
			cage3,
			cage4,
			cage5,
		},
	}
	zk := NewZookeeper()
	err := zk.AddAnimalToCage(&zoopark)
	// It's ok if all cages are busy, it's zookeeper's responsibility to figure out what to do next
	if err != nil {
		fmt.Println(err)
	}
	kiara := nala.Reproduce("lion", "Kiara")
	zoopark.Animals = append(zoopark.Animals, kiara)
	err = zk.AddAnimalToCage(&zoopark)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range zoopark.Cages {
		c.Describe()
	}
	nala.Escape()
	timon.Escape()
	for _, c := range zoopark.Cages {
		c.Describe()
	}
}
