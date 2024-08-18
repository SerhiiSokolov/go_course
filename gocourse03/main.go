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

func (r *Rodent) MoveRodent(source *Sector, dest *Sector) {
	time.Sleep(100 * time.Millisecond)
	now := time.Now()
	move := Movement{FromTo: FromTo{source, dest}, Time: now}
	r.Movements = append(r.Movements, &move)
	source.RemoveRodent(r)
	dest.AddRodent(r)
}

func (r *Rodent) Describe() {
	if len(r.Movements) > 0 {
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

}

func NewSector(name string) *Sector {
	return &Sector{Name: name}
}

func (s *Sector) AddRodent(r *Rodent) {
	s.Rodents = append(s.Rodents, r)
}

func (s *Sector) RemoveRodent(r *Rodent) {
	temp := NewSector(s.Name)
	for _, rd := range s.Rodents {
		if rd.ID != r.ID {
			temp.AddRodent(rd)
		}
	}
	s.Rodents = temp.Rodents
}

func (s Sector) Describe() {
	fmt.Println(s.Name)
	for _, rodent := range s.Rodents {
		fmt.Printf("ID=%d, Type=%s\n", rodent.ID, rodent.Type)
		rodent.Describe()
	}
}

func NewMaze() *Maze {
	return &Maze{}
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
	maze.Sectors = append(maze.Sectors, sector1, sector2)

	Beaver1 := NewRodent(1, Beaver)
	Beaver2 := NewRodent(2, Beaver)
	Hamster1 := NewRodent(3, Hamster)
	Hamster2 := NewRodent(4, Hamster)
	Squirrel1 := NewRodent(5, Squirrel)
	Squirrel2 := NewRodent(6, Squirrel)

	sector1.AddRodent(Beaver1)
	sector1.AddRodent(Hamster1)
	sector1.AddRodent(Squirrel1)
	sector2.AddRodent(Hamster2)
	sector2.AddRodent(Beaver2)
	sector2.AddRodent(Squirrel2)

	Beaver1.MoveRodent(sector1, sector2)
	Beaver2.MoveRodent(sector2, sector1)
	Beaver2.MoveRodent(sector1, sector2)
	Hamster1.MoveRodent(sector1, sector2)
	Hamster2.MoveRodent(sector2, sector1)
	Squirrel1.MoveRodent(sector1, sector2)
	Squirrel2.MoveRodent(sector2, sector1)
	fmt.Println("After moving")
	maze.Describe()
}
