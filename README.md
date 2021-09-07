# op-fire-quasar

Han Solo was recently appointed General of the Alliance Rebel and seeks to strike a great blow against the Galactic Empire to rekindle the flame of resistance.
The rebel intelligence service has detected a call for help from an Imperial cargo ship adrift in an asteroid field. The ship's manifest is ultra classified, but it is rumored that carries rations and weapons for an entire legion.

My mission: return the source and content of the distress message.

# Router mux
configuration routes and server
src/http/mux-router.go

# Documentation
    ROUTER MUX : https://github.com/gorilla/mux

    TESTIFY : https://github.com/stretchr/testify

    SWAGGER: 1.- https://goswagger.io/install.html
             2.- https://goswagger.io/use/spec.html for annoations, refer the guide


# Use app
    Execute the command :
             1.- "go mod tidy" for dependencies
             2.- "go mod tidy" for dependencies

# Without Docker compose
             1.- To build the binary run "go build src/cmd/main.go"
             2.- To run the binary "./main"

# With Docker compose
             1.- docker-compose up --build

# Run local server
    Execute the command :
             1.- "go run src/cmd/main.go"

# Test coverage
             1.- Generate coverprofile "go test ./... -covermode=count -coverprofile=coverage.out"
             2.- run coverage in Terminal "go tool cover -func=coverage.out"
             3.- (optional) run coverage in Browser "go tool cover -html=coverage.out"

# Local swagger
             1.- Generate file swagger "swagger generate spec -o ./swaggerui/swagger.json" 
             2.- UP server "go run src/cmd/main.go" and go to http://localhost:8080/swaggerui/

# Exposed service
             1.- url base: https://warm-beach-67225.herokuapp.com/  
             2.- swagger url: https://warm-beach-67225.herokuapp.com/swaggerui/
