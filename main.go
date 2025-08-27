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
	Recieve(c context.Context, e Event, p *Publisher) error
}

func NewPublisher() *Publisher {
	return &Publisher{
		Recievers: make(map[EventType][]Reciever),
	}
}

// Publisher notifies all listners (notifiers) of an event.
type Publisher struct {
	Recievers map[EventType][]Reciever
}

func (p *Publisher) Register(r Reciever, t EventType) {
	p.Recievers[t] = append(p.Recievers[t], r)
}

func (p *Publisher) Publish(ctx context.Context, e Event) {
	for _, v := range p.Recievers[e.EventType] {
		if err := v.Recieve(ctx, e, p); err != nil {
			fmt.Println(err)
		}
	}
}

type ControllerLights struct {
	name string
}

// Recieve implements Reciever
func (s *ControllerLights) Recieve(ctx context.Context, e Event, p *Publisher) error {
	switch e.EventType {
	case MotionDetectedEvent:
		fmt.Printf("ControllerLights: Room: %s: Motion Detected\n", e.Location)
		fmt.Printf("ControllerLights: Room: %s: Turning On Lights\n", e.Location)
	default:
		return fmt.Errorf("%s: event type not supported", s.name)
	}
	return nil
}

type ControllerSecurityCamera struct {
	name string
}

// Recieve implements Reciever
func (s *ControllerSecurityCamera) Recieve(ctx context.Context, e Event, p *Publisher) error {
	switch e.EventType {
	case MotionDetectedEvent:
		fmt.Printf("ControllerSecurityCamera: Room: %s: Motion Detected\n", e.Location)
		fmt.Printf("ControllerSecurityCamera: Room: %s: Recording For 30s\n", e.Location)
	default:
		return fmt.Errorf("%s: event type not supported", s.name)
	}
	return nil
}

type AlerterEmail struct {
	name string
}

// Recieve implements Reciever
func (s *AlerterEmail) Recieve(ctx context.Context, e Event, p *Publisher) error {
	switch e.EventType {
	case DoorOpenEvent:
		switch e.Location {
		case "Garage":
			fmt.Printf("AlerterEmail: Room: %s: Door Opened\n", e.Location)
			fmt.Printf("AlerterEmail: Room: %s: Sending Email\n", e.Location)
		}
	default:
		return fmt.Errorf("%s: event type not supported", s.name)
	}
	return nil
}

func main() {
	camera := &ControllerSecurityCamera{name: "Security Camera"}
	lights := &ControllerLights{name: "Lights"}
	sender := &AlerterEmail{name: "Email Alerter"}

	publisher := Publisher{
		Recievers: map[EventType][]Reciever{},
	}

	// MotionDetectedEvent Events
	publisher.Register(camera, MotionDetectedEvent)
	publisher.Register(lights, MotionDetectedEvent)

	// DoorOpenEvent Events
	publisher.Register(sender, DoorOpenEvent)

	publisher.Publish(context.TODO(), Event{
		EventType: MotionDetectedEvent,
		Location:  "Living Room",
	})
	publisher.Publish(context.TODO(), Event{
		EventType: DoorOpenEvent,
		Location:  "Garage",
	})
}

