package main

import "fmt"

type Animal struct {
	name    string
	species string
	isCaged bool
}

type Zoo struct {
	Animals []Animal
}

type Cage struct {
	name   string
	animal *Animal
}

type AllCages struct {
	Cages []Cage
}

type Zookeeper struct {
	name string
}

func NewAnimal(species string, name string) *Animal {
	return &Animal{species: species, name: name}
}

func NewCage(name string) *Cage {
	return &Cage{name: name}
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

func ReproduceAnimals(species string, name string) *Animal {
	return NewAnimal(species, name)
}

// AddAnimalToCage садить всіх звірів із зоопарку в клітки.
// Повертає nil, якщо тварина поміщена в клітку.
// Повертає помилку, якщо всі клітки зайняті і тварина не поміщена в клітку.
func (zk Zookeeper) AddAnimalToCage(c *AllCages, z *Zoo) error {

	for i := range z.Animals {
		animal := &z.Animals[i]
		if !animal.isCaged {

			for j := range c.Cages {
				cage := &c.Cages[j]
				if cage.animal == nil {
					cage.animal = animal
					animal.isCaged = true
					fmt.Printf("Added %s to the cage %s\n", animal.name, cage.name)
					break
				}
			}
			if !animal.isCaged {
				return fmt.Errorf("All cages are occupied, %s is homeless", animal.name)
			}
		}
	}
	return nil
}
func main() {

	lion := NewAnimal("lion", "Simba")
	warthog := NewAnimal("warthog", "Pumba")
	meerkat := NewAnimal("meerkat", "Timon")
	zoopark := Zoo{
		Animals: []Animal{
			*lion,
			*warthog,
			*meerkat,
		},
	}
	cage1 := NewCage("cage1")
	cage2 := NewCage("cage2")
	cage3 := NewCage("cage3")
	cages := AllCages{
		Cages: []Cage{
			*cage1,
			*cage2,
			*cage3,
		},
	}
	zk := NewZookeeper()
	err := zk.AddAnimalToCage(&cages, &zoopark)
	// Якщо усі клітки зайняті то це ок, і доглядач зоопарку має сам вирішити це питання
	if err != nil {
		fmt.Println(err)
	}
	newLion := ReproduceAnimals("lion", "Kiara")
	zoopark.Animals = append(zoopark.Animals, *newLion)
	err = zk.AddAnimalToCage(&cages, &zoopark)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range cages.Cages {
		c.Describe()
	}
}
