package main

import (
	"testing"
)

func TestNewZoo(t *testing.T) {
	zoo := NewZoo()
	if zoo == nil {
		t.Errorf("Expected new zoo to be created, got nil")
	}
}

func TestAddArea(t *testing.T) {
	zoo := NewZoo()
	ungulatesArea := NewArea("A1", Ungulates)
	zoo.AddArea("Ungulates", ungulatesArea)

	if len(zoo.areas) != 1 {
		t.Errorf("Expected 1 area to be added, got %d", len(zoo.areas))
	}
	if _, exists := zoo.areas["Ungulates"]; !exists {
		t.Errorf("Expected area 'Ungulates' to exist in the zoo, but it doesn't")
	}
}

func TestNewArea(t *testing.T) {
	area := NewArea("A1", Ungulates)
	if area.name != "A1" || area.animalType != Ungulates {
		t.Errorf("Expected area name to be A1 and type to be Ungulates, got %s and %s", area.name, area.animalType)
	}
}

func TestNewSector(t *testing.T) {
	const (
		Brush Purpose = "Grooming Brush"
	)
	ungulatesArea := NewArea("A1", Ungulates)
	brush := NewTool(Brush)
	allTools := map[string]Tool{
		"Brush": *brush,
	}
	ungulatesUtilityRoom := NewUtilityRoom(allTools)
	sector := ungulatesArea.NewSector("deer", *ungulatesUtilityRoom)
	if sector.subtype != "deer" {
		t.Errorf("Expected sector subtype to be 'deer', got '%s'", sector.subtype)
	}
	if len(sector.animals) != 0 {
		t.Errorf("Expected new sector to have no animals, got %d", len(sector.animals))
	}
	if sector.utilityRoom.tools == nil || sector.utilityRoom.tools["Brush"].purpose != "Grooming Brush" {
		t.Errorf("Expected new sector's utility room to have tool Brush with the purpose \"Grooming Brush\", "+
			"got %+v", sector.utilityRoom.tools["Brush"].purpose)
	}
}

func TestAddAnimal(t *testing.T) {
	sector := &Sector{}
	animal := NewAnimal(1, "Bambi", "deer")
	sector.AddAnimal(animal)

	if len(sector.animals) != 1 {
		t.Errorf("Expected sector to have 1 animal, got %d", len(sector.animals))
	}
	if sector.animals[0].name != "Bambi" || sector.animals[0].id != 1 || sector.animals[0].subtype != "deer" {
		t.Errorf("Expected sector's first animal to be 'Bambi' with id 1 and subtype deer, got '%s' '%d' '%s' ",
			sector.animals[0].name, sector.animals[0].id, sector.animals[0].subtype)
	}
}

func TestFindAnimalByName(t *testing.T) {
	zoo := NewZoo()
	ungulatesArea := NewArea("A1", Ungulates)
	zoo.AddArea("Ungulates", ungulatesArea)
	sector := ungulatesArea.NewSector("deer", *NewUtilityRoom(nil))
	animal := NewAnimal(1, "Bambi", "deer")
	sector.AddAnimal(animal)

	foundAnimal := zoo.FindAnimalByName("Bambi")
	if foundAnimal == nil || foundAnimal.name != "Bambi" {
		t.Errorf("Expected to find animal 'Bambi', but it was not found")
	}
}
