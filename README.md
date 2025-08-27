# Playing with design patterns: Observer Pattern

A simple demonstration of the observer design pattern implemented in Go.

_Receivers_

```Go
type Reciever interface {
	Recieve(c context.Context, e Event, p *Publisher)
}
```

Receivers react to Events. They should be stateless and have all the information they need
to react to the Event passed as arguments.

_Publisher_

Publishers alert recievers of an event. A publisher will keep a registery of recievers
that are interested in a particuluar event.

```Go
func (p *Publisher) Publish(ctx context.Context, e Event) {
	for _, v := range p.Recievers[e.EventType] {
		if err := v.Recieve(ctx, e, p); err != nil {
			fmt.Println(err)
		}
	}
}
```

```Bash
go run main.go

ControllerSecurityCamera: Room: Living Room: Motion Detected
ControllerSecurityCamera: Room: Living Room: Recording For 30s
ControllerLights: Room: Living Room: Motion Detected
ControllerLights: Room: Living Room: Turning On Lights
AlerterEmail: Room: Garage: Door Opened
AlerterEmail: Room: Garage: Sending Email
```
