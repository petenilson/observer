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

type Notifier interface {
	Notifiy(c context.Context, e Event, p *Publisher)
}

// Publisher notifies all listners (notifiers) of an event.
type Publisher struct {
	Notifiers map[EventType][]Notifier
}

func (p *Publisher) register(n Notifier, t EventType) {
	p.Notifiers[t] = append(p.Notifiers[t], n)
}

func (p *Publisher) Publish(ctx context.Context, e Event) {
	for _, v := range p.Notifiers[e.EventType] {
		v.Notifiy(ctx, e, p)
	}
}

// WelcomeGifter sends out a small gift to new customers
type WelcomeGifter struct {
}

// Notifiy implements Notifier
func (s *WelcomeGifter) Notifiy(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("WelcomeGifter: Sending Gift")

}

// WelcomeEmailer sends out emails to New Signups
type WelcomeEmailer struct {
}

// Notifiy implements Notifier
func (s *WelcomeEmailer) Notifiy(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("WelcomeEmailer: Sending Email")
}

// ProtomotionUnsubscriber will remove unregister the user from promotional emails
type ProtomotionUnsubscriber struct {
}

// Notifiy implements Notifier
func (s *ProtomotionUnsubscriber) Notifiy(ctx context.Context, e Event, p *Publisher) {
	fmt.Println("ProtomotionUnsubscriber: Removing User From Mailing List")
}

func main() {

	mailer := WelcomeEmailer{}
	gifter := WelcomeGifter{}
	unsubr := ProtomotionUnsubscriber{}

	publisher := Publisher{
		Notifiers: map[EventType][]Notifier{},
	}

	// UserSignUp Events
	publisher.register(&mailer, UserSignUp)
	publisher.register(&gifter, UserSignUp)

	// UserUnsubscribe Events
	publisher.register(&unsubr, UserUnsubscribe)

	publisher.Publish(context.TODO(), Event{EventType: UserSignUp})
	publisher.Publish(context.TODO(), Event{EventType: UserUnsubscribe})
}
