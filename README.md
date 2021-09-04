# nsr-orff-ms-decryptqrsod

Han Solo was recently appointed General of the Alliance Rebel and seeks to strike a great blow against the Galactic Empire to rekindle the flame of resistance.
The rebel intelligence service has detected a call for help from an Imperial cargo ship adrift in an asteroid field. The ship's manifest is ultra classified, but it is rumored that carries rations and weapons for an entire legion.

My mission: return the source and content of the distress message.

ROUTER MUX
configuration routes and server
src/http/mux-router.go

DOCUMENTATION
    ROUTER MUX : https://github.com/gorilla/mux

    TESTIFY : https://github.com/stretchr/testify

    SWAGGER: 1.- https://goswagger.io/install.html
             2.- https://goswagger.io/use/spec.html for annoations, refer the guide


USE APP
1.- Execute the command "go mod init app" to start workdir
2.- Execute command "go mod tidy" for dependencies

WITHOUT DOCKER COMPOSE
3.- To build the binary run "go build src/cmd/main.go"
4.- To run the binary "./main"
WITH DOCKER COMPOSE
3.- docker-compose up --build

RUN LOCAL SERVER
    "go run src/cmd/main.go"

TEST COVERAGE
1.- Generate coverprofile "go test ./... -covermode=count -coverprofile=coverage.out"
2.- run coverage in Terminal "go tool cover -func=coverage.out"
3.- (optional) run coverage in Browser "go tool cover -html=coverage.out"

LOCAL SWAGGER
Generate file swagger "swagger generate spec -o ./swaggerui/swagger.json" 
UP server "go run src/cmd/main.go" and go to http://localhost:8080/swaggerui/
