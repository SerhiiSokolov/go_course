package main

import "fmt"

type VisionType string

const (
	day   VisionType = "day"
	night VisionType = "night"
)

type Event struct {
	cameraID  int
	eventType string
	animal    string
	action    string
}

type Camera interface {
	TakeScreenShot(animal string, action string) Event
}

type DayCamera struct {
	cameraId   int
	visionType VisionType
}

func (c DayCamera) TakeScreenShot(animal string, action string) Event {
	fmt.Printf("This is a %s camera registered a %s move for %s\n", c.visionType, action, animal)
	return Event{
		cameraID:  c.cameraId,
		eventType: "Movement",
		animal:    animal,
		action:    action,
	}
}

type NightCamera struct {
	cameraId   int
	visionType VisionType
}

func (c NightCamera) Describe() {
	fmt.Printf("This is a %s camera %d\n", c.visionType, c.cameraId)
}

func (c NightCamera) TakeScreenShot(animal string, action string) Event {
	fmt.Printf("This is a %s camera registered a %s move for %s\n", c.visionType, action, animal)
	return Event{
		cameraID:  c.cameraId,
		eventType: "Movement",
		animal:    animal,
		action:    action,
	}
}

type Zone struct {
	name         string
	dayCameras   map[int]*DayCamera
	nightCameras map[int]*NightCamera
}

func NewZone(name string) *Zone {
	zone := &Zone{name: name, dayCameras: make(map[int]*DayCamera), nightCameras: make(map[int]*NightCamera)}
	return zone
}

func (z Zone) Describe() {
	fmt.Printf("Zone = %s\n", z.name)
	fmt.Printf("Day cameras IDs:")
	for _, camera := range z.dayCameras {
		fmt.Printf("%v,", camera.cameraId)
	}
	fmt.Printf("\nNight cameras IDs:")
	for _, camera := range z.nightCameras {
		fmt.Printf("%v,", camera.cameraId)
	}
	fmt.Println()
}

func (z Zone) RegisterEvent(c Camera, animal string, action string) Event {
	event := c.TakeScreenShot(animal, action)
	return event
}

func SendDataToStorage(e []Event) {
	fmt.Printf("Save the following events to storage:\n")
	for _, event := range e {
		fmt.Printf("Animal %s %s\n", event.animal, event.action)
	}
}
func main() {
	zone1 := NewZone("preadators")
	events := []Event{}
	dayCamera1 := &DayCamera{cameraId: 1, visionType: day}
	nightCamera1 := &NightCamera{cameraId: 1, visionType: night}
	nightCamera2 := &NightCamera{cameraId: 2, visionType: night}
	zone1.dayCameras[dayCamera1.cameraId] = dayCamera1
	zone1.nightCameras[nightCamera1.cameraId] = nightCamera1
	zone1.nightCameras[nightCamera2.cameraId] = nightCamera2
	zone1.Describe()

	events = append(events, zone1.RegisterEvent(dayCamera1, "bear", "MoveLeft"))
	events = append(events, zone1.RegisterEvent(nightCamera1, "tiger", "MoveRight"))
	SendDataToStorage(events)
}
