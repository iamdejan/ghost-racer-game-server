
# [Delayed] Ghost Race - Game Server  
**NOTE: This project is actually finished, but can't be tested in game, because my original plan to make the game in Godot didn't work out, since Godot didn't support MQTT protocol.** 

This is a server made in Go language that will be used of my new project, Ghost Racer. The communication protocol for sending players' position is MQTT, an IoT protocol for message broker.
  
## Getting Started  
Make sure you have all tools needed from prerequisites, then you can start running this server.  
  
### Prerequisites  
In order to run this server, you need:  
1) Go language compiler (at least v1.12.6).  
2) [Goland IDE](https://www.jetbrains.com/go/) by Jetbrains (or you can use [Visual Studio Code](https://code.visualstudio.com/))  
  
### Installing  
1) Clone this repo.  
2) Open the directory, and run this from terminal: `go run cmd/server/main.go`  
  
## Running The Tests  
For you who use Jetbrains Goland, you can simply run tests from file that has `_test.go` suffix.  
  
## Deployment  
I haven't tried to deploy this, so I'm open to suggestions.  
  
## Authors  
- [Giovanni Dejan](http://github.com/iamdejan)