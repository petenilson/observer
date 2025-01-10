package main

import (
	"context"
	"fmt"
)

type EventType int

const (
	UserSignUp EventType = iota
	UserUnsubscribe
)

type Event struct {
	EventType EventType
}

type Reciever interface {
	Recieve(c context.Context, e Event, p *Publisher)
}

// Publisher notifies all listners (notifiers) of an event.
type Publisher struct {
	Notifiers map[EventType][]Reciever
}

func (p *Publisher) register(n Reciever, t EventType) {
	p.Notifiers[t] = append(p.Notifiers[t], n)
}

func (p *Publisher) Publish(ctx context.Context, e Event) {
	for _, v := range p.Notifiers[e.EventType] {
		v.Recieve(ctx, e, p)
	}
}

// WelcomeGifter sends out a small gift to new customers
type WelcomeGifter struct {
}

// Recieve implements Reciever
func (s *WelcomeGifter) Recieve(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("WelcomeGifter: Sending Gift")

}

// WelcomeEmailer sends out emails to New Signups
type WelcomeEmailer struct {
}

// Recieve implements Reciever
func (s *WelcomeEmailer) Recieve(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("WelcomeEmailer: Sending Email")
}

// ProtomotionUnsubscriber will remove unregister the user from promotional emails
type ProtomotionUnsubscriber struct {
}

// Recieve implements Reciever
func (s *ProtomotionUnsubscriber) Recieve(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("ProtomotionUnsubscriber: Removing User From Mailing List")
}

func main() {

	mailer := WelcomeEmailer{}
	gifter := WelcomeGifter{}
	unsubr := ProtomotionUnsubscriber{}

	publisher := Publisher{
		Notifiers: map[EventType][]Reciever{},
	}

	// UserSignUp Events
	publisher.register(&mailer, UserSignUp)
	publisher.register(&gifter, UserSignUp)

	// UserUnsubscribe Events
	publisher.register(&unsubr, UserUnsubscribe)

	publisher.Publish(context.TODO(), Event{EventType: UserSignUp})
	publisher.Publish(context.TODO(), Event{EventType: UserUnsubscribe})
}
