# Say Hello Go

This application is a simple golang test project, using gin gonic for routing. Through the web ui users see the last message someone wrote and can add a new message for the next person to visit the website.

### Normal usage
To get started simply run:
```sh
# get this repo
go get github.com/clem109/say-hello-go

go build main.go
go run main.go

# or alternatively if you have go configured correctly
go install
say-hello-go
```
This exposes port 8080.

### Docker
Incase you forget docker:

```sh
go get github.com/clem109/say-hello-go

docker build -t [your tag] .

docker run -it -p [PORT]:8080 [your tag]
```