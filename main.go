/*
A toy example demonstrating the observer pattern
*/
package main

import (
	"context"
	"fmt"
)

type EventType int

const (
	MotionDetectedEvent EventType = iota
	DoorOpenEvent
)

type Event struct {
	EventType EventType
	Location  string
}

type Reciever interface {
	Recieve(c context.Context, e Event, p *Publisher)
}

// Publisher notifies all listners (notifiers) of an event.
type Publisher struct {
	Recievers map[EventType][]Reciever
}

func (p *Publisher) register(n Reciever, t EventType) {
	p.Recievers[t] = append(p.Recievers[t], n)
}

func (p *Publisher) Publish(ctx context.Context, e Event) {
	for _, v := range p.Recievers[e.EventType] {
		v.Recieve(ctx, e, p)
	}
}

type ControllerLights struct {
}

// Recieve implements Reciever
func (s *ControllerLights) Recieve(ctx context.Context, e Event, p *Publisher) {
	switch e.EventType {
	case MotionDetectedEvent:
		fmt.Printf("ControllerLights: Room: %s: Motion Detected\n", e.Location)
		fmt.Printf("ControllerLights: Room: %s: Turning On Lights\n", e.Location)
	default:
		panic("ControllerLights: Event Type Not Supported")
	}
}

type ControllerSecurityCamera struct {
}

// Recieve implements Reciever
func (s *ControllerSecurityCamera) Recieve(ctx context.Context, e Event, p *Publisher) {
	switch e.EventType {
	case MotionDetectedEvent:
		fmt.Printf("ControllerSecurityCamera: Room: %s: Motion Detected\n", e.Location)
		fmt.Printf("ControllerSecurityCamera: Room: %s: Recording For 30s\n", e.Location)
	default:
		panic("ControllerSecurityCamera: Event Type Not Supported")
	}
}

type AlerterEmail struct {
}

// Recieve implements Reciever
func (s *AlerterEmail) Recieve(ctx context.Context, e Event, p *Publisher) {
	switch e.EventType {
	case DoorOpenEvent:
		switch e.Location {
		case "Garage":
			fmt.Printf("AlerterEmail: Room: %s: Door Opened\n", e.Location)
			fmt.Printf("AlerterEmail: Room: %s: Sending Email\n", e.Location)
		}
	default:
		panic("AlerterEmail: Event Type Not Supported")
	}
}

func main() {

	camera := ControllerSecurityCamera{}
	lights := ControllerLights{}
	sender := AlerterEmail{}

	publisher := Publisher{
		Recievers: map[EventType][]Reciever{},
	}

	// UserSignUp Events
	publisher.register(&camera, MotionDetectedEvent)
	publisher.register(&lights, MotionDetectedEvent)

	// UserUnsubscribe Events
	publisher.register(&sender, DoorOpenEvent)

	publisher.Publish(context.TODO(), Event{
		EventType: MotionDetectedEvent,
		Location:  "Living Room",
	})
	publisher.Publish(context.TODO(), Event{
		EventType: DoorOpenEvent,
		Location:  "Garage",
	})
}
