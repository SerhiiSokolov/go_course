package main

import (
	"fmt"
	"time"
)

type rodentType string

const (
	Beaver   rodentType = "beaver"
	Hamster  rodentType = "hamster"
	Squirrel rodentType = "squirrel"
)

type (
	FromTo [2]*Sector

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID        int
		Type      rodentType
		Movements []*Movement
	}

	Sector struct {
		Name    string
		Rodents []*Rodent
	}

	Maze struct {
		Sectors []*Sector
	}
)

func NewRodent(id int, t rodentType) *Rodent {
	return &Rodent{id, t, nil}
}

func (r *Rodent) Move(source *Sector, dest *Sector) {
	now := time.Now()
	move := Movement{FromTo: FromTo{source, dest}, Time: now}
	r.Movements = append(r.Movements, &move)
	source.RemoveRodent(r)
	dest.AddRodent(r)
}

func (r *Rodent) Describe() {
	fmt.Printf("ID=%d, Type=%s\n", r.ID, r.Type)
	if len(r.Movements) == 0 {
		return
	}
	from := r.Movements[0].FromTo[0]
	to := r.Movements[len(r.Movements)-1].FromTo[1]
	fmt.Printf("Started at %s\n", from.Name)
	fmt.Printf("Finished at %s\n", to.Name)
	for _, movement := range r.Movements {
		from := movement.FromTo[0]
		to := movement.FromTo[1]
		fmt.Printf("Moved from %s to %s at %v\n", from.Name, to.Name, movement.Time)
	}
}

func NewSector(name string) *Sector {
	return &Sector{Name: name}
}

func (s *Sector) AddRodent(r *Rodent) {
	s.Rodents = append(s.Rodents, r)
}

func (s *Sector) RemoveRodent(r *Rodent) {
	for i, rd := range s.Rodents {
		if rd.ID == r.ID {
			s.Rodents = append(s.Rodents[:i], s.Rodents[i+1:]...)
			break
		}
	}
}

func (s Sector) Describe() {
	fmt.Println(s.Name)
	for _, rodent := range s.Rodents {
		rodent.Describe()
	}
}

func NewMaze() *Maze {
	return &Maze{}
}

func (m *Maze) AddSectors(s *Sector) {
	m.Sectors = append(m.Sectors, s)
}

func (m Maze) Describe() {
	for _, sector := range m.Sectors {
		sector.Describe()
	}
}

func main() {
	sector1 := NewSector("Sector1")
	sector2 := NewSector("Sector2")
	maze := NewMaze()
	maze.AddSectors(sector1)
	maze.AddSectors(sector2)

	beaver1 := NewRodent(1, Beaver)
	beaver2 := NewRodent(2, Beaver)
	hamster1 := NewRodent(3, Hamster)
	hamster2 := NewRodent(4, Hamster)
	squirrel1 := NewRodent(5, Squirrel)
	squirrel2 := NewRodent(6, Squirrel)

	sector1.AddRodent(beaver1)
	sector1.AddRodent(hamster1)
	sector1.AddRodent(squirrel1)
	sector2.AddRodent(hamster2)
	sector2.AddRodent(beaver2)
	sector2.AddRodent(squirrel2)

	beaver1.Move(sector1, sector2)
	beaver2.Move(sector2, sector1)
	beaver2.Move(sector1, sector2)
	hamster1.Move(sector1, sector2)
	hamster2.Move(sector2, sector1)
	squirrel1.Move(sector1, sector2)
	squirrel2.Move(sector2, sector1)
	fmt.Println("After moving")
	maze.Describe()
}
