# Observer

A simple demonstration of the observer design pattern implemented in Go.

*Receivers*

```Go
type Reciever interface {
	Recieve(c context.Context, e Event, p *Publisher)
}
```

Receivers react to Events. They should be stateless and have all the information they need
to react to the Event passed as arguments.

*Publisher*

Publishers alert recievers of an event. A publisher will keep a registery of recievers
that are interested in a particuluar event.

```
go run main.go

ControllerSecurityCamera: Room: Living Room: Motion Detected
ControllerSecurityCamera: Room: Living Room: Recording For 30s
ControllerLights: Room: Living Room: Motion Detected
ControllerLights: Room: Living Room: Turning On Lights
AlerterEmail: Room: Garage: Door Opened
AlerterEmail: Room: Garage: Sending Email
```
