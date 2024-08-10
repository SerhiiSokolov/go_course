package main

import "fmt"

// Struct that represents an Animal
type Animal struct {
	name      string // Name of the animal
	species   string // Species of the animal
	inTheCage bool   // Flag indicating whether the animal is in a cage
}

// Struct that represents a Zoo, containing multiple animals
type Zoo struct {
	Animals []Animal // Slice of animals in the zoo
}

// Struct that represents a Cage, which can hold an animal
type Cage struct {
	name   string // Name of the cage
	animal string // Species of the animal currently in the cage
}

// Struct that represents all the cages in the zoo
type AllCages struct {
	Cages []Cage // Slice of cages
}

// Struct that represents a Zookeeper
type Zookeeper struct {
	name string // Name of the zookeeper
}

// Constructor for creating a new Animal
func NewAnimal(species string, name string) *Animal {
	return &Animal{species: species, name: name}
}

// Constructor for creating a new Cage
func NewCage(name string) *Cage {
	return &Cage{name: name}
}

// Constructor for creating a new Zookeeper
func NewZookeeper() *Zookeeper {
	return &Zookeeper{name: "Bob"}
}

// Method that describes the current state of the cage
func (c Cage) DescribeCage() {
	if c.animal == "" {
		fmt.Println("Cage is empty!\n") // Print if the cage is empty
	} else {
		fmt.Printf("Cage %s is occupied by animal: %s\n", c.name, c.animal) // Print the animal in the cage
	}
}

// Function to create a new animal as a result of reproduction
func ReproduceAnimals(species string, name string) *Animal {
	return NewAnimal(species, name) // Calls the NewAnimal constructor to create the new animal
}

// Method for adding animals to cages by the zookeeper
func (zk Zookeeper) AddAnimalToCage(c *AllCages, z *Zoo) error {
	// Iterate over the animals in the zoo
	for i := range z.Animals {
		animal := &z.Animals[i] // Get a pointer to the current animal
		if !animal.inTheCage {  // If the animal is not already in a cage
			// Iterate over the cages
			for j := range c.Cages {
				cage := &c.Cages[j]    // Get a pointer to the current cage
				if cage.animal == "" { // If the cage is empty
					cage.animal = animal.species // Place the animal in the cage
					animal.inTheCage = true      // Mark the animal as being in a cage
					fmt.Printf("Added %s to the cage %s\n", animal.name, cage.name)
					break // Exit the cage loop as the animal has been placed
				}
			}
			if !animal.inTheCage { // If animal is not in the cage, this means all cages are busy
				return fmt.Errorf("All cages are occupied, %s is homeless", animal.name) // Return an error
			}
		}
	}
	return nil // Return nil if all animals were successfully placed in cages
}

func main() {
	// Create new animals
	lion := NewAnimal("lion", "Simba")
	warthog := NewAnimal("warthog", "Pumba")
	meerkat := NewAnimal("meerkat", "Timon")

	// Initialize the zoo with the created animals
	zoopark := Zoo{
		Animals: []Animal{
			*lion,
			*warthog,
			*meerkat,
		},
	}

	// Create cages for the animals
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

	// Create a zookeeper
	zk := NewZookeeper()

	// Try to add animals to cages
	err := zk.AddAnimalToCage(&cages, &zoopark)
	if err != nil {
		fmt.Println(err) // Print any error if animals couldn't be placed in cages
	}

	// Reproduce a new animal and add it to the zoo
	newLion := ReproduceAnimals("lion", "Kiara")
	zoopark.Animals = append(zoopark.Animals, *newLion) // Add the new animal to the zoo

	// Try to add the new animal to a cage
	err = zk.AddAnimalToCage(&cages, &zoopark)
	if err != nil {
		fmt.Println(err) // Print any error if the new animal couldn't be placed in a cage
	}

	// Describe the state of each cage
	for _, c := range cages.Cages {
		c.DescribeCage() // Print the details of each cage
	}
}
