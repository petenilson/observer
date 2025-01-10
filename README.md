# Observer

A simple demonstration of the observer design pattern implemented in Go.

*Receivers*

A Receiver reacts to Events that happen and they are unaware of other
Recievers. Receivers should be stateless and have all the information they need
to react to the Event passed in as arguments. This information is present
both in the context passed to the Receiver and the Event itself.


```
go run main.go

ControllerSecurityCamera: Room: Living Room: Motion Detected
ControllerSecurityCamera: Room: Living Room: Recording For 30s
ControllerLights: Room: Living Room: Motion Detected
ControllerLights: Room: Living Room: Turning On Lights
AlerterEmail: Room: Garage: Door Opened
AlerterEmail: Room: Garage: Sending Email
```
